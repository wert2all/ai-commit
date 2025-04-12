# Changelog

## [1.4.0](https://github.com/wert2all/ai-commit/compare/v1.3.0...v1.4.0) (2025-04-12)


### Features

* add method to retrieve content of changed files and add it to ([0519ad1](https://github.com/wert2all/ai-commit/commit/0519ad1798a01b367c51d466ae01dcc472236350))
* add OpenRoute support ([97a87ae](https://github.com/wert2all/ai-commit/commit/97a87ae9d83c6ff40d24690914b7bb5e5cea65d9))
* **ai:** update default models ([a4ff65a](https://github.com/wert2all/ai-commit/commit/a4ff65a66beeeff61b7c02e19b6740a8c68f4be4))
* **config:** add `--version` option and display version ([c013e77](https://github.com/wert2all/ai-commit/commit/c013e77608641e91dda3ef92ed60f08c0d76356e))
* **project:** add option to include changed files content in context generation ([b6d4fd6](https://github.com/wert2all/ai-commit/commit/b6d4fd6ad8b37e1e059cae2f79d7a152a2bac8e5))


### Bug Fixes

* update system prompt ([12977b8](https://github.com/wert2all/ai-commit/commit/12977b8487405094807e878ab5ac1e97dc25d2b7))

## [1.3.0](https://github.com/wert2all/ai-commit/compare/v1.2.0...v1.3.0) (2025-04-11)


### Features

* add Claude support ([6ecfafa](https://github.com/wert2all/ai-commit/commit/6ecfafa3a48107d655e9ae8bba2bebcb33dc1cdf))
* add option for commit control ([3c57bf2](https://github.com/wert2all/ai-commit/commit/3c57bf2250e6d28ecde0b6d9add2575f7dd5afb9))
* **ai:** add local AI provider with Ollama support ([c06ce48](https://github.com/wert2all/ai-commit/commit/c06ce483c0064954119923c77d7392657990c041))
* enable Gemini provider ([927a4a6](https://github.com/wert2all/ai-commit/commit/927a4a6324bacfdacbd380312164ca1e02d93a58))
* **ui:** add card display for generated commit message ([c6bc3f3](https://github.com/wert2all/ai-commit/commit/c6bc3f3613615dc88bd08924a3b01693918f81b8))


### Bug Fixes

* **error:** how error using tui library ([bb1ce78](https://github.com/wert2all/ai-commit/commit/bb1ce78ca6d260a7d6d7e8603527a9eaf7b9082b))
* **openai:** use project context system prompt in commit message generation ([c0fb134](https://github.com/wert2all/ai-commit/commit/c0fb134ef6a092be5d37f9c7fbe69160ff08f8a9))
* remove unnecessary error handling in NewLocalProvider function ([d55eb7d](https://github.com/wert2all/ai-commit/commit/d55eb7ddecd46107a27e7117ec1e05a546766ba4))
* remove unnecessary SystemPrompt and GenerateCommitMessagePrompt functions ([8588731](https://github.com/wert2all/ai-commit/commit/8588731fad2e1d4ae1cb269cc20f837456d08f42))
* update README and clean up local ollama implementation ([448f893](https://github.com/wert2all/ai-commit/commit/448f893bbbeac6f2658f02abf95d0feee36d37bd))

## [1.2.0](https://github.com/wert2all/ai-commit/compare/v1.1.0...v1.2.0) (2025-04-09)


### Features

* **commit:** allow 'enter' as yes for commit changes ([ecbd331](https://github.com/wert2all/ai-commit/commit/ecbd331c48f77cd63af0835a826891b8c63a6692))

## [1.1.0](https://github.com/wert2all/ai-commit/compare/v1.0.0...v1.1.0) (2025-04-09)


### Features

* add firebase studio support ([d73e9cd](https://github.com/wert2all/ai-commit/commit/d73e9cddb37ee9b57fc81bb2cc2b98e723c8f112))
* add languages detection and context builder improvements ([90c0535](https://github.com/wert2all/ai-commit/commit/90c053523681e5cb168feafcbf1f53ad39add387))
* add project directory parameter ([1e47473](https://github.com/wert2all/ai-commit/commit/1e47473410da473076c7af0f2c4c9d84c612d580))
* add support for multiple AI providers ([0137e0d](https://github.com/wert2all/ai-commit/commit/0137e0dd48df1ae44c85dd4e50a40841965cddfc))
* **ai/provider:** update GenerateCommitMessage to accept Changes object instead of strings ([67f0fa3](https://github.com/wert2all/ai-commit/commit/67f0fa334123680754c3599a19ba42195d7705d5))
* **ai:** add Claude provider support ([938b7a2](https://github.com/wert2all/ai-commit/commit/938b7a2a9d142fb4db79b207d71e5b4e32291815))
* **ai:** add Google Gemini support ([8a540f7](https://github.com/wert2all/ai-commit/commit/8a540f7a89985e0336d1ec3db379dde5fa13d1f8))
* **ai:** add Mistral AI provider support ([e17d05b](https://github.com/wert2all/ai-commit/commit/e17d05b157acd43ee50eff0159d7b4173a18d78a))
* **config:** add .env file support ([0db5602](https://github.com/wert2all/ai-commit/commit/0db560294aec5c9f835b9fbe1b69e17ae6180f0b))
* context builder for build project context ([d717223](https://github.com/wert2all/ai-commit/commit/d7172230dadb79e5c8144d1891df24cae889571a))
* **context:** add git branch information to project context ([4279a16](https://github.com/wert2all/ai-commit/commit/4279a16a01769fc1838bad1887dcbd5d3d58aa8e))
* **main:** add commit with generated message ([75e931a](https://github.com/wert2all/ai-commit/commit/75e931a35f073d29612401f71a0350fe4741c8ac))
* **main:** add debug flag and logging to main function and generateCommitMessage function ([3d56c0b](https://github.com/wert2all/ai-commit/commit/3d56c0b5f7bdf57f9971643bb08c101d2abb4423))
* **project/context:** add project structure to context output ([09fa744](https://github.com/wert2all/ai-commit/commit/09fa744df6823bb69b05b1bb079213d2db039de3))
* **project/context:** add system prompt and ProjectContext struct, update Build method signature and return value ([48b7117](https://github.com/wert2all/ai-commit/commit/48b711730e4d7eaf9737ef3af50e5e1b529ddceb))
* update OpenAI provider ([9b2a082](https://github.com/wert2all/ai-commit/commit/9b2a08271cfce32510fffbd3d85669cc3b317177))


### Bug Fixes

* add .commitlintrc.json ([4296a2a](https://github.com/wert2all/ai-commit/commit/4296a2ab61a655ae9821eeb0b29516295c77a8b2))
* **changes:** add error handling for no changes detected in the repository ([4f5814e](https://github.com/wert2all/ai-commit/commit/4f5814e2ce761a91796a7f12ade3036118f6efd9))
* fix commit lint on cicd ([bf2aa5b](https://github.com/wert2all/ai-commit/commit/bf2aa5bf7594b1e6b2b97e4ef9a6fcd3fb6c14f6))
* fix release build ([1b8f09f](https://github.com/wert2all/ai-commit/commit/1b8f09f2a17c71c4489a26bd27715b9255e93679))
* fix release build ([f0e82e0](https://github.com/wert2all/ai-commit/commit/f0e82e0688cc85633c77e301162818377d25a634))
* resolve context conflicts and improve project structure analysis ([6ff4906](https://github.com/wert2all/ai-commit/commit/6ff4906769b6dbd774ce4270ca9c69d1e5cdb4c0))
* temporaly remove other providers exept mistral ([9f18c16](https://github.com/wert2all/ai-commit/commit/9f18c16e42444b81a05eea94259d68a1ece3f344))
* update import path and fix debug logging ([089d556](https://github.com/wert2all/ai-commit/commit/089d556342aa4d60bc085d1f0698dae10f9c0381))
