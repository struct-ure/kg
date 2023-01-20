#!/bin/bash

deploy=/Users/matthew/code/struct-ure/deploy/
rootDir=/Users/matthew/code/struct-ure/root

echo "Converting the root folder to graph compatible import format..."
go run cmd/export/main.go $rootDir
if [[ $? -ne 0 ]]; then
    echo "Converstion of root failed"
    exit 1
fi

echo "Starting cluster..."
docker run --name dgraph-loader --rm -d -p 8080:8080 --env BADGER_COMPACTL0ONCLOSE=true -v $deploy/dgraph:/dgraph dgraph/standalone:v22.0.2 > /dev/null
if [[ $? -ne 0 ]]; then
    echo "Start of cluster failed"
    exit 1
fi

echo "Started, waiting for cluster readiness..."
while [[ `curl --silent http://localhost:8080/probe/graphql | jq -r '.status'` !=  up ]]; do
    sleep 1
done

echo "Cluster ready, loading the schema..."
curl --silent --data-binary '@../schema/schema.graphql' --header 'content-type: application/octet-stream' http://localhost:8080/admin/schema

template='{
    "dgraph.type": "Version",
    "Version.version": "%s",
    "Version.timestamp: "%s"
}
'
version=`git tag | sort -V | tail -1`
ts=`date "+%Y-%m-%dT%H:%M:%S%z"`
printf "$template" $version $ts > 2.version.json

echo "Schema loaded, loading the data..."
mkdir $deploy/dgraph/import
cp 0.root.json $deploy/dgraph/import/
cp 1.categories.json $deploy/dgraph/import/
cp 2.version.json $deploy/dgraph/import
docker exec dgraph-loader dgraph live -f /dgraph/import

echo "Data loaded, testing the graph..."
./test.sh
if [[ $? -ne 0 ]]; then
    echo "Test failed"
    exit 1
fi

echo "Graph tested, stopping the alpha..."
curl --silent http://localhost:8080/admin/shutdown

sleep 3

echo "Alpha stopped, stopping the container..."
docker stop dgraph-loader > /dev/null 2>&1

rm -rf $deploy/dgraph/import

echo "Done"
