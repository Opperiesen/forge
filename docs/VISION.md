# Vision

## Product statement

Forge is a small GitOps engine for a single Linux container host. It continuously
compares a desired revision in Git with observable local state and performs the
minimum safe actions needed to converge them.

## Intended users

The primary user is an infrastructure practitioner or small platform team that:

- operates one Linux server at a time;
- wants reviewable container changes and deterministic recovery;
- finds ad-hoc Compose, systemd, CI, and SSH glue increasingly fragile;
- does not need the scheduling and distributed control plane of Kubernetes.

The MVP optimizes for an experienced operator. A beginner-friendly web interface
may be considered later, but cannot shape the initial engine.

## Job to be done

Given a reviewed Git revision, make the container workload on one host match it,
show precisely what happened, and provide a controlled path back to the last
known successful revision.

## Value proposition

Forge should make these properties routine:

- **Reproducibility:** the desired configuration and exact images are versioned.
- **Traceability:** every plan and result refers to an immutable Git commit.
- **Convergence:** drift is detected and corrected by an idempotent agent.
- **Understandability:** a plan is visible before destructive replacement.
- **Recovery:** successful revisions are locally recorded and can be reapplied.
- **Portability:** OCI boundaries avoid coupling the product to one runtime.

## Design principles

1. Git describes desired state; runtime inspection describes observed state.
2. The engine must remain useful on one host before any fleet feature exists.
3. A failed operation is reported honestly; partial work is never called atomic.
4. Inputs are validated and artifacts resolved before host mutation begins.
5. Runtime-specific behavior stays behind a narrow adapter contract.
6. Secure defaults beat convenience when configuration can control the host.
7. Every dependency and subsystem must justify its place in the MVP lifecycle.

## Explicit non-goals

The MVP will not provide:

- multi-host scheduling, leader election, quorum, or cluster networking;
- Kubernetes API compatibility, controllers, or custom resources;
- a web UI, hosted control plane, marketplace, or account system;
- Windows, macOS, Android, or iOS workload management;
- endpoint inventory, patch management, MDM, or remote desktop;
- a new implementation of namespaces, cgroups, or the OCI runtime specification;
- a general secret-management service;
- zero-downtime deployment guarantees for every workload.

## MVP success

The MVP succeeds when a fresh supported Linux host can:

1. follow an explicit Git revision policy;
2. validate and plan a small container stack without mutating the host;
3. converge that stack twice with the second run producing no changes;
4. detect manual drift and restore the desired state;
5. survive an agent restart without losing the last applied revision;
6. report a partial failure without corrupting its state record;
7. reapply the last successful revision through the normal planning path;
8. expose enough status and logs to explain every action.

Performance, cluster scale, and UI polish are not MVP success criteria.

## Open decisions

The framework intentionally leaves three implementation choices open until a
vertical prototype provides evidence:

- the first reference OCI runtime;
- the exact manifest schema and compatibility policy;
- the durable local state backend.

Each choice requires an ADR before it becomes a public contract.
