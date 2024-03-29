# go-bin-template
[![Actions Status](https://github.com/Eun/go-bin-template/workflows/push/badge.svg)](https://github.com/Eun/go-bin-template/actions)
[![Coverage Status](https://coveralls.io/repos/github/Eun/go-bin-template/badge.svg?branch=master)](https://coveralls.io/github/Eun/go-bin-template?branch=master)
[![PkgGoDev](https://img.shields.io/badge/pkg.go.dev-reference-blue)](https://pkg.go.dev/github.com/Eun/go-bin-template)
[![go-report](https://goreportcard.com/badge/github.com/Eun/go-bin-template)](https://goreportcard.com/report/github.com/Eun/go-bin-template)
---
A golang template for executables

## Setup
1. Adjust `.env`
2. Adjust `.goreleaser.yml`
3. Adjust `.golangci.yml`


## Docker
Docker images will be pushed to GitHub's registry by default.
> If you don't need docker remove the `dockers` section in `.goreleaser.yml`
> and the `Login to Docker Registry` step in `release_published.yml`.

## Release Process
Every closed pull request will update (or create) a new drafted [release](https://github.com/Eun/go-bin-template/releases).
If you wish to release, just edit the drafted release and publish it.
GoReleaser will build the binary and add the binaries it to the release.

## Notes
Testing and building will always be done with the go version specified in the `go.mod` file.

## Build History
[![Build history](https://buildstats.info/github/chart/Eun/go-bin-template?branch=master)](https://github.com/Eun/go-bin-template/actions)