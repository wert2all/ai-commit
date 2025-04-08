# AI Commit Message Generator

A Go application that generates contextual conventional commit messages using OpenAI's GPT model by analyzing your git changes.

## Prerequisites

1. You need an OpenAI API key. You can get one from [OpenAI's website](https://platform.openai.com/)
2. Git must be installed and the application must be run from within a git repository
3. Set your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY='your-api-key-here'
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
go run main.go
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

- Uses OpenAI's GPT-3.5-turbo model for intelligent commit message generation
- Analyzes your actual git changes to generate contextual commit messages
- Considers staged changes, unstaged changes, and untracked files
- Follows [Conventional Commits](https://www.conventionalcommits.org/) specification
- Generates precise and meaningful commit messages based on your actual code changes
