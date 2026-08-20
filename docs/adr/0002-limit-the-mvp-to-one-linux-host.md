# 0002 — Limit the MVP to one Linux host

- Status: accepted
- Date: 2026-08-20

## Context

The useful gap is between ad-hoc single-host scripts and cluster orchestrators.
Adding host coordination would introduce scheduling, identity, networking,
consensus, and failure domains before local convergence is proven.

## Decision

The MVP manages one Linux host per Forge agent. It has no distributed control
plane and makes no cross-host availability guarantee.

## Alternatives considered

- **Mini-Kubernetes:** solves a different, distributed problem and obscures the
  single-host use case.
- **Fleet manager first:** repeats the abandoned endpoint-management direction
  and competes on inventory rather than simple container convergence.
- **Multi-OS first:** forces weak abstractions over incompatible host primitives.

## Consequences

State, locking, failure recovery, and security can be designed locally. Future
multi-host administration must compose independent agents and cannot silently
turn the local engine into a scheduler.
