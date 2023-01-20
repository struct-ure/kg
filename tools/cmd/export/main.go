package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/struct-ure/kg/tools/model"
	"github.com/struct-ure/kg/tools/util"
	"github.com/struct-ure/kg/tools/wikidata"

	"github.com/pkg/errors"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: export <directory to export>")
		os.Exit(2)
	}
	path := os.Args[1]
	if !util.DirectoryExists(path) {
		fmt.Fprintf(os.Stderr, "Directory '%s' does not exists", path)
		os.Exit(2)
	}
	wdMap := make(map[string]*model.EncodedEntry)
	categories := make([]*model.EncodedEntry, 0)
	entry := &model.Entry{
		Name: []struct {
			Lang  string "json:\"lang\""
			Value string "json:\"value\""
		}{
			{
				Lang:  "en",
				Value: "root",
			},
		},
	}
	root, err := model.NewEncodedEntry(os.Args[1], entry)
	if err != nil {
		panic(err)
	}
	err = loadEntries(os.Args[1], root, wdMap, &categories)
	if err != nil {
		panic(err)
	}

	err = loadWikiDataEntries(wdMap)
	if err != nil {
		panic(err)
	}

	applyRanks([]*model.EncodedEntry{root})

	b, err := json.MarshalIndent(categories, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("1.categories.json", b, 0666)
	if err != nil {
		panic(err)
	}

	b, err = json.MarshalIndent(root, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("0.root.json", b, 0666)
	if err != nil {
		panic(err)
	}
}

func loadEntries(path string, root *model.EncodedEntry,
	wdMap map[string]*model.EncodedEntry, categories *[]*model.EncodedEntry) error {

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ensure the directory is read/created first
		if filepath.Base(path) == "_this.json" {
			return nil
		}
		if info.IsDir() && !strings.HasSuffix(path, "/_categories") {
			path = path + "/_this.json"
		}

		if strings.HasSuffix(path, ".json") {
			entry, err := model.LoadFromFile(path)
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, "_this.json") {
				path = strings.ReplaceAll(path, "/_this.json", "")
			}
			encodedEntry, err := model.NewEncodedEntry(path, entry)
			if err != nil {
				return err
			}
			uri := util.URIFromPath(path)
			if encodedEntry.WDID != "" {
				wdMap[encodedEntry.WDID] = encodedEntry
			}
			if strings.Contains(path, "/_categories/") {
				*categories = append(*categories, encodedEntry)
			} else {
				parentURI := util.ParentFromURI(uri)
				added := model.AddChildEntry(root, parentURI, encodedEntry)
				if !added {
					err = errors.Errorf("child %s not added to parentURI %s", encodedEntry.GetName(""), parentURI)
					return err
				}
			}
		}
		return nil
	})
	return err
}

// Maximum wikidata entries per REST call
const maxWDEntries = 50

func loadWikiDataEntries(wdMap map[string]*model.EncodedEntry) error {

	lookups := make([][]string, len(wdMap)/maxWDEntries+1)
	n := 0
	for k := range wdMap {
		lookups[n/maxWDEntries] = append(lookups[n/maxWDEntries], k)
		n++
	}

	for _, v := range lookups {
		entities, err := wikidata.GetEntities(v)
		if err != nil {
			return errors.Wrap(err, "getting entities from wikidata")
		}
		for k, entity := range entities {
			err = wdMap[k].MergeWikiDataEntity(entity, wdMap)
			if err != nil {
				return errors.Wrap(err, "merging wikidata entity")
			}
		}
	}

	return nil
}

func applyRanks(elements []*model.EncodedEntry) {
	for n, element := range elements {
		element.Rank = n
		if len(element.Children) > 0 {
			applyRanks(element.Children)
		}
	}
}
