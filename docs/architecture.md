# Architecture notes

## Initial shape

```text
Kubernetes API <──> Go observers, controllers, autoscaler, and fleet auditor
                              │
                       observe or manage
                              ▼
                    OPA service <── OPA load lab
                              │
                         telemetry
                              ▼
                       OTel Collector
```

This is a learning direction, not a deployment diagram. OPA is the in-cluster reference workload. Go programs observe it and may later manage its configuration or scaling.

Here, a controller means a worker that reconciles desired and observed state which usually calls the Kubernetes API. The proposed programs and package boundaries are exercises to explore, rather than a requirement to deploy each responsibility as a separate service. See the [glossary](glossary.md) for the terminology.

## Boundaries to define as you implement

- **Workload contract:** record which OPA image/version, configuration, policy source, ports, probes, and resources your learning deployment uses. Make the active version and configuration observable.
- **Management scope:** start with observation. For changes, use dry-run or explicit bounds, validate inputs, and make rollout status visible before broadening what a controller can modify.
- **Kubernetes permissions:** start read-only. Add only the exact resource verbs needed by each observer or controller, and review generated RBAC before applying it.
- **Telemetry:** use stable, low-cardinality attributes. Never attach credentials, full policy input, or arbitrary object payloads to logs or spans.
- **Actions:** separate recommendation from mutation. Make controller writes idempotent, bounded, and visible through status conditions and metrics.
- **Optional decision integration:** if a Go service calls OPA's Data API, document the decision input/output and distinguish a denial, undefined result, invalid response, and request failure. Choose failure behavior for that specific operation.
- **Diagnostics:** keep profiling and debug handlers on a separate, protected interface. Support bundles should have size/time limits and redact secrets by default.
- **Load generation:** keep targets explicit and traffic bounded to the disposable learning environment. Make cancellation and maximum duration part of the tool contract.
- **Fleet views:** label cluster observations with freshness and partial-failure state; never present unavailable data as current.

## Package direction

Start with packages under `internal/` when there is real code to share, such as configuration parsing, OPA health/status clients, Kubernetes adapters, or telemetry setup. Keep each service's responsibilities local. Avoid a generic framework until at least two services demonstrate the same need.

Keep each executable's wiring in `cmd/<name>`. Avoid package-level mutable state. Pass `context.Context` to work that can be cancelled, and make resource ownership and shutdown order explicit.

## Deployment notes

See [deploy/kubernetes/README.md](../deploy/kubernetes/README.md) before creating manifests. This scaffold does not choose a cluster version, OPA image, policy source, or credentials mechanism on your behalf.
