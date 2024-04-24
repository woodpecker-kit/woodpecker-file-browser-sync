[![ci](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/workflows/ci/badge.svg)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-file-browser-sync?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-file-browser-sync?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-file-browser-sync)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-file-browser-sync)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-file-browser-sync)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-file-browser-sync)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync)
[![codecov](https://codecov.io/gh/woodpecker-kit/woodpecker-file-browser-sync/branch/main/graph/badge.svg)](https://codecov.io/gh/woodpecker-kit/woodpecker-file-browser-sync)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-file-browser-sync)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-file-browser-sync)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/releases)

## for what

- this project used to woodpecker plugin

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/woodpecker-kit/woodpecker-file-browser-sync)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- use this template, replace list below and add usage
    - `github.com/woodpecker-kit/woodpecker-file-browser-sync` to your package name
    - `woodpecker-kit` to your owner name
    - `woodpecker-file-browser-sync` to your project name

- use github action for this workflow push to docker hub, must add
    - variables `ENV_DOCKERHUB_OWNER` user of docker hub
    - variables `ENV_DOCKERHUB_REPO_NAME` repo name of docker hub
    - secrets `DOCKERHUB_TOKEN` token of docker hub user from [hub.docker](https://hub.docker.com/settings/security)

- check `docker-bake.hcl` config, change to your docker image

- if you use `wd_steps_transfer` just add `.woodpecker_kit.steps.transfer` at git ignore
- change code start with `// change or remove`

### workflow usage

- see [doc](doc/docs.md)

## Notice

- want dev this project, see [dev doc](doc/README.md)