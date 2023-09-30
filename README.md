# better-science

This is a better way to encode and use builing science rules. It has a few basic principles:

1. Maintain a high standard of code quality.
2. Use a knowledge graph style of [MathGraph](https://www.researchgate.net/publication/332580594_MathGraph_A_Knowledge_Graph_for_Automatically_Solving_Mathematical_Exercises) to encode and use knowledge.
3. Separate the encoded knowledge into a data representation that is distinct from how the knowledge graph engine uses the knowledge.

Fundementally, the expectation is the is a set of knowledge graphs: one for each "application" (eg. "sizing a GSHP"). Each knowledge graph will have three types of nodes: input, interim, and output nodes. Input nodes will always be pre-populated with a value (either by a user input, or with a default value), interim nodes represent interim steps in a calculation, which may or may not be needed to be filled to get an output. And output nodes are in the nodes that the application will try to calculate.


## Development Environment Setup

To get the development environment set up, install the following tools.

### Go

Follow online guides to set this up.

### Golds (Auto Documentation)

For automated documentation generation, use [Golds](https://github.com/go101/golds). Installation instructions are [here](https://go101.org/apps-and-libs/golds.html). You can then use this to run a 

### Add Go Binaries to your Path

Golang binaries are usually installed in `~/go/bin`. You can check this exists if you install Golds (above). Add this to your path.

## Code Structure

The structure of this repo is heavily based on the [example Golang repository](https://github.com/golang-standards/project-layout/tree/master) structure by [golang-standards](https://github.com/golang-standards).


Most code is in the `/internal/` directory. These are the main different layers that are used to set up the separation of concerns.

| layer | description |
| --- | ---|
| `common` | Core utility functions. Often these will be third party packages wrapped in my own code to avoid vendor lock in |
| `data_access` | Provides the basic data model and operations to use data stored in DBs. Should only be relied upon by the `data_service` |
| `data_service` | Provides the business operations needed by the rest of the application |
| `kg_engine` | Provides a way to use a knowledge graph to calculate new values |
| `knowledge_graph` | Provides a way to represent knowledge as a graph. Is dependant on the `/knowledge'`directory, where the knowledge is actually stored. | 

The final important directories are: `/knowledge` and `/cmd`. 

`/knowledge` is built of three directories:

1. `./terms` which serves as the single source of truth of terms.
2. `./applications` which provides a set of applications, each with it's own `input_nodes`, `interim_nodes`, `output_nodes` and `equations` directories.
3. `./default_values` which provides a way of getting default values for input terms, either via a calculation from other 

`/cmd` is the control layer, where the application is set up and run.


