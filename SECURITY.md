# Security policy

## Project maturity

Forge is pre-alpha and does not yet manage containers. There are no supported
production releases. Security properties described here are design requirements,
not claims about an implemented system.

## Reporting a vulnerability

Do not open a public issue for a suspected vulnerability. Use GitHub's private
vulnerability reporting for this repository. Include affected revision, impact,
reproduction steps, and any suggested mitigation. Acknowledgement and disclosure
timelines will be agreed in the private report while the project is pre-alpha.

## Initial threat model

Forge will run on a Linux host with enough authority to control container
resources. A compromised source repository, forged OCI artifact, unsafe manifest,
runtime socket, or over-privileged agent can therefore compromise the host.

The MVP must treat these as separate trust boundaries:

- Git transport, credentials, and immutable revision resolution;
- manifest parsing, schema validation, and dangerous host options;
- OCI artifact provenance and digest verification;
- agent privileges and access to the runtime boundary;
- host mounts, devices, capabilities, namespaces, and networking;
- durable state integrity and operation recovery;
- logs and status exposed to local users.

## Required defaults

- Prefer images pinned by digest; never silently replace a resolved digest.
- Reject privileged containers, host namespaces, devices, and arbitrary host
  mounts until an explicit policy is designed and documented.
- Keep Git and registry credentials outside manifests and persisted status.
- Redact credential material and secret-like values from errors and logs.
- Validate the complete desired state before any mutation.
- Refuse mutation when durable operation state cannot be written safely.
- Run the agent with the least authority supported by the chosen runtime model.
- Never claim rollback can undo application-level external side effects.

Secret distribution is outside the MVP. A future integration must reference an
external secret source without committing secret values to Git.
