# struct-ure/kg

struct-ure/kg is a self-contained [knowledge graph](https://en.wikipedia.org/wiki/Knowledge_graph) (KG) of IT skills and technology. It presents a GraphQL API to retrieve information about the graph. Transparent management of the structure and content of the graph is accomplished using git.

### Features

* delivered via Docker
* multilingual
* integration with Wikidata
* GraphQL API

### What can I do with it?

* replace simple tag concepts in your software with identifiers from the graph, e.g., `"C"` becomes `https://struct-ure.org/kg/it/programming-languages/c`
* query the entire graph to build a tree-control to present IT skills and technologies in a UI
* fork the repo and add your company-specific knowledge for use within your organization
* find nodes that have a particular class, e.g., all database nodes that are graph-oriented
* find KG nodes by known aliases
* let your AI infer relationships amongst graph nodes

### Quick start

```sh
docker run -it -p 8080:8080 structureorg/kg
```

then, query the graph using your favorite GraphQL tool. See the `query` folder for example queries.

### Contributing
At present the KG contains over 1,700 concepts — everything from programming languages to electronic health care systems. While a promising start, there's still so much more to add! We're hopeful that domain experts, companies and tech enthusiasts will help move the KG forward.

Additions and improvements to the KG are accomplished by editing the folder and file structure under the /root folder. Please fork this repo and create a pull request with your changes/additions. For more detail on conventions used in the /root folder and how to add/edit KG entries, please see [CONTRIBUTING.md](CONTRIBUTING.md).

### Versioning
struct-ure/kg uses timestamp-based versioning (YY.MM.DD). The first version, 23.01.19 was tagged on January 19, 2023. You can query the KG `queryVersion {version}` for the version.

### Technology
struct-ure/kg is built upon [Dgraph](https://github.com/draph-io/dgraph), a horizontally scalable and distributed GraphQL database with a graph backend. struct-ure/kg is deployed as a simple single-node cluster in its [published Docker image](https://hub.docker.com/r/structureorg/kg/tags) as it's graph structure has fewer than 80k edges (Dgraph can support graphs with hundreds of millions of edges).

The tools used to build Dgraph-compatible import files are written in Go. See the /tools folder for more information.

### To Do
* move the graph build steps to Github actions
* investigate other non-IT domains for inclusion into the KG