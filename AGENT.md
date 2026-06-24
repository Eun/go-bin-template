# AGENT.md — Instructions for adapting this template

This document is for AI agents (and humans) who are using this repository as a
template for a new Go project. It describes what needs to be changed and how the
project is structured.

## Checklist: adapting for a new project

### 1. `go.mod`
- [ ] Change the module path from `github.com/Eun/go-bin-template` to your module path
- [ ] Run `go mod tidy`

### 2. `mise.toml`
- [ ] Update Go version if needed (must match `go.mod`)
- [ ] Update `tasks.build` — change binary name from `go-bin-template` to your project name
- [ ] Update `tasks.clean` — change binary name
- [ ] Tool versions (`golangci-lint`, `goreleaser`) can be updated as needed

### 3. `.goreleaser.yml`
- [ ] Change `project_name` to your project name
- [ ] Update `builds[].id` and `builds[].binary`
- [ ] Update `builds[].ldflags` — change the `-X` paths to match your module path
- [ ] Update `archives[].name_template` — change project name prefix
- [ ] Update `nfpms[]` — change `id`, `package_name`, `homepage`, `maintainer`, `description`
- [ ] Update `dockers[].build_flag_templates` — change `--build-arg=BINARY=` value
- [ ] Remove `dockers` and `docker_manifests` sections entirely if you don't need Docker

### 4. `.golangci.yml`
- [ ] Update `formatters.settings.goimports.local-prefixes` to your module path
- [ ] Adjust linter settings/rules as needed for your project

### 5. `.github/workflows/release_published.yml`
- [ ] Update `IMAGE_NAME` env var to your Docker image name (e.g. `yourorg/yourproject`)
- [ ] Update `REGISTRY` if not using `ghcr.io`
- [ ] Update `docker/login-action` credentials if using a different registry
- [ ] Remove Docker-related steps if not publishing images

### 6. `.gitignore`
- [ ] Change `/go-bin-template` to your binary name

### 7. Source code
- [ ] Replace `cmd/hello/` with your actual commands
- [ ] Update `cmd/root/root.go` — app name, description, commands
- [ ] Update import paths everywhere to match your module path

### 8. `README.md`
- [ ] Rewrite for your project

### 9. `.github/FUNDING.yml`
- [ ] Update or remove sponsor links

## Architecture

### Tool management
All tools (Go, golangci-lint, goreleaser) are managed by [mise](https://mise.jdx.dev/).
Versions are pinned in `mise.toml`. Run `mise install` to install them.

### Tasks
All build/lint/test/release operations are defined as mise tasks in `mise.toml`.
CI workflows only checkout, install mise, and run these tasks — no build logic
lives in GitHub Actions YAML.

| Task | What it does |
|---|---|
| `build` | `go build -o <binary> .` |
| `lint` | `golangci-lint run ./...` |
| `test` | `go test` with `-cover -coverprofile=coverage.out -covermode=atomic` |
| `release` | `goreleaser release --snapshot --clean` |
| `clean` | Remove binary and `dist/` directory |

### CI workflows

| File | Trigger | Purpose |
|---|---|---|
| `ci.yml` | `push` (all branches) + `pull_request` | Lint, vulnerability scan, test + coverage upload, coverage PR comment, build |
| `release_draft.yml` | `push` (default branch only, via `if` condition) | Draft release notes |
| `release_published.yml` | GitHub release published | GoReleaser: build binaries, packages, Docker images |
| `pull_request_target_opened.yml` | PR opened | Auto-label based on title (feat/fix/chore) |

### Code coverage
- The `test` mise task produces `coverage.out`
- In CI, the `test` job uploads this as a `code-coverage` artifact
- The `code_coverage` job (PR-only) uses `fgrosse/go-coverage-report` to download
  the artifact and post a coverage diff as a PR comment
- Coverage from pushes to the default branch is used as the baseline for comparison

### Docker releases
- `.goreleaser.yml` uses `{{.Env.REGISTRY}}/{{.Env.IMAGE_NAME}}` for image names
- `release_published.yml` passes `REGISTRY`, `IMAGE_NAME`, and `IS_LATEST_RELEASE` as env vars
- `IS_LATEST_RELEASE` controls whether the `:latest` tag is pushed (skipped for prereleases)
- Multi-arch manifests combine `amd64` and `arm64` images

### Linting
- `.golangci.yml` uses the v2 format
- Local import prefix is configured under `formatters.settings.goimports.local-prefixes`
- The `err` shadow warning from `govet` is excluded via an exclusion rule

### Dependencies
- Dependabot is configured for both `gomod` and `github-actions` (daily at 04:00)
- Dependencies are vendored in `vendor/`

### Release notes
- release-drafter auto-drafts release notes from merged PRs
- PRs are auto-labeled based on title prefix: `feat` → feature, `fix` → fix, `chore` → chore
- Version bumps are determined by labels: `major`, `minor`, `patch` (default: patch)
