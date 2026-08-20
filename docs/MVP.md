# Minimum viable product

## Outcome

The MVP proves that Forge can safely reconcile one containerized application
stack on one Linux host from one Git repository. It is a product slice, not a
collection of disconnected runtime experiments.

## Included capabilities

### Configuration and Git

- Read one repository from a configured local path or authenticated remote.
- Resolve the configured branch, tag, or commit to an immutable commit ID.
- Load one versioned Forge manifest from that revision.
- Validate the complete manifest before planning or mutation.
- Reject unknown or unsafe fields rather than silently ignoring them.
- Never store secret values directly in the manifest or persisted status.

### Planning and reconciliation

- Inspect desired state, observed runtime state, and last applied state.
- Produce a deterministic plan with create, update, restart, and remove actions.
- Offer a non-mutating validation and plan path through the CLI.
- Resolve images before mutation and prefer immutable digest references.
- Apply a plan sequentially through one reference runtime adapter.
- Re-run idempotently: no drift means no runtime operation.
- Detect supported forms of manual drift and reconcile them.

### State and operation

- Persist the desired commit, last attempted commit, last successful commit,
  observed resources, operation result, and timestamps.
- Expose concise machine-readable and human-readable status.
- Record structured operational logs without credentials or secret values.
- Resume safely after an agent or host restart.
- Serialize mutations so the CLI and agent cannot apply competing plans.

### Rollback

- Keep enough local metadata to locate the last successfully applied revision.
- Run rollback through validation and planning like any other apply.
- Stop and report if required source or OCI artifacts are unavailable.
- Describe rollback as best effort: container operations are not transactional.
- Mark automatic reconciliation paused if an emergency local rollback diverges
  from the configured Git revision, until the operator resolves that divergence.

## Reconciliation lifecycle

1. **Fetch:** obtain repository metadata without altering workload state.
2. **Resolve:** pin the configured reference to an immutable commit.
3. **Load:** read the manifest from that exact commit.
4. **Validate:** reject invalid, unsupported, or unsafe desired state.
5. **Observe:** query the runtime and durable local state.
6. **Plan:** calculate and display a deterministic ordered action list.
7. **Prepare:** resolve and fetch all required OCI artifacts.
8. **Apply:** execute one action at a time under a host-wide mutation lock.
9. **Verify:** inspect resulting resources and health information.
10. **Record:** persist the complete result and emit structured logs.

Only a fully verified revision becomes the last successful revision.

## Failure contract

- Git unavailable before resolution: keep the current workload and report stale
  source state; never guess a revision.
- Invalid configuration: perform no workload mutation.
- Artifact unavailable during preparation: perform no workload mutation.
- Runtime unavailable: preserve desired state, report degraded observed state,
  and retry with bounded backoff.
- Apply interrupted or partially failed: stop remaining actions, inspect actual
  runtime state, persist the partial result, and require a fresh plan.
- State store unavailable: refuse new mutations because recovery cannot be
  explained reliably.
- Agent restart: recover from durable state and observed runtime state; never
  assume the interrupted plan completed.

The MVP does not automatically roll back a partially failed apply. Automatic
compensation without transactional runtime guarantees could destroy a workload
that is still serving traffic.

## CLI surface

The target command surface is deliberately small:

```text
forge validate
forge plan
forge apply
forge status [--json]
forge logs
forge rollback [<revision>]
forge version
```

The long-running agent will use a separate `forged` executable. Exact flags and
configuration locations remain provisional until the vertical prototype.

## Excluded from MVP

- more than one managed host per Forge instance;
- cluster scheduling and cross-host service discovery;
- web UI or remote control plane;
- automatic Git writes or pull-request creation;
- secret storage or secret synchronization;
- build pipelines and OCI registry hosting;
- custom networking, storage, or low-level runtime implementations;
- Windows and macOS agents;
- self-update, plugins, policy languages, and extension marketplaces.

## Acceptance tests

The MVP is complete only when automated integration tests demonstrate:

- clean install and first apply on the reference Linux distribution;
- a second identical reconciliation with zero runtime mutations;
- image update, configuration update, restart, and removal;
- manual drift detection and repair;
- invalid manifest and unavailable image causing zero mutations;
- runtime outage and recovery;
- interruption mid-apply followed by truthful recovery;
- agent restart preserving the last successful revision;
- explicit rollback to the previous successful revision;
- secret-like values absent from status and logs.
