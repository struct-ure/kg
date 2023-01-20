package util

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	stripRank    = regexp.MustCompile(`^\d.`)
	stripRanks   = regexp.MustCompile(`/\d.`)
	domainPrefix = "https://struct-ure.org/kg"
)

// URIDomainPrefix returns the struct-ure/kg domain prefix.
func URIDomainPrefix() string {
	return domainPrefix
}

// URIFromPath calculates a kg URI from the file path.
func URIFromPath(path string) string {
	// find the top level root directory in the path
	index := strings.Index(path, "/root")
	if index == -1 {
		panic(errors.Errorf("Invalid path, /root not found in path %s", os.Args[1]))
	}
	path = path[index+5:] // len of '/root'

	// strip the filename ranking (e.g., the zero from 0.NFS.json)
	uri := stripRanks.ReplaceAllString(path, "/")
	// convert to lowercase
	uri = strings.ToLower(uri)
	// replace spaces with dashes
	uri = strings.ReplaceAll(uri, " ", "-")
	// trim off trailing file type
	uri = strings.TrimSuffix(uri, ".json")
	// if within a categories folder, strip the underscore
	uri = strings.ReplaceAll(uri, "/_categories/", "/categories/")
	// append the prefix
	uri = domainPrefix + uri
	return uri
}

// ParentFromURI returns the parent of a given kg URI.
func ParentFromURI(uri string) string {
	elements := strings.Split(uri, "/")
	return strings.Join(elements[:len(elements)-1], "/")
}

// RankFromPath returns the numeral rank of a file.
func RankFromPath(path string) int {
	var (
		err  error
		rank int
	)
	base := filepath.Base(path)
	textRank := stripRank.FindString(base)
	if textRank != "" {
		textRank = strings.TrimRight(textRank, ".")
		rank, err = strconv.Atoi(textRank)
		if err != nil {
			rank = 0
		}
	}
	return rank
}

func LabelFromPath(path string) string {
	path = strings.TrimSuffix(path, "/_this.json")
	base := filepath.Base(path)
	base = string(stripRank.ReplaceAll([]byte(base), []byte{}))
	base = strings.ReplaceAll(base, "(slash)", "/")
	return strings.TrimSuffix(base, ".json")
}
