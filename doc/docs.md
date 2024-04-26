---
name: woodpecker-file-browser-sync
description: woodpecker file_browser_sync template
author: woodpecker-kit
tags: [ file-browser, woodpecker-file-browser-sync ]
containerImage: sinlov/woodpecker-file-browser-sync
containerImageUrl: https://hub.docker.com/r/sinlov/woodpecker-file-browser-sync
url: https://github.com/woodpecker-kit/woodpecker-file-browser-sync
icon: https://raw.githubusercontent.com/woodpecker-kit/woodpecker-file-browser-sync/main/doc/logo.jpeg
---

# woodpecker-file-browser-sync

- file sync for woodpecker-ci use https://github.com/filebrowser/filebrowser

## Features

- [x] sync as [github.com/filebrowser/filebrowser](https://github.com/filebrowser/filebrowser)
    - local path will sync to remote path graph `{.sync}/${repo-Hostname}/${repo-Owner}/${repo-Name}/{build-number}`

- [x] support link choose by web test
    - [x] `file-browser-urls` support multi urls, will auto switch host fast
    - [x] `file-browser-standby-url` if multi urls not work, will use this

- [x] support `dry-run` mode
- [x] support sync mode `download` and `upload`
    - each mode support `sync-include-globs` and `sync-exclude-globs`

## Settings

| Name                                  | Required | Default value | Description                                                       |
|---------------------------------------|----------|---------------|-------------------------------------------------------------------|
| `debug`                               | **no**   | *false*       | open debug log or open by env `PLUGIN_DEBUG`                      |
| `sync-dry-run`                        | **no**   | *false*       | dry run, only print some info, not real sync                      |
| `file-browser-urls`                   | **yes**  |               | file browser urls, support multi urls, will auto switch host fast |
| `file-browser-username`               | **yes**  |               | file browser username, support from secret                        |
| `file-browser-user-password`          | **yes**  |               | file browser password, support from secret                        |
| `file-browser-standby-url`            | **yes**  |               | file browser standby url, if multi urls not work, will use this   |
| `file-browser-standby-username`       | **yes**  |               | file browser standby username, support from secret                |
| `file-browser-standby-user-passwords` | **yes**  |               | file browser standby password, support from secret                |
| `sync-mode`                           | **yes**  |               | sync mode, support: upload, download                              |
| `sync-work-space-path`                | **yes**  |               | sync path under workspace path                                    |
| `sync-include-globs`                  | **no**   |               | include globs do not use with exclude globs                       |
| `sync-exclude-globs`                  | **no**   |               | exclude globs do not use with include globs                       |

**Hide Settings:**

| Name                                        | Required | Default value                    | Description                                                                      |
|---------------------------------------------|----------|----------------------------------|----------------------------------------------------------------------------------|
| `timeout_second`                            | **no**   | *10*                             | command timeout setting by second, minimum 10s                                   |
| `file-browser-sync-timeout-second`          | **no**   | *60*                             | file browser sync files timeout second, minimum 60s                              |
| `woodpecker-kit-steps-transfer-file-path`   | **no**   | `.woodpecker_kit.steps.transfer` | Steps transfer file path, default by `wd_steps_transfer.DefaultKitStepsFileName` |
| `woodpecker-kit-steps-transfer-disable-out` | **no**   | *false*                          | Steps transfer write disable out                                                 |

## Example

- workflow with backend `docker`

```yml
labels:
  backend: docker
steps:
  woodpecker-file-browser-sync:
    image: sinlov/woodpecker-file-browser-sync:latest
    pull: false
    settings:
      # debug: true
      # sync-dry-run: true # dry run, only print some info, not real sync
      ## connect config
      file-browser-urls:
        - https://filebrowser-inner.example.com
        - https://filebrowser.example.com
        - https://filebrowser-zone.example.com
      file-browser-username:
        from_secret: file_browser_sync_username
      file-browser-user-password:
        from_secret: file_browser_sync_password
      # standby url
      file-browser-standby-url: https://filebrowser-standby.example.com
      file-browser-standby-username:
        from_secret: file_browser_sync_standby_username
      file-browser-standby-user-passwords:
        from_secret: file_browser_sync_standby_password
      ## sync config
      sync-work-space-path: dist # sync path under workspace path, Required
      sync-mode: upload # sync mode, support: upload, download
      sync-include-globs: # include globs do not use with exclude globs
        - "**/**/*.dll"
        - "**/*.md"
```

- workflow with backend `local`, must install at local and effective at evn `PATH`
    - can download by [github release](https://github.com/woodpecker-kit/woodpecker-file-browser-sync/releases)
- install at ${GOPATH}/bin, latest

```bash
go install -a github.com/woodpecker-kit/woodpecker-file-browser-sync/cmd/woodpecker-file-browser-sync@latest
```

- install at ${GOPATH}/bin, v1.0.0

```bash
go install -v github.com/woodpecker-kit/woodpecker-file-browser-sync/cmd/woodpecker-file-browser-sync@v1.0.0
```

```yml
labels:
  backend: local
steps:
  woodpecker-file-browser-sync:
    image: woodpecker-file-browser-sync
    settings:
      # debug: true
      # sync-dry-run: true # dry run, only print some info, not real sync
      ## connect config
      file-browser-urls:
        - https://filebrowser-inner.example.com
        - https://filebrowser.example.com
        - https://filebrowser-zone.example.com
      file-browser-username:
        from_secret: file_browser_sync_username
      file-browser-user-password:
        from_secret: file_browser_sync_password
      # standby url
      file-browser-standby-url: https://filebrowser-standby.example.com
      file-browser-standby-username:
        from_secret: file_browser_sync_standby_username
      file-browser-standby-user-passwords:
        from_secret: file_browser_sync_standby_password
      ## sync config
      sync-work-space-path: dist # sync path under workspace path, Required
      sync-mode: upload # sync mode, support: upload, download
      sync-include-globs: # include globs do not use with exclude globs
        - "**/**/*.dll"
        - "**/*.md"
```

- full config

```yaml
labels:
  backend: docker
steps:
  woodpecker-file-browser-sync:
    image: sinlov/woodpecker-file-browser-sync:latest
    pull: false
    settings:
      debug: true
      sync-dry-run: true # dry run, only print some info, not real sync
      file-browser-sync-timeout-second: 120 # file browser sync files timeout second, minimum 60s
      ## connect config
      file-browser-urls:
        - https://filebrowser-inner.example.com
        - https://filebrowser.example.com
        - https://filebrowser-zone.example.com
      file-browser-username:
        from_secret: file_browser_sync_username
      file-browser-user-password:
        from_secret: file_browser_sync_password
      # standby url
      file-browser-standby-url: https://filebrowser-standby.example.com
      file-browser-standby-username:
        from_secret: file_browser_sync_standby_username
      file-browser-standby-user-passwords:
        from_secret: file_browser_sync_standby_password
      ## sync config
      sync-work-space-path: dist # sync path under workspace path, Required
      sync-mode: upload # sync mode, support: upload, download
      sync-include-globs: # include globs do not use with exclude globs
        - "**/**/*.dll"
        - "**/*.md"
      sync-exclude-globs: # exclude globs do not use with include globs
        - "**/node_modules/**"
        - "**/vendor/**"
```

## Notes

- `file-browser-urls` support multi urls, However, these URLs actually point to the same service, preventing files from
  being downloaded or uploaded normally in the same build.

- `file-browser-standby-url`and related parameters can be set in the server environment
  variable `WOODPECKER_ENVIRONMENT` for use in multiple environments

```ini
WOODPECKER_ENVIRONMENT="PLUGIN_FILE_BROWSER_STANDBY_URL:https://filebrowser-standby.example.com,PLUGIN_FILE_BROWSER_USERNAME:sync-filebrowser,PLUGIN_FILE_BROWSER_STANDBY_USERNAME:sync-filebrowser"
```

then it can be simplified

```yaml
labels:
  backend: docker
steps:
  woodpecker-file-browser-sync:
    image: sinlov/woodpecker-file-browser-sync:latest
    pull: false
    settings:
      # debug: true
      # sync-dry-run: true # dry run, only print some info, not real sync
      ## connect config
      file-browser-urls:
        - https://filebrowser-inner.example.com
        - https://filebrowser.example.com
        - https://filebrowser-zone.example.com
      file-browser-user-password:
        from_secret: file_browser_sync_password
      # standby url
      file-browser-standby-user-passwords:
        from_secret: file_browser_sync_standby_password
      ## sync config
      sync-work-space-path: dist # sync path under workspace path, Required
      sync-mode: upload # sync mode, support: upload, download
      sync-include-globs: # include globs do not use with exclude globs
        - "**/**/*.dll"
        - "**/*.md"
```

## Known limitations

- `file-browser-standby-url` is required, if `file-browser-urls` not work, will use this, but
  if `file-browser-standby-url`
  not work, will not retry.
