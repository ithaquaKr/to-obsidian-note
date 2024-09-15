# to-obsidian-note

`to-obsidian-note` is a simple Go-based CLI tool that reads content from `stdin` and generates a Markdown note file based on a predefined template. It allows users to specify the name of the file, tags, and the directory where the file will be saved. The tool also automatically adds a creation timestamp to the note.

## Features

- Read note content from `stdin`.
- Specify the output file name using the `--name` flag.
- Add tags to the note using the `--tag` flag.
- Automatically embeds the current date and time into the note template.
- Saves the generated file to a directory specified by the `OBS_PATH` environment variable.

## Template Structure

The generated Markdown note file follows this structure:

```markdown
---
id: <note-name>
aliases:
  - <note-name>
tags:
  - <tag>
time: <timestamp>
---

<Content from stdin>
```

## Requirements

- Go (Version 1.16 or later)

## Install

## Usage

### Basic usage

### Flags

- `--name`:
- `--tag`

## Environment variable

- `OBS_PATH`: Specifies the directory where the generated Markdown file will be saved. Ensure this is set before running the tool.

Example

```bash
OBS_PATH="${HOME}/my-obsidian/inbox/"
```

## Contribution

Feel free to fork this repository, open issues, or create pull requests for improvements and bug fixes.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

