---
title: wp-git-action
---

[![Build Status](https://ci.thegeeklab.de/api/badges/thegeeklab/wp-git-action/status.svg)](https://ci.thegeeklab.de/repos/thegeeklab/wp-git-action)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/thegeeklab/wp-git-action)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/wp-git-action)
[![Go Report Card](https://goreportcard.com/badge/github.com/thegeeklab/wp-git-action)](https://goreportcard.com/report/github.com/thegeeklab/wp-git-action)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/wp-git-action)](https://github.com/thegeeklab/wp-git-action/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/thegeeklab/wp-git-action)
[![License: Apache-2.0](https://img.shields.io/github/license/thegeeklab/wp-git-action)](https://github.com/thegeeklab/wp-git-action/blob/main/LICENSE)

Woodpecker CI plugin to perform git actions.

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< toc >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

## Usage

```YAML
steps:
  - name: commit changelog
    image: quay.io/thegeeklab/wp-git-action
    settings:
      action:
        - commit
        - push
      netrc_password: ghp_randomstring
      author_name: octobot
      author_email: octobot@example.com
      message: "[skip ci] update changelog"
```

### Parameters

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< propertylist name=wp-git-action.data sort=name >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

### Examples

#### Publish GitHub pages

The plugin can be used to publish GitHub pages to the pages branch. Remember that the `pages` action cannot be combined with other actions.

```YAML
steps:
  - name: publish
    image: quay.io/thegeeklab/wp-git-action
    settings:
      action:
        - pages
      author_email: bot@texample.com
      author_name: octocat-bot
      message: "update pages"
      branch: gh-pages
      pages_directory: docs/
      netrc_password: ghp_randomstring
```

## Build

Build the binary with the following command:

```Shell
make build
```

Build the Container image with the following command:

```Shell
docker build --file Containerfile.multiarch --tag thegeeklab/wp-git-action .
```

## Test

```Shell
docker run --rm \
  -e PLUGIN_ACTION=clone \
  -e PLUGIN_BRANCH=master \
  -e PLUGIN_AUTHOR_EMAIL=octocat@example.copm \
  -e PLUGIN_AUTHOR_NAME=octocat \
  -e PLUGIN_REMOTE_URL=https://github.com/octocat/Hello-World \
  -w /tmp \
  thegeeklab/wp-git-action
```
