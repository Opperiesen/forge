# Forge repository instructions

## Product invariant

Forge is a GitOps reconciliation engine for one Linux container host. The MVP is
not a cluster orchestrator, fleet manager, web platform, multi-OS abstraction, or
low-level OCI runtime.

Read `docs/VISION.md`, `docs/MVP.md`, `docs/ARCHITECTURE.md`, and
`docs/ROADMAP.md` before changing behavior. Accepted decisions live in
`docs/adr/` and must be superseded by a new ADR rather than silently rewritten.

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

## Operating protocol

### Resume a work session

1. Read this file and the current position in `docs/ROADMAP.md`.
2. Confirm that no other writer is working in this repository; inspect any
   suspicious process completely rather than assuming ownership from its name.
3. Inspect the worktree, fetch the remote, and start from the current `main`.
4. Inspect the active GitHub milestone, its open issues, blockers, and open pull
   requests. GitHub is authoritative for live work status.
5. Read only the product documents and ADRs relevant to the selected issue.
6. Select the first unblocked issue in the active milestone unless an incident or
   maintainer decision explicitly changes priority.

### Authority and change control

- `docs/VISION.md` owns users, problem, and principles.
- `docs/MVP.md` owns the release boundary and lifecycle promises.
- `docs/ARCHITECTURE.md` owns component and trust boundaries.
- `docs/adr/` owns durable technical decisions and their rationale.
- `docs/ROADMAP.md` owns milestone order, dependencies, and exit gates.
- GitHub milestones and issues own live execution status and acceptance work.
- Pull requests own proposed changes; commits do not redefine product scope.

Do not duplicate live percentages, issue checklists, or target dates in Markdown.
If sources disagree, stop implementation and reconcile the highest relevant
authority through review.

### Execute an issue

- Keep one implementation issue and one code pull request in progress whenever
  practical. Parallel agents may analyze independent questions, but only one
  writer may modify the worktree.
- Branch from current `main`; implement one coherent behavior; include tests,
  failure behavior, and documentation in the same pull request.
- Treat an issue dependency as a hard blocker. A spike must produce testable
  evidence, an ADR, or an explicit rejected option.
- Run all validation commands in this file before commit and before merge.
- Open a draft pull request linked to the issue, review the diff and check
  results, then mark it ready only when the acceptance criteria are met.
- Require the protected checks, including CodeQL, to pass. Squash merge so the
  pull request title becomes the permanent `main` history entry.
- Confirm the issue closed, the remote branch was removed, local `main` is
  aligned, and the worktree is clean before selecting more work.

### Close a milestone

- Reproduce every exit gate from a fresh environment; CI success by itself is
  insufficient.
- Account for every issue: close it, move it deliberately, or document why it no
  longer applies. Do not hide incomplete work in prose.
- Update `docs/ROADMAP.md` only when status, sequence, outcome, or gate changes.
- Add or supersede ADRs for durable decisions discovered during delivery.
- Close the GitHub milestone, then fully decompose the next milestone before
  beginning its implementation.

Long-term notes outside the repository retain only durable project status and
decisions. The repository and GitHub remain authoritative for engineering detail.
