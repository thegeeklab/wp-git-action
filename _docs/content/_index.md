---
title: drone-git-action
---

[![Build Status](https://img.shields.io/drone/build/thegeeklab/drone-git-action?logo=drone&server=https%3A%2F%2Fdrone.thegeeklab.de)](https://drone.thegeeklab.de/thegeeklab/drone-git-action)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/thegeeklab/drone-git-action)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/drone-git-action)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/drone-git-action)](https://github.com/thegeeklab/drone-git-action/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/thegeeklab/drone-git-action)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/drone-git-action)](https://github.com/thegeeklab/drone-git-action/blob/main/LICENSE)

Drone plugin to execute git actions.

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< toc >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

## Usage

```YAML
kind: pipeline
name: default

steps:
  - name: commit artifact
    image: thegeeklab/drone-git-action
    settings:
      netrc_password: ghp_3LbMg9Kncpdkhjp3bh3dMnKNXLjVMTsXk4sM
      author_name: octobot
      author_email: octobot@example.com
      message: "[skip ci] update changelog"
```

### Parameters

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< propertylist name=drone-git-action.data sort=name >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

## Build

Build the binary with the following command:

```Shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

make build
```

Build the Docker image with the following command:

```Shell
docker build --file docker/Dockerfile.amd64 --tag thegeeklab/drone-git-action .
```

## Test

```Shell
docker run --rm \
  -e DRONE_BUILD_EVENT=pull_request \
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=foo \
  -e DRONE_PULL_REQUEST=1
  -e PLUGIN_API_KEY=abc123 \
  -e PLUGIN_MESSAGE="Demo comment" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  thegeeklab/drone-git-action
```
