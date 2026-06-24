# go-bin-template
[![CI](https://github.com/Eun/go-bin-template/actions/workflows/ci.yml/badge.svg)](https://github.com/Eun/go-bin-template/actions/workflows/ci.yml)
[![PkgGoDev](https://img.shields.io/badge/pkg.go.dev-reference-blue)](https://pkg.go.dev/github.com/Eun/go-bin-template)
[![go-report](https://goreportcard.com/badge/github.com/Eun/go-bin-template)](https://goreportcard.com/report/github.com/Eun/go-bin-template)
---
A golang template for executables.

## Quick Start
```bash
mise install        # install Go, golangci-lint, goreleaser
mise run build      # build the binary
mise run test       # run tests with coverage
mise run lint       # run linter
```

## Prerequisites
- [mise](https://mise.jdx.dev/) — manages tool versions and tasks

All tool versions (Go, golangci-lint, goreleaser) are pinned in `mise.toml`.

## Setup
When using this template for a new project, adjust the following:
1. `mise.toml` — tool versions if needed
2. `.goreleaser.yml` — project name, ldflags, docker image references
3. `.golangci.yml` — linter settings and local import prefix
4. `.github/workflows/release_published.yml` — `REGISTRY`, `IMAGE_NAME` env vars
5. `go.mod` — module name

See [AGENT.md](AGENT.md) for a detailed checklist.

## Mise Tasks
| Task | Description |
|---|---|
| `mise run build` | Build the binary |
| `mise run lint` | Run golangci-lint |
| `mise run test` | Run tests with coverage (`coverage.out`) |
| `mise run release` | Run GoReleaser snapshot |
| `mise run clean` | Remove build artifacts |

## CI / CD
All CI logic is driven by mise — GitHub Actions workflows only checkout, install mise, and run tasks.

| Workflow | Trigger | What it does |
|---|---|---|
| `ci.yml` | Push + PR | Lint, vulnerability scan (nancy), test with coverage, build, coverage PR comment |
| `release_draft.yml` | Push to default branch | Drafts release notes via release-drafter |
| `release_published.yml` | Release published | GoReleaser builds binaries, packages, and Docker images |
| `pull_request_target_opened.yml` | PR opened | Auto-labels PRs based on title |

### Code Coverage
- Tests produce `coverage.out` via `-coverprofile`
- On PRs, the `code_coverage` job uses [fgrosse/go-coverage-report](https://github.com/fgrosse/go-coverage-report) to post a coverage diff comment

## Docker
Docker images are pushed to GitHub Container Registry (`ghcr.io`) on release.
Multi-arch manifests are created for `linux/amd64` and `linux/arm64`.

The `:latest` tag is only pushed for non-prerelease releases.

> If you don't need Docker, remove the `dockers` and `docker_manifests` sections in `.goreleaser.yml`
> and the Docker-related steps in `release_published.yml`.

## Release Process
1. Every merged PR updates a drafted [release](https://github.com/Eun/go-bin-template/releases) via release-drafter
2. Edit the draft and publish it
3. GoReleaser builds binaries, packages, and Docker images automatically

## Project Structure
```
go-bin-template/
├── main.go                              # Entry point
├── cmd/
│   ├── root/root.go                     # CLI app setup (flags, commands, logging)
│   └── hello/
│       ├── hello.go                     # Example "hello" command
│       └── hello_test.go                # Tests for hello command
├── mise.toml                            # Tool versions + tasks (build, lint, test, release, clean)
├── .goreleaser.yml                      # GoReleaser config (multi-arch Docker + deb/rpm)
├── .golangci.yml                        # Linter configuration (v2 format)
├── .github/
│   ├── Dockerfile                       # Distroless container image
│   ├── FUNDING.yml                      # GitHub sponsors
│   ├── PULL_REQUEST_TEMPLATE.md
│   ├── dependabot.yml                   # Auto-update Go modules + GitHub Actions
│   ├── labeler.yml                      # PR label rules
│   ├── release-drafter.yml              # Release notes template
│   └── workflows/
│       ├── ci.yml                       # Lint, vuln scan, test, coverage, build
│       ├── release_draft.yml            # Draft release notes on push to default branch
│       ├── release_published.yml        # GoReleaser release + Docker push
│       └── pull_request_target_opened.yml  # Auto-label PRs
├── go.mod
├── go.sum
├── vendor/                              # Vendored dependencies
├── LICENSE
├── README.md
└── AGENT.md                             # Instructions for AI agents adapting this template
```
