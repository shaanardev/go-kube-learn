# Kubernetes deployment learning area

Add manifests only after a service has a documented runtime contract. Choose a local Kubernetes version first, then match the Kubernetes Go libraries and controller tooling to it.

For each workload, work through these items and record the reason for each setting:

- dedicated service account and least-privilege RBAC
- non-root execution, read-only root filesystem where practical, dropped capabilities, and seccomp profile
- CPU/memory requests and limits based on observation
- startup, readiness, and liveness probes with distinct meanings
- graceful termination and an adequate termination grace period
- network access limited to required services
- profiling/debug endpoints kept on a separate protected interface rather than exposed by the normal Service
- policy and credentials delivery with rotation in mind
- replica count, disruption behavior, and failure-domain assumptions
- resource quotas and namespace boundaries for experiments

For OPA, readiness should reflect the policy-loading model you choose. OPA's health endpoint can include bundle activation checks; a process being alive does not by itself prove the required policy is ready. Keep OPA's management endpoints protected and expose only the API surface that callers need.

This directory intentionally contains no generated deployment or RBAC manifests yet. Create and review them as part of the relevant exercise, and apply them only to a disposable learning cluster until their permissions and effects are understood.
