# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.2.0](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/compare/1.1.0...v1.2.0) (2025-02-06)

### ✨ Features

* add support for multiple docker bake targets ([2cfc5320](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/2cfc53206cb65cefbf1e1a90952e72bb22dd9f84)), feat [#28](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/issues/28)

### 👷‍ Build System

* comment out zymosis installation and execution ([c8503100](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/c8503100ff5169061e3d36c9a18adb5270c2e519))

* update docker-bake.hcl for image building ([0e4a61fb](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/0e4a61fb0ba200ce460319f7b31aeed21648b0c8))

* pin zymosis version to v1.1.3 ([56587cea](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/56587cea6ff1654c328eb58f225522dd229b0e70))

* bump github.com/woodpecker-kit/woodpecker-tools ([5ef266ba](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/5ef266ba98dab56bd38758d63ebe06f613d21e19))

## [1.1.0](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/compare/1.0.0...v1.1.0) (2024-12-21)

### ✨ Features

* update Go version to 1.21 ([ae4721b6](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/ae4721b62b90d94bec6233939d8f1fd5919c15c1)), re [#23](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/issues/23)

### 📝 Documentation

* update contributing guidelines and code of conduct ([71f418c0](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/71f418c08038b12583ff4b8a7e6a2c533a87b943))

### 👷‍ Build System

* bump github.com/woodpecker-kit/woodpecker-tools ([eff952c8](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/eff952c81b514355c0e2045007fce94c6d4f5730))

* update GitHub release action to v2 ([d95740c0](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/d95740c0b9e9ae14d0aa9b35fb64e55625edce6d))

* bump github.com/Masterminds/semver/v3 from 3.3.0 to 3.3.1 ([8dd23fee](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/8dd23fee0b8200bab111d6e6e8f5810aa74757d5))

* bump github.com/urfave/cli/v2 from 2.27.4 to 2.27.5 ([80df0e87](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/80df0e8769990bd018518bca8b4da58dd4f97136))

* bump github.com/sinlov-go/unittest-kit from 1.1.1 to 1.2.1 ([3e4d2939](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/3e4d29397bd2cd9cfe1c85321b31f88709d0fcf1))

* bump github.com/Masterminds/semver/v3 from 3.2.1 to 3.3.0 ([1236dbb1](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/1236dbb1326ac97c409685d8ec03877f7cb73e2a))

* bump github.com/sinlov-go/go-common-lib from 1.7.0 to 1.7.1 ([15e7f703](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/15e7f70391e062446f61afbbbdda31d13c6638e2))

* bump github.com/sinlov-go/unittest-kit from 1.1.0 to 1.1.1 ([69c4f9ca](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/69c4f9caa6a978de44cd4050aed751d9e8981a78))

* bump github.com/urfave/cli/v2 from 2.27.3 to 2.27.4 ([cef4e778](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/cef4e7785da61ad3f9a3f5eb709f26f150f2d754))

## 1.0.0 (2024-04-26)

### 🐛 Bug Fixes

* github.com/sinlov/filebrowser-client v0.7.2 and fix usage doc ([59324d8f](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/59324d8f248fd8c6e8e421b56b1afe1eba4ccd80))

* fix send file will be error ([c99b7bd3](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/c99b7bd381e84922d62f5947d90e52c5a8c2a947)), fix [#5](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/issues/5)

### ✨ Features

* filebrowser client mode `download` and `upload` ([295c1ea6](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/295c1ea6683988b68c786f5544959771fc87e249)), feat [#1](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/issues/1)

* finish `doSyncModeUpload` and pass test case ([27a9ac26](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/27a9ac2634e3362f9ab627210987194bbd594fea))

* add basic flag and test case for plugin ([f9061304](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/f90613048a17638f9f5bb0841cf105be432e7c65))

* `HttpLinkSpeed` can test as http link `BestLinkIgnoreRetry` can fast find out best ([5f848586](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/5f8485865c233d14d339975ef87099c12a209424))

### 📝 Documentation

* add version badge at doc/docs.md ([fadf91ae](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/fadf91aeb5cb46e62c7ba8d326f8d0f4a9798243))

* add full usage of this plugin ([8e596151](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/commit/8e596151f173a8ab3d1892860dea74e381cc434c)), feat [#3](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/issues/3)
