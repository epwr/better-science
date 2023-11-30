# better-science

For over two years, I have been developing an expert system that encodes knowledge from Building Scientists and uses it to create decarbonization plans for carbon real estate. This repo is my escape from the confines of a startup that is focused on shipping new functionality and is where I can dream of new ways to encode expert knowledge.

## Development Environment Setup

To get the development environment set up, install the following tools.

### Go

Follow online guides to set this up.

### Golds (Auto Documentation)

For automated documentation generation, use [Golds](https://github.com/go101/golds). Install `golds` by running:

```
go install go101.org/golds@latest
```

### Add Go Binaries to your Path

Golang binaries are usually installed in `~/go/bin`. You can check this exists if you install Golds (above). 

If you are running `zsh`, you can add this to your path by running the following:

```
echo 'export PATH=$PATH:~/go/bin' >> ~/.zshrc
source ~/.zshrc
```

## Code Structure

The structure of this repo is heavily based on the [example Golang repository](https://github.com/golang-standards/project-layout/tree/master) structure by [golang-standards](https://github.com/golang-standards).

Most code is in the `/internal/` directory. These are the main different layers that are used to set up the separation of concerns.

| layer    | description                                                                                                     |
|----------|-----------------------------------------------------------------------------------------------------------------|
| `common` | Core utility functions. Often these will be third party packages wrapped in my own code to avoid vendor lock in |
| `graph`  | General graph data structures and path finding algorithms.                                                      |



