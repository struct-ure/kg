#!/bin/bash

query='query GetByID($getStructureId: String!) {
  getStructure(id: $getStructureId) {
    id
  }
}
'
vars='
{
  "getStructureId": "https://struct-ure.org/kg/it/programming-languages/c"
}
'

query=`echo $query | tr -d '\n'`
vars=`echo $vars | tr -d '\n'`

curl -s \
    -X POST \
    -H "Content-Type: application/json" \
    -d "{\"query\":\"$query\", \"variables\":$vars}" \
    http://localhost:8080/graphql