# go-bin-template [![Actions Status](https://github.com/Eun/go-bin-template/workflows/CI/badge.svg)](https://github.com/Eun/go-bin-template/actions) [![Codecov](https://img.shields.io/codecov/c/github/Eun/go-bin-template.svg)](https://codecov.io/gh/Eun/go-bin-template) [![GoDoc](https://godoc.org/github.com/Eun/go-bin-template?status.svg)](https://godoc.org/github.com/Eun/go-bin-template) [![go-report](https://goreportcard.com/badge/github.com/Eun/go-bin-template)](https://goreportcard.com/report/github.com/Eun/go-bin-template)
A golang template for executables

## Setup
1. Adjust `.goreleaser.yml`
2. Adjust `.golangci.yml`
3. Setup Dockerhub


## Dockerhub
1. Adjust the `Dockerfile` in the `.dockerhub` folder
2. Setup following `Build Rules` to enable automatic building:
    ```
    Source Type | Source         | Docker Tag | Dockerfile location | Build Context
    Tag         | /^v([0-9.]+)$/ | latest     | Dockerfile          | /.dockerhub
    Tag         | /^v([0-9.]+)$/ | v{\1}      | Dockerfile          | /.dockerhub
    ```
