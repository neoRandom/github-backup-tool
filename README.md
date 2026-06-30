# GitHub Backup Tool

A small Go CLI that backs up your GitHub repositories by cloning each repo in two formats:

- a **mirror clone** (`--mirror`) for complete Git history and refs
- a **standard clone** for easy browsing of files

## How it works

1. Reads `GITHUB_TOKEN` from a local `.env` file.
2. Uses the GitHub API to get your username and repository list.
3. Clones each repository under `./backup/<owner>/<repo>/` as:
   - `<repo>.git` (mirror clone)
   - `<repo>` (normal clone)
4. Prompts to move the backup into a timestamped folder under `./snapshots/`.

## Requirements

- Go 1.26+
- Git installed and available on `PATH`
- A GitHub personal access token in `.env`

## Setup

1. Copy the example env file:

   ```bash
   cp .env.example .env
   ```

2. Edit `.env` and set your token:

   ```env
   GITHUB_TOKEN=your_token_here
   ```

## Usage

Run the tool from the repository root:

```bash
go run .
```

After cloning, confirm the prompt to move the backup to a snapshot folder.

## Development

Build:

```bash
go build ./...
```

Test:

```bash
go test ./...
```

## Important notes

- Keep your `.env` file private and never commit it.
- The token is embedded in clone URLs at runtime for authentication.
- Large accounts may take time to back up depending on repository count and size.
