# struct-ure/kg

struct-ure/kg is a self-contained [knowledge graph](https://en.wikipedia.org/wiki/Knowledge_graph) (KG) of tech skills and IT stuff (software, platforms, etc.). It presents a GraphQL API to retrieve information from the graph. Transparent management of the structure and content of the graph is accomplished using gitflow. "Editing" the KG is as simple as making changes to directories and files.

### Features

* simple to contribute (fork -> edit files -> submit pr)
* single Docker image
* multilingual
* integration with Wikidata
* GraphQL API

### What can I do with it?

* replace simple tag concepts in your software with identifiers from the graph, e.g., `"C"` becomes `https://struct-ure.org/kg/it/programming-languages/c`
* query the entire graph to build a tree-control to present tech skills and software in a UI (for example: [this](https://struct-ure.github.io/kg/examples/ui-tree/))
* fork the repo and add your company/domain-specific knowledge for use within your organization
* find nodes that have a particular class, e.g., all database nodes that are graph-oriented
* find nodes by known aliases, e.g., "Golang" is an alias for the programming language "Go"
* let your AI infer relationships, e.g., "EC2" is a part of "AWS" which is part of "Cloud Computing" which is part of "IT"

### Quick start

```sh
docker run -it -p 8080:8080 structureorg/kg
```

then, query the graph using your favorite GraphQL tool. See the `query` folder for example queries. 

Images for both amd64 and arm64 are available on [Dockerhub](https://hub.docker.com/r/structureorg/kg/tags).

### Contributing
At present the KG contains over 1,700 concepts — everything from programming languages to electronic health care systems. While a promising start, there's still so much more to add! We're hopeful that domain experts, companies and tech enthusiasts will help move the KG forward.

Additions and improvements to the KG are accomplished by editing the directory and file structure under the /root folder. Please fork this repo and create a pull request with your changes/additions. For more detail on conventions used in the /root folder and how to add/edit KG entries, please see [CONTRIBUTING.md](CONTRIBUTING.md). For simple fixes (e.g., a spelling error, feel free to open an issue).

### Questions
Please use the [Discussions](https://github.com/struct-ure/kg/discussions) board for questions and suggestions.
### Versioning
struct-ure/kg uses timestamp-based versioning in the YY.MM.DD format. The first version, 23.01.19, was tagged on January 19, 2023. You can query the KG's current version via GraphQL `queryVersion {version}`.

### Technology
struct-ure/kg is built upon [Dgraph](https://github.com/draph-io/dgraph), a horizontally scalable and distributed GraphQL database with a graph backend. struct-ure/kg is deployed as a simple single-node cluster in its [published Docker image](https://hub.docker.com/r/structureorg/kg/tags). At present, the KG has fewer than 60k edges (Dgraph can support graphs with hundreds of millions of edges when deployed in high availability).

The [tools](/tools) used to build graph-compatible import files are written in Go.

### To Do
* move the graph build and image publish steps to Github actions
* investigate other non-IT domains for inclusion into the KG
