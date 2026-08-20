# Forge

**GitOps reconciliation for a single Linux container host.**

Forge is an early-stage open-source project for operators who want containers to
converge from a Git repository without adopting a cluster orchestrator. A local
agent will read a pinned Git revision, validate the desired state, calculate a
plan, and reconcile the host through an interchangeable OCI runtime adapter.

> [!IMPORTANT]
> Forge is currently a design scaffold. It does not deploy containers yet. The
> repository deliberately defines the product contract before implementing it.

## The problem

Running a few containers on one Linux server is easy. Keeping their complete
state reproducible, reviewed, observable, and recoverable over time often grows
into a collection of Compose files, systemd units, CI jobs, and SSH scripts.
Forge aims to provide one small reconciliation loop for that gap.

The initial user is an individual operator or small platform team managing one
Linux host. Git remains the desired-state authority; Forge owns local
validation, planning, application, status, logs, and explicit rollback.

## Intended workflow

The following is the target experience, not an implemented interface:

```console
$ forge validate
configuration is valid at 8f31c42

$ forge plan
~ web  ghcr.io/example/web@sha256:old -> ghcr.io/example/web@sha256:new

$ forge apply
applied 8f31c42

$ forge status
revision 8f31c42  state converged

$ forge rollback
restored the last successfully applied revision
```

An illustrative manifest may look like this; the schema is intentionally not
stable yet:

```yaml
version: forge/v1alpha1
services:
  web:
    image: ghcr.io/example/web@sha256:0123456789abcdef
    ports:
      - "8080:8080"
    restart: always
```

## MVP boundary

The first usable release is deliberately narrow:

- Linux and one host;
- one Git repository and an explicit revision policy;
- declarative validation, planning, apply, status, logs, and rollback;
- an idempotent local reconciliation agent;
- OCI-compatible images and an interchangeable low-level runtime boundary;
- durable local records of desired, observed, and last successful state.

Forge is not initially a cluster scheduler, a mini-Kubernetes, a web platform,
a fleet manager, a multi-OS abstraction, or a new low-level OCI runtime. It also
does not promise atomic rollback where the host runtime cannot provide it.

## Repository map

- [`docs/VISION.md`](docs/VISION.md) — users, problem, principles, and success.
- [`docs/MVP.md`](docs/MVP.md) — exact lifecycle, failure behavior, and acceptance.
- [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md) — component and state boundaries.
- [`docs/adr/`](docs/adr/) — decisions that must remain explainable.
- [`SECURITY.md`](SECURITY.md) — trust boundaries and vulnerability reporting.
- [`CONTRIBUTING.md`](CONTRIBUTING.md) — scope and engineering workflow.
- [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md) — expected community behavior.

## Development

Forge requires Go 1.26 or newer.

```console
go test ./...
go vet ./...
go build ./cmd/forge
go run ./tools/repolint
go run ./cmd/forge version
```

No third-party Go module is required by the scaffold.

## License

Forge is licensed under the [Apache License 2.0](LICENSE).
