# 0005 — Make apply and rollback honest about partial failure

- Status: accepted
- Date: 2026-08-20

## Context

Container lifecycle operations and application side effects are not generally
transactional. Claiming atomic deployment or automatic rollback would create a
dangerous false guarantee.

## Decision

Forge validates the full configuration and prepares required artifacts before
mutation. It then applies a deterministic sequential plan under one mutation
lock. On failure it stops, observes actual state, persists the partial result,
and requires a newly calculated plan.

Rollback reapplies a previous successful immutable revision through the same
validation, planning, preparation, application, and verification lifecycle. It
is explicit and best effort; partial rollback is reported as degraded.

## Alternatives considered

- **Automatic compensation:** may destroy a still-serving container or repeat
  irreversible external effects.
- **Pretend the complete stack is transactional:** unsupported by common host
  runtime primitives.
- **No rollback command:** forces operators to reconstruct recovery metadata
  during an incident.

## Consequences

Status and operation journals are first-class MVP data. Recovery may require
operator action, but Forge will not hide uncertainty or report false health.
