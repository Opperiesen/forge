# 0004 — Keep the OCI runtime interchangeable

- Status: accepted
- Date: 2026-08-20

## Context

Building a low-level runtime requires correct namespaces, cgroups, capabilities,
seccomp, filesystems, and lifecycle handling. That work does not validate the
Forge product proposition. Coupling to one complete engine would instead make
Forge a thin wrapper around that engine's private behavior.

## Decision

Forge will consume OCI-compatible artifacts and define a narrow internal runtime
adapter. The MVP will select one reference adapter after a vertical prototype.
It will not implement an OCI runtime from scratch.

The adapter contract must explicitly cover the operations OCI does not define,
including image transport, networking, storage, logs, and restart behavior.

## Alternatives considered

- **Build a runtime immediately:** high security and correctness risk with no
  evidence that it differentiates the product.
- **Hard-code Podman or Docker:** fastest prototype, but risks inheriting a large
  external CLI contract as Forge's architecture.
- **Support several runtimes immediately:** multiplies integration tests before
  the contract is known.

## Consequences

One implementation will be supported first, while the boundary remains testable
and replaceable. Interchangeability is an architectural property, not a promise
that every runtime works in the first release.
