#!/bin/bash

count=`../query/countStructures.sh | jq ".data.allNodes.count"`
if [[ $count -lt 10 ]]; then
    echo "countStructures failed"
    exit 1
fi

id=`../query/getStructureByID.sh | jq -r ".data.getStructure.id"`
if [[ "$id" != "https://struct-ure.org/kg/it/programming-languages/c" ]]; then
    echo "getStructureByID failed"
    exit 1
fi

version=`../query/queryVersion.sh | jq -r ".data.queryVersion.version"`
if [[ "version" == "" ]]; then
    echo "queryVersion failed"
    exit 1
fi