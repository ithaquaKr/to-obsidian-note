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

> TBD

## Usage

### Basic usage

To create a new note, run the `to-obsidian-note` command with the following options:

```bash
to-obsidian-note -n "note" -t "ai"
```

After running the command, you will be prompted to enter the content of your note. Once you're done, press `Ctrl + D` (On Linux/macOS) or `Ctrl + Z` (On Windows) to finish and generate the note.

Example

```bash
to-obsidian-note --name "prompt-engineering" --tag "ai,prompt"
```

After entering the note content and confirming, the following file will be generated.

```markdown
---
id: prompt-engineering
aliases:
  - prompt-engineering
tags:
  - ai
  - prompt
time: 2024-09-15-Mon 10:30:00
---

This is a note about prompt engineering.
```

### Advance usage

```bash
echo "This is content of note" | to_obsidian_note -n "prompt-engineering" -t "ai,prompt"
```

```bash
pbpaste | to_obsidian_note -n "take-note" -t "note"
```

```bash
echo "Explain this content" | fabric -p ai | to_obsidian_note -n "explain" -t "ai"
```

### Flags

- `--name`:
- `--tag`

## Environment variable

- `OBS_PATH`: Specifies the directory where the generated Markdown file will be saved. Ensure this is set before running the tool.

Example

```bash
OBS_PATH="User/your_user/obsidian/inbox/"
```

## Contribution

Feel free to fork this repository, open issues, or create pull requests for improvements and bug fixes.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

