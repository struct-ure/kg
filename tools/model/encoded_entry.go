package model

import (
	"fmt"
	"strings"

	"github.com/struct-ure/kg/tools/util"
	"github.com/struct-ure/kg/tools/wikidata"

	gowikidata "github.com/Navid2zp/go-wikidata"
	"github.com/pkg/errors"
)

type uidEntry struct {
	UID string `json:"uid"`
}

// EncodedEntry represents a structure entry that will be marshaled to a Dgraph JSON format.
type EncodedEntry struct {
	ID    string `json:"Structure.id"`
	UID   string `json:"uid"`
	Type  string `json:"dgraph.type"`
	Label []struct {
		Type   string    `json:"dgraph.type"`
		Entity *uidEntry `json:"MultilingualLabel.entity"`
		Lang   string    `json:"MultilingualLabel.lang"`
		Value  string    `json:"MultilingualLabel.value"`
	} `json:"Structure.label,omitempty"`
	Name []struct {
		Type   string    `json:"dgraph.type"`
		Entity *uidEntry `json:"MultilingualName.entity"`
		Lang   string    `json:"MultilingualName.lang"`
		Value  string    `json:"MultilingualName.value"`
	} `json:"Structure.name,omitempty"`
	Description []struct {
		Type   string    `json:"dgraph.type"`
		Entity *uidEntry `json:"MultilingualDesc.entity"`
		Lang   string    `json:"MultilingualDesc.lang"`
		Value  string    `json:"MultilingualDesc.value"`
	} `json:"Structure.description,omitempty"`
	Aliases []*struct {
		Type   string    `json:"dgraph.type"`
		Entity *uidEntry `json:"MultilingualAlias.entity"`
		Lang   string    `json:"MultilingualAlias.lang"`
		Values []string  `json:"MultilingualAlias.values"`
	} `json:"Structure.aliases,omitempty"`
	WDID             string      `json:"Structure.wdID,omitempty"`
	Rank             int         `json:"Structure.rank"`
	PriorID          string      `json:"Structure.priorID,omitempty"`
	URL              string      `json:"Structure.url,omitempty"`
	TypeOf           []*uidEntry `json:"Structure.typeOf,omitempty"`
	Related          []*uidEntry `json:"Structure.related,omitempty"`
	Notes            string      `json:"Structure.notes,omitempty"`
	StackOverflowTag string      `json:"Structure.stackOverflowTag,omitempty"`

	Children []*EncodedEntry `json:"Structure.children,omitempty"`
	Parent   *uidEntry       `json:"Structure.parent,omitempty"`
}

// MergeWikiDataEntity merges an Entry record (as read from disk) with results from the wikidata API.
func (entry *EncodedEntry) MergeWikiDataEntity(entity *wikidata.Entity, wdMap map[string]*EncodedEntry) error {
	if entity == nil {
		return errors.New("entity cannot be nil")
	}

	for lang, value := range entity.Labels {
		if util.LangSupported(lang) {
			entry.UpsertName(lang, value.Value)
		}
	}
	for lang, value := range entity.Descriptions {
		if util.LangSupported(lang) {
			entry.UpsertDescription(lang, value.Value)
		}
	}
	for lang, value := range entity.Aliases {
		if util.LangSupported(lang) {
			entry.UpsertAliases(lang, value)
		}
	}
	exists := func(list []*uidEntry, id string) bool {
		for _, v := range list {
			if v.UID == id {
				return true
			}
		}
		return false
	}
	toUID := func(uid string) string { return fmt.Sprintf("_:%s", uid) }
	for _, value := range entity.Categories {
		e, ok := wdMap[value]
		ok = ok && strings.Contains(e.ID, "/categories/")
		if ok {
			if !exists(entry.TypeOf, e.ID) {
				entry.TypeOf = append(entry.TypeOf, &uidEntry{UID: toUID(e.ID)})
			}
		}
	}
	for _, value := range entity.Related {
		e, ok := wdMap[value]
		if ok {
			if !exists(entry.Related, e.ID) {
				entry.Related = append(entry.Related, &uidEntry{UID: toUID(e.ID)})
			}
		}
	}
	if entry.URL == "" {
		entry.URL = entity.Website
		if entry.URL == "" {
			entry.URL = entity.WikipediaURL
		}
	}
	if entry.StackOverflowTag == "" {
		entry.StackOverflowTag = entity.StackOverflowTag
	}
	return nil
}

// UpsertName adds a name to the entry if not already defined.
func (entry *EncodedEntry) UpsertName(lang, value string) {
	found := false
	for _, v := range entry.Name {
		if v.Lang == lang {
			found = true
		}
	}
	if !found {
		entry.Name = append(entry.Name, struct {
			Type   string    `json:"dgraph.type"`
			Entity *uidEntry `json:"MultilingualName.entity"`
			Lang   string    `json:"MultilingualName.lang"`
			Value  string    `json:"MultilingualName.value"`
		}{
			Type:   "MultilingualName",
			Entity: &uidEntry{UID: entry.UID},
			Lang:   lang,
			Value:  value,
		})
	}
}

// UpsertDescription adds a description to the entry if not already defined.
func (entry *EncodedEntry) UpsertDescription(lang, value string) {
	found := false
	for _, v := range entry.Description {
		if v.Lang == lang {
			found = true
		}
	}
	if !found {
		entry.Description = append(entry.Description, struct {
			Type   string    `json:"dgraph.type"`
			Entity *uidEntry `json:"MultilingualDesc.entity"`
			Lang   string    `json:"MultilingualDesc.lang"`
			Value  string    `json:"MultilingualDesc.value"`
		}{
			Entity: &uidEntry{UID: entry.UID},
			Lang:   lang,
			Value:  value,
		})
	}
}

// UpsertAliases adds aliases to the entry if not already defined.
func (entry *EncodedEntry) UpsertAliases(lang string, aliases []gowikidata.Alias) {
	exists := func(a []string, val string) bool {
		for _, v := range a {
			if v == val {
				return true
			}
		}
		return false
	}

	// create the aliases array if not present
	if entry.Aliases == nil {
		entry.Aliases = make([]*struct {
			Type   string    `json:"dgraph.type"`
			Entity *uidEntry `json:"MultilingualAlias.entity"`
			Lang   string    `json:"MultilingualAlias.lang"`
			Values []string  `json:"MultilingualAlias.values"`
		}, 0)
	}
	// if no array exists for the lang, add one
	var found = new(struct {
		Type   string    `json:"dgraph.type"`
		Entity *uidEntry `json:"MultilingualAlias.entity"`
		Lang   string    "json:\"MultilingualAlias.lang\""
		Values []string  "json:\"MultilingualAlias.values\""
	})
	for _, e := range entry.Aliases {
		if e.Lang == lang {
			found = e
		}
	}
	if found.Lang == "" {
		found.Type = "MultilingualAlias"
		found.Entity = &uidEntry{UID: entry.UID}
		found.Lang = lang
		entry.Aliases = append(entry.Aliases, found)
	}
	// for each new alias, if the alias doesn't exist append it
	for _, v := range aliases {
		if !exists(found.Values, v.Value) {
			found.Values = append(found.Values, v.Value)
		}
	}
}

// GetName returns the name in the language (default 'en') specified.
func (e *EncodedEntry) GetName(lang string) string {
	if lang == "" {
		lang = "en"
	}
	for _, v := range e.Name {
		if v.Lang == lang {
			return v.Value
		}
	}
	return ""
}

// NewEncodedEntry creates a record using an Entry.
func NewEncodedEntry(path string, entry *Entry) (*EncodedEntry, error) {
	if path == "" {
		return nil, errors.New("path cannot be empty")
	}
	if entry == nil {
		return nil, errors.New("entry cannot be nil")
	}
	encoded := &EncodedEntry{
		ID:               util.URIFromPath(path),
		UID:              fmt.Sprintf("_:%s", util.URIFromPath(path)),
		Type:             "Structure",
		WDID:             entry.WDID,
		Rank:             util.RankFromPath(path),
		PriorID:          entry.PriorID,
		URL:              entry.URL,
		Notes:            entry.Notes,
		StackOverflowTag: entry.StackOverflowTag,
		Children:         make([]*EncodedEntry, 0),
	}
	for _, v := range entry.Label {
		encoded.Label = []struct {
			Type   string    `json:"dgraph.type"`
			Entity *uidEntry `json:"MultilingualLabel.entity"`
			Lang   string    `json:"MultilingualLabel.lang"`
			Value  string    `json:"MultilingualLabel.value"`
		}{
			{
				Type:   "MultilingualLabel",
				Entity: &uidEntry{UID: encoded.UID},
				Lang:   v.Lang,
				Value:  v.Value,
			},
		}
	}
	if len(entry.Name) > 0 {
		if entry.Name[0].Value != "" {
			for _, v := range entry.Name {
				encoded.Name = []struct {
					Type   string    `json:"dgraph.type"`
					Entity *uidEntry `json:"MultilingualName.entity"`
					Lang   string    `json:"MultilingualName.lang"`
					Value  string    `json:"MultilingualName.value"`
				}{
					{
						Type:   "MultilingualName",
						Entity: &uidEntry{UID: encoded.UID},
						Lang:   v.Lang,
						Value:  v.Value,
					},
				}
			}
		}
	}
	if len(entry.Description) > 0 {
		if entry.Description[0].Value != "" {
			for _, v := range entry.Description {
				encoded.Description = []struct {
					Type   string    `json:"dgraph.type"`
					Entity *uidEntry `json:"MultilingualDesc.entity"`
					Lang   string    `json:"MultilingualDesc.lang"`
					Value  string    `json:"MultilingualDesc.value"`
				}{
					{
						Type:   "MultilingualDesc",
						Entity: &uidEntry{UID: encoded.UID},
						Lang:   v.Lang,
						Value:  v.Value,
					},
				}
			}
		}
	}
	if len(entry.Aliases) > 0 {
		if entry.Aliases[0].Values[0] != "" {
			for _, v := range entry.Aliases {
				encoded.Aliases = []*struct {
					Type   string    `json:"dgraph.type"`
					Entity *uidEntry `json:"MultilingualAlias.entity"`
					Lang   string    `json:"MultilingualAlias.lang"`
					Values []string  `json:"MultilingualAlias.values"`
				}{
					{
						Type:   "MultilingualAlias",
						Entity: &uidEntry{UID: encoded.UID},
						Lang:   v.Lang,
						Values: v.Values,
					},
				}
			}
		}
	}
	toUID := func(uid string) string { return fmt.Sprintf("_:%s", uid) }
	if len(entry.TypeOf) > 0 {
		for _, v := range entry.TypeOf {
			if v != "" {
				encoded.TypeOf = append(encoded.TypeOf, &uidEntry{UID: toUID(v)})
			}
		}
	}
	if len(entry.Related) > 0 {
		for _, v := range entry.Related {
			if v != "" {
				encoded.Related = append(encoded.Related, &uidEntry{UID: toUID(v)})
			}
		}
	}

	switch {
	case encoded.ID == "https://struct-ure.org/kg":
		// do nothing
	case strings.Contains(encoded.ID, "/categories/"):
		// do nothing
	default:
		encoded.Parent = &uidEntry{
			UID: fmt.Sprintf("_:%s", util.ParentFromURI(encoded.ID)),
		}
	}
	return encoded, nil
}

// AddChildEntry recursively adds an entry to specified parent.
func AddChildEntry(tree *EncodedEntry, parentID string, entry *EncodedEntry) bool {
	if tree == nil {
		err := errors.Errorf("tree unexpected nil for %s, parent %s", entry.GetName("en"), parentID)
		panic(err)
	}
	if entry.GetName("en") == "root" {
		return true
	}
	if tree.ID == parentID {
		tree.Children = append(tree.Children, entry)
		return true
	}
	for _, v := range tree.Children {
		added := AddChildEntry(v, parentID, entry)
		if added {
			return true
		}
	}
	return false
}
