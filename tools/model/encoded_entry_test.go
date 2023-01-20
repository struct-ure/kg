package model

import (
	"testing"

	"github.com/struct-ure/kg/tools/wikidata"

	"github.com/stretchr/testify/require"
)

func TestNewEncodedEntry(t *testing.T) {
	entry := &Entry{
		Name: []struct {
			Lang  string "json:\"lang\""
			Value string "json:\"value\""
		}{
			{
				Lang:  "en",
				Value: "OpenCV",
			},
		},
	}
	encoded, err := NewEncodedEntry("/foo/bar/root/0.IT/2.APIs and Libraries/0.OpenCV1.json", entry)

	require.NoError(t, err)
	require.NotNil(t, encoded)

	require.Equal(t, "https://struct-ure.org/kg/it/apis-and-libraries/opencv1", encoded.ID)
	require.Equal(t, "_:https://struct-ure.org/kg/it/apis-and-libraries/opencv1", encoded.UID)
}

func TestAliasUpsert(t *testing.T) {
	entity, err := wikidata.GetEntities([]string{"Q37227"}) // Golang
	require.NoError(t, err)

	entry := &Entry{
		Aliases: []struct {
			Lang   string   `json:"lang"`
			Values []string `json:"values"`
		}{
			{
				Lang:   "en",
				Values: []string{"golang"},
			},
		},
	}
	encodedEntry, err := NewEncodedEntry("/foo/bar/root/0.IT/2.APIs and Libraries/0.Open CV1.json", entry)
	require.NoError(t, err)
	err = encodedEntry.MergeWikiDataEntity(entity["Q37227"], map[string]*EncodedEntry{})
	require.NoError(t, err)
	require.Equal(t, "غو", encodedEntry.GetName("ar"))

	found := false
	for _, v := range encodedEntry.Aliases {
		if v.Lang == "ar" {
			found = true
			require.Subset(t, v.Values, []string{"GO لغة برمجة", "جو"})
		}
	}
	require.True(t, found)
}
