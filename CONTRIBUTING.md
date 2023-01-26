# Contributing

Additions and improvements to the KG are accomplished by editing the directory and file structure under the /root folder (in your forked repo). Simply submit a PR back to our main branch. Once merged, a new Docker image will be built and published with your changes.

When making changes, please limit the PR to specific areas. For example, please don't add a programming language definition and an ERP system definition in the same PR. This way, our domain experts (see below) can focus on the domain-specific changes without the distraction of non-related entries.

## Conventions
The /root folder contains the definition of the graph.

![Image of root folder](/img/root-folder.png)

### File names

#### Integer Prefix
File names of both folders and files are prefixed with an integer. This integer is an explicit ranking. For instance, let's say you wanted to add an obscure markup language like RTML (Remote Telescope Markup Language) to the graph under `0.Markup Languages`. While it may be important to you, its use isn't common. You can still add the entry, but rank it lower by prefixing the file name with a higher integer, say 2. So you'd add a new file named `2.RTML.json` to the folder.

The integer ranking also lets graph users [present concepts](https://struct-ure.github.io/kg/examples/ui-tree/) in a logic order. For example, we define Markup Languages, Programming Languages, and other core IT concepts with a zero prefix. We rank things like ERP systems with a '4' prefix.

Without this integer ranking, we'd be at the mercy of the alphabetical sort of folder entries.

#### Labels
The text between the integer prefix and the .json suffix will become the English-specific label in the graph. You can override this (as well as add labels for other languages) by editing the .json file (see 'File Contents' below).

#### File Contents

Every node in the graph corresponds to a .json file in which the entry is defined. For folders, the `_this.json` file present in the folder defines its attributes.

Here is a empty .json file:
```json
{
    "label": [
        {
            "lang": "en",
            "value": ""
        }
    ],
    "name": [
        {
            "lang": "en",
            "value": ""
        }
    ],
    "description": [
        {
            "lang": "en",
            "value": ""
        }
    ],
    "aliases": [
        {
            "lang": "en",
            "values": [
                ""
            ]
        }
    ],
    "wdID": "",
    "priorID": "",
    "url": "",
    "stackOverflowTag": "",
    "typeOf": [
        ""
    ],
    "related": [
        ""
    ],
    "notes": ""
}
```
* `label` is an array of language-specific entries for the node's *label*. If not specified the textual part of the file name will be used as the English entry
* `name` is an array of language-specific entries for the node's *name*. Often they are the same, but in some cases the `label`s will be abbreviations (e.g., XML) and the `name`s will be the full name (e.g., Extensible Markup Language)
* `description` is an array of language-specific entries for descriptions of the node
* `aliases` is an array of language-specific entries of aliases for the the node
* `wdID` is the Wikidata ID for the concept (see below)
* `priorID` notes the ID in the graph this concept was previously located (see Relocating Nodes below)
* `url` notes the primary URL for the concept
* `stackOverflowTag` notes the StackOverflow tag for the concept
* `typeOf` is an array of KG category IDs by which the concept can be categorized (see Categories below)
* `related` is an array of KG category IDs to which the concept relates
* `notes` is a place to record unstructured comments from those maintaining the graph

## URIs
Unique identifiers for nodes in the graph use the `https://struct-ure.org/kg` prefix. For example, the concept of 'skill of writing software for mobile devices' has the unique identifier `https://struct-ure.org/kg/it/skills/mobile-development`. These IDs are constructed automatically using the .json file's location in the /root folder when the graph import data is prepared. See [/tools/util/uri.go]() for details.

## WikiData Integration
When building the standalone graph, struct-ure/kg can pull information from Wikidata regarding the concept. For example, check out the .json definition for `/root/0.IT/0.Programming Languages/0.Go.json`:

```json
{
	"wdID": "Q37227"
}
```

When the conversion tool sees this, it will pull information from Wikidata for the Go node. When a `wdID` value is present, defined attributes in the .json take precedence. For example, I could add an alias to the above `0.Go.json` like so:

```json
{
	"wdID": "Q37227",
    "aliases": [
        {
            "lang": "en",
            "values": [
                "Go-lang"
            ]
        }
    ]
}
```
This would result in a graph node that has the "Golang" alias (from Wikidata) and the "Go-lang" alias (from my entry).

## Languages
struct-ure/kg supports 14 popular languages. See the [/tools/util/lang.go]() file for the definitions.

## Relocating Nodes
When deciding to rename a node or change its location within the graph (by moving it in the /root folder), please note its prior URI-based ID in the `priorID` field.

## Domain Experts
The KG has over 1,700 concepts at present. We welcome domain experts to both flesh out existing areas or add areas that are missing. If you'd like to take on an area of the graph as a domain expert, please reach out to `matthew@struct-ure.org`. As a domain expert, we'll rely on you to approve and merge domain-specific changes to the graph.

## Knowledge Graph
Ontology experts might notice that our [A-Boxes](https://en.wikipedia.org/wiki/Abox) and T-Boxes are intermixed in the graph. This was a conscience choice primarily made to keep the knowledge graph easy to consume and use in real-world applications. 

### Categories
We also use *Categories* liberally. For example, *PostgreSQL* is first and foremost a database. Fittingly, it's located under the /root/0.IT/1.Databases folder. But it's also a *relational* database. Instead of creating folders under the `1.Databases` folder for Relational, Graph, Columnar types (and moving entries under those) we apply *Categories*. This way the high level concept of *Databases* stays uncluttered. You could then query all Database nodes that have the Relational category for example.

Use of categories also avoids sticky situations where a concept is more than one thing. For instance, the programming language 'C' is both a *structured* and *compiled* language. If we were to further sub-divide `Programming Languages` by `Programming Languages/Structured` and `Programming Languages/Compiled`, where would we put `C.json`?

