package model

import (
	"encoding/json"
	"os"

	"github.com/struct-ure/kg/tools/util"
)

// Entry represents a structure (usually defined in a JSON file).
type Entry struct {
	Label []struct {
		Lang  string `json:"lang"`
		Value string `json:"value"`
	} `json:"label"`
	Name []struct {
		Lang  string `json:"lang"`
		Value string `json:"value"`
	} `json:"name"`
	Description []struct {
		Lang  string `json:"lang"`
		Value string `json:"value"`
	} `json:"description"`
	Aliases []struct {
		Lang   string   `json:"lang"`
		Values []string `json:"values"`
	} `json:"aliases"`
	WDID             string   `json:"wdID"`
	PriorID          string   `json:"priorID"`
	URL              string   `json:"url"`
	TypeOf           []string `json:"typeOf"`
	Related          []string `json:"related"`
	Notes            string   `json:"notes"`
	StackOverflowTag string   `json:"stackOverflowTag"`
}

// GetName returns the name in the language (default 'en') specified.
func (e *Entry) GetName(lang string) string {
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

// GetDescription returns the description in the language (default 'en') specified.
func (e *Entry) GetDescription(lang string) string {
	if lang == "" {
		lang = "en"
	}
	for _, v := range e.Description {
		if v.Lang == lang {
			return v.Value
		}
	}
	return ""
}

// SetDescription sets the description by lang.
func (e *Entry) SetDescription(lang, value string) {
	if lang == "" {
		lang = "en"
	}
	for k, v := range e.Description {
		if v.Lang == lang {
			e.Description[k].Value = value
		}
	}
}

// LoadFromFile loads and unmarshals an Entry from a json file. If the label is not
// set in the json file, it sets the label to the parsed file or directory name.
func LoadFromFile(path string) (*Entry, error) {
	var entry Entry
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &entry)
	if err != nil {
		return nil, err
	}
	if len(entry.Label) == 0 {
		entry.Label = []struct {
			Lang  string "json:\"lang\""
			Value string "json:\"value\""
		}{
			{
				Lang:  "en",
				Value: util.LabelFromPath(path),
			},
		}
	}
	return &entry, err
}
