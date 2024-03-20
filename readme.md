# htmgenie

htmgenie is a command-line interface (CLI) tool designed to convert Markdown files to their HTML equivalents.

## Getting Started

### Prerequisites

Before you begin, make sure you have the following prerequisites installed:

#### [Go](https://golang.org/doc/install)

You'll need Go version 1.22 or higher for compiling the CLI and managing other dependencies.

#### Install Make

- **Windows**: Download and install [make](https://gnuwin32.sourceforge.net/packages/make.htm). Alternatively, if you have [Chocolatey](https://chocolatey.org/install) installed, you can run the following command in an admin command prompt:

  ```shell
  choco install make
  ```

- **Mac**: Use Homebrew to install make by running the following command in the terminal:

  ```shell
  brew install make
  ```

#### Set Environment Variables

Make sure you set the following environment variable to customize the output path for generated html file:

- `HTMGENIE_OPATH`: Override the default path set to current directory `.`

### Installation

To install the CLI, follow these steps:

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Run the following commands:

```shell
make build
make install
```

### Using the CLI

Once the CLI is installed, you can use the following syntax to generate HTML from Markdown files:

```shell
htmgenie generate -f <filename_with_path>
```

This command converts a Markdown file to its HTML equivalent and saves the result to a new file. The tool supports a subset of Markdown syntax and generates HTML tags accordingly.

### Command Options

The `generate` command supports the following flags:

- `-f, --file`: Specifies the input Markdown file to be converted to HTML. Only `.md` files are accepted.

For more information, you can use the `--help` flag with the `generate` command:

```shell
htmgenie generate --help
```
