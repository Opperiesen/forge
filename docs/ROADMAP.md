# Forge roadmap

This roadmap orders Forge by demonstrated product capabilities. It is not a
calendar or a second issue tracker. A milestone closes only when its exit gate
can be reproduced from a fresh checkout.

## Sources of truth

| Question | Authority |
| --- | --- |
| Why Forge exists and for whom | [`VISION.md`](VISION.md) |
| What the first product includes and excludes | [`MVP.md`](MVP.md) |
| How the system is divided | [`ARCHITECTURE.md`](ARCHITECTURE.md) |
| Why a durable technical choice was made | [`adr/`](adr/) |
| Which outcomes come next and what proves them | This roadmap |
| What work is executable now | GitHub milestone issues |
| What change is under review | One linked pull request |

Status is not duplicated between documents. GitHub owns issue and milestone
progress; this file owns ordering, dependencies, and exit gates. Completed work
is described here only when it changes a gate or the roadmap itself.

## Planning model

Forge uses a rolling horizon:

- the current milestone is decomposed into small, ordered issues;
- the next milestone has a defined outcome and gate, then is decomposed near the
  end of the current milestone;
- later milestones retain outcomes, dependencies, and gates without speculative
  implementation tasks;
- dates are added only after M1 and M2 provide measured delivery evidence.

Only one milestone is active and one implementation issue should normally be in
progress. A blocked issue names its dependency. Research work must end in a
testable artifact, an ADR, or an explicit rejection.

## Current position

**Active milestone:** [M1 — Manifest v1alpha1](https://github.com/Opperiesen/forge/milestone/2)

**Next executable issue:** [#3 — Define the canonical Stack model](https://github.com/Opperiesen/forge/issues/3)

M1 is deliberately isolated from Git, host discovery, networking, and container
runtimes. It first establishes the trusted canonical specification consumed by
all later stages.

## Delivery sequence

`M0 → M1 → M2 → M3 → M4 → M5 → M6 → M7 → M8 → M9 → M10`

The sequence is dependency-driven. A later milestone may be researched early,
but its production implementation does not begin before its prerequisites close
unless an ADR records why the ordering changed.

## Milestones and exit gates

### [M0 — Foundation](https://github.com/Opperiesen/forge/milestone/1) — complete

The repository establishes the product boundary, engineering rules, architecture,
security model, contribution workflow, and protected delivery path.

Exit gate:

- vision, MVP, architecture, and initial ADRs agree on the single-host product;
- local formatting, vet, race tests, build, and repository hygiene pass;
- `main` requires review-quality CI and CodeQL checks through pull requests;
- security reporting, dependency updates, secret scanning, and merge policy are
  configured;
- the roadmap and operating protocol make future work resumable.

### [M1 — Manifest v1alpha1](https://github.com/Opperiesen/forge/milestone/2) — active

Untrusted manifest bytes become one versioned, validated, deterministic,
runtime-independent canonical `Stack`.

Exit gate:

- the canonical model and every invariant are documented and unit tested;
- the manifest syntax and parser are justified by an ADR and strict-decoding
  evidence;
- parsing, defaults, normalization, and semantic validation fail closed;
- equivalent documents produce equal canonical specifications;
- `forge validate` has stable diagnostics and exit behavior;
- normative examples are continuously tested;
- acceptance tests require no Git repository, network, runtime, privilege, or
  hidden machine state.

Executable chain: [#3](https://github.com/Opperiesen/forge/issues/3) →
[#4](https://github.com/Opperiesen/forge/issues/4) →
[#5](https://github.com/Opperiesen/forge/issues/5) →
[#6](https://github.com/Opperiesen/forge/issues/6) →
[#7](https://github.com/Opperiesen/forge/issues/7) →
[#8](https://github.com/Opperiesen/forge/issues/8) →
[#9](https://github.com/Opperiesen/forge/issues/9) →
[#10](https://github.com/Opperiesen/forge/issues/10).

### [M2 — Git Snapshot](https://github.com/Opperiesen/forge/milestone/3) — planned

Forge resolves one configured source to an immutable commit and loads the
manifest from that snapshot without treating a moving branch as applied state.

Exit gate:

- source URL, revision policy, commit identity, and manifest path have explicit
  contracts;
- fetch, authentication, timeout, unavailable source, missing file, and invalid
  revision failures are bounded and actionable;
- the exact commit and content identity flow into later status records;
- tests use isolated repositories and perform no runtime mutation.

### [M3 — Runtime Contract](https://github.com/Opperiesen/forge/milestone/4) — planned

Forge obtains observed state and requests operations through a narrow OCI
runtime adapter, with one Linux reference adapter selected from evidence.

Exit gate:

- desired-state logic contains no runtime-specific types;
- observation, lifecycle operations, capabilities, errors, and cancellation are
  represented by a tested adapter contract;
- a Linux prototype validates the riskiest API and privilege assumptions;
- an ADR selects the reference adapter and records rejected alternatives.

### [M4 — Deterministic Plan](https://github.com/Opperiesen/forge/milestone/5) — planned

Forge compares canonical desired state with normalized observed state and emits
an ordered, explainable plan without mutating the host.

Exit gate:

- identical desired and observed inputs always produce the same plan;
- create, update, remove, replace, and no-op decisions are explicit;
- ordering and dependency behavior are stable and tested;
- capability gaps and unsafe transitions fail before apply;
- planning has no runtime mutation or durable-state side effect.

### [M5 — Manual Apply](https://github.com/Opperiesen/forge/milestone/6) — planned

An operator can manually validate, plan, and apply one real container stack on
one Linux host through the reference adapter.

Exit gate:

- the first apply reaches the requested state and records its attempt and result;
- a second apply with unchanged inputs is a no-op;
- partial failure is reported without claiming atomicity or false convergence;
- the last successful revision remains distinguishable from the latest attempt;
- a fresh integration environment reproduces the full manual lifecycle.

Release checkpoint: `v0.1.0-alpha.1` may be cut when this gate closes.

### [M6 — Durable Convergence](https://github.com/Opperiesen/forge/milestone/7) — planned

Reconciliation remains correct across process restarts, concurrent invocation,
and interrupted operations.

Exit gate:

- desired, observed, attempted, and last-successful records have durable schemas;
- a host-level lock prevents overlapping mutation;
- restart and interruption recovery are deterministic and tested;
- persisted state can be inspected and corruption fails safely;
- repeated convergence is idempotent.

### [M7 — Autonomous Agent](https://github.com/Opperiesen/forge/milestone/8) — planned

The local `forged` agent watches the configured Git source and continuously
converges one host without requiring a central control plane.

Exit gate:

- polling, cancellation, backoff, and retry policies are bounded;
- a new commit triggers reconciliation and records the resolved identity;
- runtime drift is detected and corrected according to policy;
- invalid desired state never replaces the last successful state;
- status distinguishes fetching, planning, applying, converged, degraded, and
  blocked states.

Release checkpoint: `v0.1.0-alpha.2` may be cut when this gate closes.

### [M8 — Failure and Rollback](https://github.com/Opperiesen/forge/milestone/9) — planned

Forge makes interrupted and partially failed operations inspectable and can
explicitly restore the last successfully applied revision where capabilities
permit.

Exit gate:

- preparation and operation journaling expose the exact completed boundary;
- restart resumes or safely replans interrupted work;
- rollback creates a new explicit attempt rather than rewriting history;
- unsupported or unsafe rollback is rejected before mutation;
- failure-injection tests cover interruption at every mutating boundary.

### [M9 — Operability and Security](https://github.com/Opperiesen/forge/milestone/10) — planned

Forge is diagnosable and hardened for unattended operation on a single Linux
host.

Exit gate:

- human and structured status, logs, and stable error categories are available;
- credentials and environment values are redacted and never persisted by
  default;
- least-privilege execution and filesystem permissions are documented and tested;
- image digest and Git revision provenance are visible;
- a hardened system service definition and operational runbook are verified;
- threat boundaries and recovery procedures match the implementation.

Release checkpoint: `v0.1.0-rc.1` may be cut when this gate closes.

### [M10 — v0.1.0](https://github.com/Opperiesen/forge/milestone/11) — planned

The first supported release installs cleanly and demonstrates the complete MVP
lifecycle on a fresh Linux host.

Exit gate:

- a clean host can install, configure, validate, plan, apply, observe, update,
  recover, and roll back the documented example;
- end-to-end and failure-injection suites pass against the reference adapter;
- user, operator, security, troubleshooting, and upgrade documentation agree;
- versioned binaries, checksums, SBOM, provenance, and signatures are published;
- known limitations and support boundaries are explicit;
- the `v0.1.0` tag points to the reviewed release commit.

## Milestone lifecycle

A milestone moves through three states:

1. **Planned:** outcome, dependency, and exit gate exist here; work is not yet
   fully decomposed.
2. **Active:** all known work is represented by ordered GitHub issues with tests,
   failure cases, dependencies, and non-goals.
3. **Complete:** all gate evidence passes from a fresh environment, no hidden
   blocker remains, residual work is moved deliberately, and the GitHub milestone
   is closed.

At closure, update this roadmap only for a changed gate, sequence, or status;
close or move every remaining issue; record any durable decision in an ADR; then
decompose the next milestone. Passing CI alone does not close a milestone.
