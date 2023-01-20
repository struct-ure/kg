#!/bin/bash

query='query QueryVersion {
  queryVersion {
    version
    timestamp
  }
}
'
vars='
{}
'

query=`echo $query | tr -d '\n'`
vars=`echo $vars | tr -d '\n'`

curl -s \
    -X POST \
    -H "Content-Type: application/json" \
    -d "{\"query\":\"$query\", \"variables\":$vars}" \
    http://localhost:8080/graphql