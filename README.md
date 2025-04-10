# AI Commit Message Generator

A Go application that generates contextual conventional commit messages using AI (OpenAI or others) by analyzing your git changes.

## Prerequisites

1. You need an API key from one of the supported providers:
   - [OpenAI](https://platform.openai.com/)
   - [Mistral AI](https://mistral.ai/)
   - [Google AI (Gemini)](https://ai.google.dev/) (comming soon)
   - **Local Ollama**
   - [Anthropic (Claude)](https://www.anthropic.com/) (comming soon)
2. Git must be installed and the application must be run from within a git repository
3. Set your API key as an environment variable:

   Option 1: Using environment variables

   ```bash
   # For OpenAI
   export OPENAI_API_KEY='your-api-key-here'

   # For Claude
   export CLAUDE_API_KEY='your-api-key-here'

   # For Mistral
   export MISTRAL_API_KEY='your-api-key-here'

   # For Gemini
   export GEMINI_API_KEY='your-api-key-here'
   ```

## Installation

```bash
go mod download
```

## Usage

To generate a commit message:

```bash
# Use OpenAI (default)
go run main.go

# Use OpenAI with specific model
go run main.go --provider openai --model gpt-4

# Use Claude
go run main.go --provider claude --model claude-2

# Use Mistral
go run main.go --provider mistral --model mistral-medium

# Use Gemini
go run main.go --provider gemini --model gemini-pro

# Use Local AI (Ollama)
go run main.go --provider local --model llama2
```

The program will analyze your current git changes and generate an AI-powered commit message following the conventional commit format:
`type(scope): description`

For example:

- `feat(auth): implement OAuth2 authentication flow`
- `fix(api): resolve race condition in database connection pool`
- `docs(readme): update installation instructions`

## Supported AI Providers

- OpenAI
- Mistral
- Gemini
- Local Ollama

## Local AI Provider Setup

To use a local Ollama:

1. Install [Ollama](https://ollama.com/)
2. Pull a model (e.g., `ollama pull llama2`)
3. Start the Ollama server
4. Run the tool with:

```bash
go run main.go --provider local --model llama2 -endpoint http://localhost:11434
```

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
  - Mistral AI (mistral-medium, mistral-small, mistral-tiny)
  - Google Gemini (gemini-pro, gemini-pro-vision)
  - **Local AI (Ollama)**
- Analyzes your actual git changes to generate contextual commit messages
- Considers staged changes
- Follows [Conventional Commits](https://www.conventionalcommits.org/) specification
- Generates precise and meaningful commit messages based on your actual code changes
