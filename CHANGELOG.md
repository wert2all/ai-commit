# Changelog

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
