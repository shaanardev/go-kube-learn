# Executables

Each directory here represents a runnable Go program. The name `cmd` is a Go project convention: programs here can be interactive command-line tools or long-running services.

A program starts at `func main` in `package main`. Aim to keep startup understandable: read configuration, construct what the program needs, start its work, and handle shutdown. Let supporting packages emerge as you learn what needs separating; the directory layout does not require a framework or a particular set of layers.

## Planned exercises

- `log-parser`: streaming file parser and summary output
- `task`: small CLI with persisted task data
- `health-checker`: bounded concurrent HTTP health checks
- `pod-observer`: read-only Kubernetes pod monitor
- `opa-observer`: diagnose the OPA workload and create a redacted support bundle
- `opa-controller`: roll out validated OPA configuration and report progress
- `opa-loadlab`: run bounded load and soak measurements against a lab OPA service
- `opa-fleet`: audit OPA versions, health, and configuration across clusters
- `scaling-controller`: reconcile a custom scaling policy resource
- `autoscaler`: recommend a bounded scaling action before enabling writes
- `otelcol`: optional custom OPA health receiver after operating the standard Collector
