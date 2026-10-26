# Learning roadmap

Each project is an exercise. Implement one small behavior at a time, write down design choices, and make the next project reuse an idea you now understand. The sequence is adapted from the [Go project ideas list](https://dev.to/kumarkalyan/top-50-project-ideas-to-become-a-golang-hero-98k).

Use the [glossary](glossary.md) when terminology is unfamiliar. The milestones describe behaviors to understand; suggested tools are options to investigate as the need arises.

## 0. Set up and learn the Go toolchain

Get comfortable with `go version`, `go help`, modules, `gofmt`, `go doc`, and reading compiler errors. Use the Go tour and standard-library documentation as references. Keep this repo as one module for now.

## 1. Go fundamentals: small command-line tools

Suggested projects: log file parser, task manager CLI, system information command, HTTP status checker.

Practice:

- functions, structs, interfaces, slices, maps, and errors
- reading files and streams without loading everything into memory
- command arguments, configuration, and exit codes
- table-driven tests and testing error cases
- context-aware HTTP clients, request timeouts, and bounded concurrency

Milestone: each command has a clear input/output contract and useful errors; malformed input and network failures do not crash the process.

## 2. HTTP service and operational basics

Build a small REST API, then a service health checker that calls it. Use `net/http` first. Add structured logs, configuration validation, request IDs, graceful shutdown, and separate liveness from readiness.

Milestone: the service can start, report whether it is alive and ready, stop on cancellation, and explain why it cannot start.

## 3. Kubernetes observation and telemetry

Build a pod monitor that lists and watches a deliberately small set of Kubernetes resources. Begin with a read-only service account and a local development cluster. Add metrics and traces using OpenTelemetry APIs/exporters; learn the upstream Collector deployment before writing a custom Collector distribution.

Practice API clients, watches, retries with backoff, rate limits, RBAC, resource requests, and shutdown behavior. Make local logic testable without a cluster.

Milestone: the monitor has narrow permissions, bounded resource use, and useful signals when the API server is slow or unavailable.

## 4. Run OPA in Kubernetes and diagnose it

Deploy OPA in a disposable local cluster as a service you can observe and operate. Start with its Deployment, Service, health probes, resource profile, and configuration. Evolve `cmd/opa-observer` into a doctor CLI that reports pod health, endpoints, active revisions, and service availability. Add a bounded, redacted support bundle.

Practice Kubernetes service discovery, health versus readiness, rollout behavior, config delivery, resource requests and limits, and the difference between process health and serving the expected configuration. Add OPA decision calls only if you want a separate Rego exercise; the rest of this roadmap does not depend on them.

Milestone: you can deploy OPA, tell whether it is ready for its intended use, diagnose a failed or partial rollout, and explain the service's resource and configuration state.

## 5. Configuration rollout controller

Explore a worker that keeps the OPA workload aligned with a desired configuration or image revision and reports progress. This repeated comparison and action is reconciliation; it is what “controller” means here. `cmd/opa-controller` is the executable entry point, while `api/` could define custom resource types if you choose that way of expressing desired state. Decide where intent comes from and which behavior belongs to your worker versus Kubernetes' existing controllers. Observation and a preview of changes are useful steps before enabling writes. Kubebuilder/controller-runtime are possible tools once their purpose is clear.

Practice API versioning, status conditions, API conflict handling, bounded retries, least-privilege RBAC, leader election, predicates, cache behavior, and controller metrics. Use `envtest` for API and reconciliation behavior, then a disposable cluster for pod scheduling, Services, and rollout behavior.

Milestone: repeated reconcile produces the same result, invalid specs give actionable status, and the controller only has permissions it needs.

## 6. Load and soak measurement

Build `cmd/opa-loadlab` to generate bounded, time-limited traffic against the OPA service in your disposable cluster. Measure latency and failures under normal load, startup, and controlled unavailability. Make cancellation work and report partial measurements on shutdown.

Practice bounded goroutine lifetimes, backpressure, HTTP connection reuse, latency distributions, and profiling. Keep rate, concurrency, and duration limits explicit.

Milestone: you can explain the workload and limits behind each measurement and identify the service's bottleneck with profiles or traces.

## 7. Autoscaler as a staged project

Start with an observer that computes and reports a recommendation for the OPA Deployment. Compare it with Kubernetes' built-in scaling mechanisms. Then add a dry-run mode and explicit bounds before considering any write path. Do not begin by changing replica counts automatically.

Practice noisy signals, stabilization windows, cooldowns, min/max bounds, stale metrics, partial failures, and an operator-controlled pause. Keep measurement and actuation logic in Go. Using OPA to approve proposed changes is optional.

Milestone: recommendations are explainable and repeatable; missing or stale metrics cannot trigger an unbounded scale action.

## 8. Multi-cluster inventory and drift

Build `cmd/opa-fleet` as a read-only auditor over an explicit set of Kubernetes contexts. Report OPA versions, readiness, and configuration revisions, while preserving useful results when one cluster is unavailable.

Practice bounded fan-out, per-cluster deadlines, cancellation, partial failures, stale-cache indicators, and stable machine-readable output. Do not add automated fleet changes in this exercise.

Milestone: operators can distinguish current, stale, unavailable, and inconsistent observations across clusters.

## 9. OpenTelemetry Collector extension

Operate the standard Collector first and collect signals from OPA and your Go services. If a concrete gap remains, use the Collector Builder to assemble a small custom distribution, for example with a receiver that polls OPA health. Keep collection, transformation, and export behavior observable and configurable.

Milestone: you can explain why a custom component is necessary, what data it handles, and how it behaves under backpressure and exporter failure.

## Suggested learning loop

For each project, write a short design note, implement the smallest vertical slice, test normal and failure paths, inspect logs/metrics, and write a brief retrospective. Keep the Go code yours; use generated code only when the Kubernetes API tooling genuinely requires it, and inspect the generated result.

For advanced language and production topics, follow [docs/go-feature-labs.md](go-feature-labs.md) alongside these projects.
