# codectl

Inspired by [boiler](https://github.com/tmrts/boilr), codectl is a templating tool powered by Go templates. It is designed to generate boilerplate starting points for new projects based on a predefined template.

## Features
* **Go Templates** - Built around [Go templates](https://golang.org/pkg/text/template/), a powerful data-driven templating engine
* **Repositories** - Supports template repositories by pulling templates from Github repositories and storing them locally
* **Parameter Definitions** - Provides powerful template definitions that allow for specifying parameters to prompt for when generating

## Installation
Currently, codectl can be installed by using `go get`:

```bash
$ go get -u github.com/thmhoag/codectl/cmd/codectl
```

In the future, codectl will be available as a self-contained binary via Github releases.

## Getting started
For a list of commands, you can start with `codectl --help`. 

### Adding a Template Repository
Currently, codectl supports Github repositories as template repositories.

To add a new repository, simply provide a repository name and a fully qualified URL to the Github repo:

```bash
$ codectl repo add default https://github.com/thmhoag/codectl-templates
```

Additionally, you can specify a specific folder of the repository by using `//` to denote that it is a folder off the root:

```bash
$ codectl repo add default https://github.com/thmhoag/codectl//templates
```

And now you can list the repos to verify it has been added:

```bash
$ codectl repo ls
```

### Updating Template Repositories
Template repositories are downloaded and cached from their sources. Once a repository has been added, the cache must be updated to reflect the change:

```bash
$ codectl repo update
```

**NOTE**: At this time, codectl only supports Github.com. In the future it will support Github Enterprise, as well as other remote sources.

### Generating from a Template

List available templates:
```bash
$ codectl template ls
```

Choose a template name, then generate it:
```bash
$ codectl generate default.go.cli
```

codectl should now prompt for any required parameters for the template chosen, and generate the output to the current directory.

To specify a different output directory, use the output flag:
```bash
$ codectl generate default.go.cli -o ./my-desired-output
```

