# Forge repository instructions

## Product invariant

Forge is a GitOps reconciliation engine for one Linux container host. The MVP is
not a cluster orchestrator, fleet manager, web platform, multi-OS abstraction, or
low-level OCI runtime.

Read `docs/VISION.md`, `docs/MVP.md`, and `docs/ARCHITECTURE.md` before changing
behavior. Accepted decisions live in `docs/adr/` and must be superseded by a new
ADR rather than silently rewritten.

## Engineering rules

- Go is the primary language; keep the core compatible with Go 1.26.
- Prefer the standard library. Every dependency needs an explicit MVP use case.
- Keep source, state, and runtime behind narrow boundaries.
- Treat Git and manifests as untrusted input.
- Preserve idempotence and deterministic plans.
- Never claim atomic apply or rollback across runtime operations.
- Never place secrets in fixtures, manifests, output, logs, or commits.
- Do not implement deferred features without first changing the MVP decision.
- Follow existing package conventions and keep packages cohesive.
- Add no commit trailers or tool signatures.

## Validation

Run before each commit:

```console
test -z "$(gofmt -l .)"
go vet ./...
go test -race ./...
go build ./cmd/forge
go run ./tools/repolint
```

Runtime integration must eventually run on Linux against the reference adapter;
unit tests must not require a running container daemon.
