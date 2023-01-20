package wikidata

import (
	"fmt"
	"strings"

	gowikidata "github.com/Navid2zp/go-wikidata"
	"github.com/pkg/errors"
)

type wdResult struct {
	ID          string
	Label       string
	Description string
}

// Searches for up to ten wikidata english-language entities by a term.
func SearchEntity(term string) ([]*wdResult, error) {
	req, _ := gowikidata.NewSearch(term, "en")
	req = req.SetLimit(10)

	res, err := req.Get()
	if err != nil {
		return nil, errors.Wrapf(err, "making wd request for term %s", term)
	}
	if len(res.SearchResult) == 0 {
		return nil, errors.Errorf("No results for '%s'", term)
	}

	results := make([]*wdResult, 0)
	for _, v := range res.SearchResult {
		// don't include published papers in results
		if strings.Contains(strings.ToLower(v.Description), "scientific article") {
			continue
		}
		results = append(results, &wdResult{
			ID:          v.ID,
			Label:       v.Label,
			Description: v.Description,
		})
	}
	if len(results) == 0 {
		return nil, errors.Errorf("no results for '%s'", term)
	}
	return results, nil
}

// Entity represents the result of a Wikidata search/get.
type Entity struct {
	ID               string
	Labels           map[string]gowikidata.Label
	Descriptions     map[string]gowikidata.Description
	Aliases          map[string][]gowikidata.Alias
	StackOverflowTag string
	WikipediaURL     string
	Website          string
	Categories       []string
	Related          []string
}

// Label returns a label for given lang.
func (e *Entity) Label(lang string) (label string) {
	entry, ok := e.Labels[lang]
	if ok {
		label = entry.Value
	}
	return
}

// Description returns a description for given lang.
func (e *Entity) Description(lang string) (description string) {
	entry, ok := e.Descriptions[lang]
	if ok {
		description = entry.Value
	}
	return
}

// GetEntities returns Entities from the Wikidata API.
func GetEntities(entities []string) (map[string]*Entity, error) {
	req, err := gowikidata.NewGetEntities(entities)
	if err != nil {
		return nil, err
	}
	req.SetSites([]string{"enwiki"})

	res, err := req.Get()
	if err != nil {
		return nil, err
	}

	results := make(map[string]*Entity)
	for k, v := range *res {
		results[k] = &Entity{
			ID:           v.ID,
			Labels:       v.Labels,
			Descriptions: v.Descriptions,
			Aliases:      v.Aliases,
		}
		if stackOverflowTag, ok := v.Claims["P1482"]; ok {
			results[k].StackOverflowTag = stackOverflowTag[0].MainSnak.DataValue.Value.Data.(string)
		}

		// instanceOf
		claims, ok := v.Claims["P31"]
		if ok {
			for _, claim := range claims {
				results[k].Categories = append(results[k].Categories, claim.MainSnak.DataValue.Value.Data.(gowikidata.DataValueFields).ID)
			}
		}
		// wikipedia page
		if len(v.SiteLinks) > 0 {
			results[k].WikipediaURL = fmt.Sprintf("https://en.wikipedia.org/wiki/%s",
				strings.ReplaceAll(v.SiteLinks["enwiki"].Title, " ", "_"))
		}
		// P856: official website
		website, ok := v.Claims["P856"]
		if ok {
			claim := website[0].MainSnak
			results[k].Website = claim.DataValue.Value.S
		}

		const (
			programmedIn = "P277"
			basedOn      = "P144"
			influencedBy = "P737"
			supports     = "P3985"
			protocol     = "P2700"
		)
		relatedPropertyIDs := []string{programmedIn, basedOn, influencedBy, supports, protocol}
		// related
		for _, claim := range relatedPropertyIDs {
			related, ok := v.Claims[claim]
			if ok {
				for _, claim := range related {
					results[k].Related = append(results[k].Related, claim.MainSnak.DataValue.Value.Data.(gowikidata.DataValueFields).ID)
				}
			}
		}
	}

	return results, nil
}
