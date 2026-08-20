# Contributing to Forge

Forge is currently defining and validating a deliberately narrow MVP. Changes
that improve that vertical slice are welcome; speculative cluster, web, multi-OS,
plugin, or custom-runtime features are not yet in scope.

## Prerequisites

- Go 1.26 or newer.
- Git.
- Linux for future runtime integration tests. Unit tests must remain portable.

The scaffold has no third-party Go dependencies.

## Local checks

Run these before submitting a change:

```console
test -z "$(gofmt -l .)"
go vet ./...
go test -race ./...
go build ./cmd/forge
go run ./tools/repolint
```

## Change rules

- Keep the dependency set minimal and justify every new module in the PR.
- Preserve deterministic, idempotent, and inspectable reconciliation behavior.
- Add tests for both the successful path and meaningful failure boundaries.
- Do not log credentials, tokens, manifest secrets, or unredacted environment.
- Update `docs/MVP.md` when acceptance behavior changes.
- Add an ADR for public contracts, dependencies, persistent state, supported
  runtimes, security boundaries, or changes to the MVP scope.
- Update `SECURITY.md` when privileges or trust boundaries change.
- Do not mix unrelated refactors with behavioral changes.

## Issues and pull requests

Open an issue before implementing a substantial change so its fit with the MVP
can be established. A pull request should explain the problem, chosen approach,
tests, failure behavior, and documentation impact. Keep commits focused and do
not add generated-by or co-author trailers.

The `main` branch accepts changes only through a green pull request. Pull request
titles become squash commit subjects and should use `type: summary`, where type
is one of `feat`, `fix`, `docs`, `test`, `refactor`, `build`, `ci`, or `chore`.
Commits and release tags merged into `main` must carry a verified signature.

Participation is governed by [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md).

By contributing, you agree that your contribution is licensed under Apache-2.0.
