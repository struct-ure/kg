#!/bin/bash

query='query QueryStructuresByTermByLang($term: String!, $lang: String!) {
	aggregateStructureByTermByLang(term: $term, lang: $lang) {
		count
	}
	queryStructureByTermByLang(term: $term, lang: $lang) {
		id
		parent {
			id
		}
		label {
			value
		}
		name {
			lang
			value
		}
		description {
			lang
			value
		}
		aliases {
			lang
			values
		}
	}
}
'

vars='
{
  "term": "C",
  "lang": "en"
}
'

query=`echo $query | tr -d '\n'`
vars=`echo $vars | tr -d '\n'`


curl -s \
    -X POST \
    -H "Content-Type: application/json" \
    -d "{\"query\":\"$query\", \"variables\":$vars}" \
    http://localhost:8080/graphql