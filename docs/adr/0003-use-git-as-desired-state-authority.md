# 0003 — Use Git as the desired-state authority

- Status: accepted
- Date: 2026-08-20

## Context

Forge promises reviewable, reproducible changes and deterministic rollback. A
mutable runtime database cannot provide that contract by itself.

## Decision

An operator configures one Git repository and a revision policy. Forge resolves
the mutable reference to an immutable commit before reading configuration. Every
plan, apply, status record, and rollback refers to a commit ID.

Local state records observations and recovery metadata but never becomes a
second desired-state authority. An emergency local rollback pauses automatic
reconciliation until Git and the selected local revision agree again.

## Alternatives considered

- **Local configuration files:** simple, but no built-in review or remote source
  of truth.
- **Central database:** requires a control plane and weakens repository-native
  workflows.
- **Automatic writes back to Git:** requires credentials and policy decisions
  that are outside the MVP.

## Consequences

Network or Git failure cannot invalidate the currently running workload. Forge
must define reference resolution, caching, authentication, and divergence
clearly. Rollback semantics remain visible instead of silently fighting Git.
