# AI Commit Message Generator

A Go application that generates contextual conventional commit messages using AI (OpenAI or others) by analyzing your git changes.

## Features

- Supports multiple AI providers:
  - OpenAI (GPT-3.5-turbo, GPT-4)
  - Claude (claude-2, claude-instant-1)
  - Mistral AI (mistral-medium, mistral-small, mistral-tiny)
  - Google Gemini (gemini-pro, gemini-pro-vision)
  - OpenRouter (with access to free and paid models)
  - **Local AI (Ollama)**
- Analyzes your actual git changes to generate contextual commit messages
- Considers staged changes
- Follows [Conventional Commits](https://www.conventionalcommits.org/) specification
- Generates precise and meaningful commit messages based on your actual code changes
- Commit changes with generated message

## Supported AI Providers

| Provider     | Default Model             |
| ------------ | ------------------------- |
| OpenAI       | gpt-3.5-turbo             |
| Claude       | claude-3-7-sonnet-latest  |
| Mistral      | codestral-latest          |
| Gemini       | gemini-2.0-flash          |
| OpenRouter   | openrouter/optimus-alpha  |
| Local Ollama | no default (must specify) |

## Options

| Option             | Description                                                                  |
| ------------------ | ---------------------------------------------------------------------------- |
| `--provider`       | Specify the AI provider (openai, claude, mistral, gemini, openrouter, local) |
| `--model`          | Specify the model to use with the selected provider                          |
| `--endpoint`       | Custom API endpoint URL (useful for local deployments)                       |
|                    |                                                                              |
| `--without-commit` | Generate a commit message without committing changes                         |

## Prerequisites

1. You need an API key from one of the supported providers:
   - [OpenAI](https://platform.openai.com/)
   - [Mistral AI](https://mistral.ai/)
   - [Google AI (Gemini)](https://ai.google.dev/)
   - [Anthropic (Claude)](https://www.anthropic.com/)
   - [OpenRouter](https://openrouter.ai)
   - **Local Ollama**
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

   # For OpenRouter
   export OPENROUTER_API_KEY='your-api-key-here'
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

# Use OpenRouter
go run main.go --provider openrouter --model optimus-alpha

# Use Local AI (Ollama)
go run main.go --provider local --model llama2
```

The program will analyze your current git changes and generate an AI-powered commit message following the conventional commit format:
`type(scope): description`

For example:

- `feat(auth): implement OAuth2 authentication flow`
- `fix(api): resolve race condition in database connection pool`
- `docs(readme): update installation instructions`

## OpenRouter Setup

OpenRouter provides access to various AI models, including some free ones that don't require billing information:

1. Sign up at [OpenRouter](https://openrouter.ai)
2. Navigate to your account settings to generate an API key
3. To use a free model, specify one of the available free models:

```bash
# Use a free OpenRouter model
go run main.go --provider openrouter --model openrouter/auto

# Other free tiers may include:
go run main.go --provider openrouter --model mistralai/mistral-7b-instruct
go run main.go --provider openrouter --model openchat/openchat-7b
```

The `openrouter/auto` model will automatically route to the best available free model.

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
