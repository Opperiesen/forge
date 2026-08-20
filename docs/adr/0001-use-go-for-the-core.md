# 0001 — Use Go for the Forge core

- Status: accepted
- Date: 2026-08-20

## Context

Forge needs a CLI, a long-running agent, concurrent reconciliation, Linux system
integration, OCI ecosystem libraries, and distributable binaries. Avoiding a
future whole-project migration matters more than language novelty.

## Decision

Implement the CLI, agent, reconciliation engine, runtime adapters, and any future
control-plane service in Go. The initial module targets Go 1.26 and uses the
standard library until an external dependency has a concrete MVP requirement.

Keep the OCI and runtime boundary language-neutral. A future low-level component
may use Rust or C without moving the product core away from Go.

## Alternatives considered

- **Full Rust:** strong low-level safety, but higher implementation friction for
  the control-plane and cloud-native integration work that dominates the MVP.
- **Mixed Go and Rust from day one:** adds build and interface complexity before
  a low-level component has justified it.
- **Python:** productive for prototypes, but less suitable for the single-binary
  agent and systems lifecycle expected here.

## Consequences

The project follows familiar cloud-native conventions and can ship static-like
Go binaries. Contributors need one primary toolchain. Go does not remove the need
for careful privilege, process, filesystem, and concurrency design.
