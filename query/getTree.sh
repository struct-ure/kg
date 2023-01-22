#!/bin/bash

query='query GetTree($getStructureId: String!) {
  getStructure(id: $getStructureId) {
    ...structureFields
    children(order: {asc: rank}) {
      ...structureFields
      children(order: {asc: rank}) {
        ...structureFields
        children(order: {asc: rank}) {
          ...structureFields
          children(order: {asc: rank}) {
            ...structureFields
            children(order: {asc: rank}) {
              ...structureFields
            }
          }
        }
      }
    }
  }
}

fragment structureFields on Structure {
  id
  rank
  label {
    lang
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
  url
  stackOverflowTag
  wdID
  typeOf {
    id
  }
  related {
    id
  }
}
'

vars='
{
  "getStructureId": "https://struct-ure.org/kg/it"
}
'

query=`echo $query | tr -d '\n'`
vars=`echo $vars | tr -d '\n'`

curl -s \
    -X POST \
    -H "Content-Type: application/json" \
    -d "{\"query\":\"$query\", \"variables\":$vars}" \
    http://localhost:8080/graphql