# struct-ure/kg

struct-ure/kg is a self-contained [knowledge graph](https://en.wikipedia.org/wiki/Knowledge_graph) (KG) of IT skills and technology. It presents a GraphQL API to retrieve information about the graph. Management of the structure and content of the graph is accomplished using git, which provides transparency and consensus review for evolution.

### Features

* delivered via Docker
* multilingual support for 14 languages
* can pull node information from wikidata
* GraphQL API

### What can I do with it?

* replace simple tag concepts in your software with entries from the graph, e.g., `"C"` becomes `https://struct-ure.org/kg/it/programming-languages/c`
* query the entire graph to build a tree-control to present IT skills and technologies in a UI
* fork the repo and add your company-specific knowledge for use within your organization
* query simple terms to understand where they "fit" within the KG, e.g. the term EC2 falls under the path of `"https://struct-ure.org/kg/it/cloud-computing/amazon-web-services`
* find nodes that have a particular class, e.g., all database nodes that are graph-oriented
* find KG entries by known aliases, in japanese
* let your AI infer relationships amongst graph nodes

### Quickstart

```sh
docker run -it -p 8080:8080 structureorg/kg
```

then, query the graph using your favorite GraphQL tool.

### 