package wikidata

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEntities(t *testing.T) {
	require := require.New(t)

	results, err := GetEntities([]string{"Q67280318"}) //Dgraph
	require.NoError(err)
	require.Len(results, 1)
	result, ok := results["Q67280318"]
	require.True(ok)
	require.Equal("https://dgraph.io/", result.Website)
	require.Contains(result.Related, "Q37227") // Go

	results, err = GetEntities([]string{"Q306144"}) //Nginx
	require.NoError(err)
	require.Len(results, 1)
	result, ok = results["Q306144"]
	require.True(ok)
	require.Equal("https://nginx.org/", result.Website)
	require.Subset(result.Related, []string{"Q44484", "Q160453"}) // HTTPS, SMTP
}
