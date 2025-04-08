# AI Commit Message Generator

A Go application that generates contextual conventional commit messages using AI (OpenAI or Claude) by analyzing your git changes.

## Prerequisites

1. You need an API key from one of the supported providers:
   - [OpenAI](https://platform.openai.com/)
   - [Anthropic (Claude)](https://www.anthropic.com/)
2. Git must be installed and the application must be run from within a git repository
3. Set your API key as an environment variable:

```bash
# For OpenAI
export OPENAI_API_KEY='your-api-key-here'

# For Claude
export CLAUDE_API_KEY='your-api-key-here'
```

## Installation

```bash
go mod download
```

## Usage

The application should be run from within a git repository. It will analyze:
- Staged changes (`git diff --cached`)
- Unstaged changes (`git diff`)
- Untracked files

To generate a commit message:

```bash
# Use OpenAI (default)
go run main.go

# Use OpenAI with specific model
go run main.go --provider openai --model gpt-4

# Use Claude
go run main.go --provider claude --model claude-2
```

The program will analyze your current git changes and generate an AI-powered commit message following the conventional commit format:
`type(scope): description`

For example:
- `feat(auth): implement OAuth2 authentication flow`
- `fix(api): resolve race condition in database connection pool`
- `docs(readme): update installation instructions`

## Build

To build the application:

```bash
go build -o commit-generator
```

Then you can run the executable:

```bash
./commit-generator
```

## Features

- Supports multiple AI providers:
  - OpenAI (GPT-3.5-turbo, GPT-4)
  - Claude (claude-2, claude-instant-1)
- Analyzes your actual git changes to generate contextual commit messages
- Considers staged changes, unstaged changes, and untracked files
- Follows [Conventional Commits](https://www.conventionalcommits.org/) specification
- Generates precise and meaningful commit messages based on your actual code changes
