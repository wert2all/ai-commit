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

| Option                 | Description                                                                  |
| ---------------------- | ---------------------------------------------------------------------------- |
| `--provider`           | Specify the AI provider (openai, claude, mistral, gemini, openrouter, local) |
| `--model`              | Specify the model to use with the selected provider                          |
| `--endpoint`           | Custom API endpoint URL (useful for local deployments)                       |
|                        |                                                                              |
| `--without-commit`     | Generate a commit message without committing changes                         |
| `--with-files-content` | Append content of changes files to context                                   |
|                        |                                                                              |
| `--version`            | Show application version                                                     |

## Prerequisites

1. You need an API key from one of the supported providers:
   - [OpenAI](https://platform.openai.com/)
   - [Mistral AI](https://mistral.ai/)
   - [Google AI (Gemini)](https://ai.google.dev/)
   - [Anthropic (Claude)](https://www.anthropic.com/)
   - [OpenRouter](https://openrouter.ai)
   - **Local Ollama**
2. Set your API key as an environment variable:

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

1. Download the latest release from the [GitHub Releases page](https://github.com/wert2all/ai-commit/releases).
2. Set your API key as an environment variable for your chosen provider (see "Prerequisites" above).
3. Run the application with your selected provider, for example:

   ```bash
   ./ai-commit --provider openai
   ```

## Usage

To generate a commit message:

```bash
# Use OpenAI (default)
./ai-commit

# Use OpenAI with specific model
./ai-commit --provider openai --model gpt-4

# Use Claude
./ai-commit --provider claude --model claude-2

# Use Mistral
./ai-commit --provider mistral --model mistral-medium

# Use Gemini
./ai-commit --provider gemini --model gemini-pro

# Use OpenRouter
./ai-commit --provider openrouter --model optimus-alpha

# Use Local AI (Ollama)
./ai-commit --provider local --model llama2
```

The program will analyze your current git changes and generate an AI-powered commit message following the conventional commit format:
`type(scope): description`

For example:

- `feat(auth): implement OAuth2 authentication flow`
- `fix(api): resolve race condition in database connection pool`
- `docs(readme): update installation instructions`

## OpenRouter Setup

OpenRouter provides access to various AI models. By default, the tool uses OpenRouter's default model, but you can also specify a model if desired.

1. Sign up at [OpenRouter](https://openrouter.ai)
2. Navigate to your account settings to generate an API key
3. Run the tool with OpenRouter as the provider (default model):

   ```bash
   ./ai-commit --provider openrouter
   ```

   Or specify a model explicitly:

   ```bash
   ./ai-commit --provider openrouter --model mistralai/mistral-7b-instruct
   ```

## Local AI Provider Setup

To use a local Ollama:

1. Install [Ollama](https://ollama.com/)
2. Pull a model (e.g., `ollama pull llama2`)
3. Start the Ollama server
4. Run the tool with:

```bash
./ai-commit --provider local --model llama2 -endpoint http://localhost:11434
```
