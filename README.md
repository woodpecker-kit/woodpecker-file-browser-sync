[![ci](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/workflows/ci/badge.svg)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-file-browser-sync?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-file-browser-sync?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-file-browser-sync)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-file-browser-sync)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-file-browser-sync)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-file-browser-sync)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync)
[![codecov](https://codecov.io/gh/woodpecker-kit/woodpecker-file-browser-sync/branch/main/graph/badge.svg)](https://codecov.io/gh/woodpecker-kit/woodpecker-file-browser-sync)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-file-browser-sync)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-file-browser-sync)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/releases)

## for what

- file sync for woodpecker-ci use https://github.com/filebrowser/filebrowser

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/woodpecker-kit/woodpecker-file-browser-sync)](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [x] sync as [github.com/filebrowser/filebrowser](https://github.com/filebrowser/filebrowser)
    - local path will sync to remote path graph `{.sync}/${repo-Hostname}/${repo-Owner}/${repo-Name}/{build-number}`

- [x] support link choose by web test
    - [x] `file-browser-urls` support multi urls, will auto switch host fast
    - [x] `file-browser-standby-url` if multi urls not work, will use this

- [x] support `dry-run` mode
- [x] support sync mode `download` and `upload`
    - each mode support `sync-include-globs` and `sync-exclude-globs`

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

### workflow usage

- see [doc](doc/docs.md)

## Notice

- want dev this project, see [dev doc](doc/README.md)