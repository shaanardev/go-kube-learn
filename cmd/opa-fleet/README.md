# OPA fleet inventory and drift auditor

Build a read-only CLI that inspects an explicit list of Kubernetes contexts and reports OPA version, readiness, and selected configuration revisions. Show partial results when one cluster is unreachable and clearly mark stale or unknown observations. Do not modify clusters.

Learn: per-cluster client configuration, bounded fan-out, context deadlines, cancellation, partial failure handling, caches, stable JSON output, and secret-safe diagnostics.
