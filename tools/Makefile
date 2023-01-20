DGRAPH_VERSION = v22.0.2

current_dir = $(shell pwd)

up: ## Start the zero and alpha containers (local development)
	DGRAPH_VERSION=$(DGRAPH_VERSION) docker-compose up

down: ## Stop the containers (local development)
	DGRAPH_VERSION=$(DGRAPH_VERSION) docker-compose stop

schema-gql: ## Load/update a GraphQL schema (local development)
ifneq (,$(wildcard ../schema/schema.graphql))
	curl --data-binary '@../schema/schema.graphql' --header 'content-type: application/octet-stream' http://localhost:8080/admin/schema
else
	@echo "No schema.graphql found"
endif

drop-data: ## Drops all data (but not the schema, local development)
	curl -X POST localhost:8080/alter -d '{"drop_op": "DATA"}'

drop-all: ## Drops data and schema (local development)
	curl -X POST localhost:8080/alter -d '{"drop_all": true}'

load-data-gql: ## Loads data from a gql-data.json file
ifneq (,$(wildcard ./gql-data.json))
	docker run -it -v $(current_dir):/export dgraph/dgraph:$(DGRAPH_VERSION) dgraph live -a host.docker.internal:9080 -z host.docker.internal:5080 -f /export/gql-data.json
else
	@echo "No gql-data.json file found"
endif

lint: ## Runs lint on the golang tools folder
	staticcheck ./...

help: ## Print target help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
