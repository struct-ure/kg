#!/bin/bash

query='query CountStructures($allFilter: StructureFilter, $leafFilter: StructureFilter) {
  allNodes: aggregateStructure(filter: $allFilter) {
    count
  }
  leafNodes: aggregateStructure(filter: $leafFilter) {
    count
  }
}
'
vars='
{
  "allFilter": {
    "has": "parent"
  },
  "leafFilter": {
    "has": "parent",
    "and": {
      "not": {
        "has": "children"
      }
    }
  }
}
'

query=`echo $query | tr -d '\n'`
vars=`echo $vars | tr -d '\n'`

curl -s \
    -X POST \
    -H "Content-Type: application/json" \
    -d "{\"query\":\"$query\", \"variables\":$vars}" \
    http://localhost:8080/graphql