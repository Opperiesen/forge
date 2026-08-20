# Architecture decision records

ADRs record decisions that would otherwise be expensive to rediscover. They are
append-only: a later decision supersedes an earlier ADR instead of rewriting its
history.

Use the next four-digit number and this structure:

```markdown
# NNNN — Decision title

- Status: proposed | accepted | superseded
- Date: YYYY-MM-DD

## Context
## Decision
## Alternatives considered
## Consequences
```

An ADR is required for a public contract, a new external dependency, persistent
state, security boundaries, runtime support, or a change to the MVP boundary.
