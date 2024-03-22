# Design Document

## Overview

htmgenie is a command-line interface (CLI) tool designed to convert Markdown files to their HTML equivalents. This document outlines the design considerations, architecture, and implementation details of htmgenie.

## Goals

The primary goal of htmgenie is to provide developers with a simple and efficient way to convert Markdown files to HTML format. Key objectives include:

- Support a subset of Markdown syntax for conversion.
- Generate HTML tags corresponding to Markdown elements.
- Provide a user-friendly CLI interface for ease of use.

## Non goals

- Build a sophisticated html generator that supports all markdown formatting for a given markdown file
- Hard target on the performance

## In Scope

| Markdown                               | HTML                                                 |
| -------------------------------------- | -------------------------------------------------    |
| `# Heading 1`                          | `<h1>Heading 1</h1>`                                 |
| `## Heading 2`                         | `<h2>Heading 2</h2>`                                 |
| `### Heading 3`                        | `<h3>Heading 3</h3>`                                 |
| `#### Heading 4`                       | `<h4>Heading 4</h4>`                                 |
| `##### Heading 5`                      | `<h5>Heading 5</h5>`                                 |
| `###### Heading 6`                     | `<h6>Heading 6</h6>`                                 |
| `Unformatted text`                     | `<p>Unformatted text</p>`                            |
| `[Link text](https://www.example.com)` | `<a href="https://www.example.com">Link text</a>`    |
| `Blank line`                           | `Ignored`                                            |

## Assumptions

- There must be a **new-line** separating each type of element except for _links_.
  - This means that, if there is unformatted text, it must be ended with an empty line. For the elements currently in scope, if a heading is present at the end of unformatted text without an empty line, it is considered to be part of the paragraph/unformatted text
- Empty spaces are considered as an empty line
- treated as invalid url formats:
  - `[[link]](https://www.more.example.com)`
  - `[link](www.google.com)`
  - More cases are covered in [url unit tests](./pkg/mapper/url_test.go)

## Implementation Decisions

- Is there a desired size limit for the input markdown file? - _Limiting to 5MB size_
- Is there a desired performance expectation in generating the html file? - _No specific goal set._
- Is there a future scope/goal set for this application? This could help in building the command structure. - _Keeping options open to extend the CLI to include actions such as validation, selective update, etc..,_
- What should be the behavior when:
  - url link in the markdown is invalid - _Going with valid URL format. Everything else is treated as plain text_
  - potentially harmful strings are found in the markdown file such as embedded scripts. _Keeping it plain for this implementation considering it is a locally hosted application posing minimal risk_
  - markdown content not listed in the formatting table - _Treated as plain text_
  - file provided in the input is invalid due to various reasons including non md files, non-existant file path, access denied files, no file path provided. - _Handled in the validation to exit with an error_
- There is no mention of forming a proposal/design document in the description. Should this not be included in the repository? - _Minimal documentation added for instructions._

## Architecture

htmgenie follows a modular architecture comprising the following components:

1. **CLI Interface**:
    - Responsible for parsing user input, executing commands, and displaying useful logs.
    - Handles file operations such as reading input Markdown files and writing output HTML files.
2. **Markdown Parser**:
    - Parses the input Markdown file and extracts Markdown elements.
    - Converts Markdown elements to HTML tags based on predefined rules.

## Workflow

1. The user invokes the `generate` command via the CLI and specifies the input Markdown file.
2. The CLI interface validates the user input and triggers the Markdown parsing process.
3. The CLI command reads the input Markdown file, extracts Markdown elements, and passes them to the Markdown parser.
4. The Markdown parser converts Markdown elements to HTML tags according to predefined rules.
5. The Markdown parser adds the converted content to the stream writer.
6. The resulting HTML content is written to a new file in the configured output path with the same name and appropriate extension.

## Use Cases

- `-f` parameter
  - parameter not provided in the command [error]
  - leads to an unknown path [error]
  - file supplied is not a markdown file (!`md`) [error]
  - no read access on the file [error]
  - A valid file path provided [success]
- Content
  - Empty file [valid]
  - file with content [valid]
  - Unreadable file (such as binary files) [error]
  - Valid markdown content but not included in the scope. [valid] - only recognized patterns are handled.

## Trade-offs

1. Use of `io.Reader` and `io.Writer` in the Parser.
    - This provided two benefits:
        - Separation of concern where file operations are not of Parser's responsibility
        - This provides an option to extend the implementation easily to support other `io` types such as: reading data from `stdin` and/or writing to multiple streams (`io.MultiWriter()`) such as printing on `stdout` along with writing to a file.
2. Use of Mapper `map` to handle direct/key based replacement such as headings. - Simpler/easier mapping framework for multiple replacement functions
3. Use of `cobra` instead of custom/handwritten framework 
    - cobra is a widely adopted CLI framework
    - It provides a lot of customization options wrt configuration, managing arguments, generating command help
    - strong community support
4. Simple implementation for `Chunk`
    - There is only one use case to support at this time `<p></p>`. However, pattern is modifiable.
5. More tests for mapper functions that get narrower as it reaches `generate` (or command implementation)
    - Due to limited time available, I chose to make the individual component fool proof where majority of cases reside.

## Future Enhancements

### Features

- Support for Markdown syntax elements currently not in scope.
- Introduce config options to handle parameters such as:
  - Override the default file size limit (5MB). The value should be denoted in bytes.
  - Write to stdout instead of to a file
  - Support inline input
- Validate markdown file before starting to generate html.

### Refinements

- Add few more unit tests to `generate` and `parser` implementation.
- Use interface to extend `Chunk` implementation to handle different complex elements.
- Introduce `--debug` arguement to handle `log` prints on stdout.
  - Suppliment it with configurable user friendly text (and emojis?).
- Write function help to describe each public function and its signature
- CI/CD pipeline
- integrate `tilt` and `envrc` for developer productivity
- GitHub `dependabot` to automate dependency updates
