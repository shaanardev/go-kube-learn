# Advanced Go feature labs

Use these topics inside the OPA observer, controller, load lab, fleet auditor, and Collector work. The goal is to recognize where a feature helps in production and where the simpler approach is better.

## Concurrency, cancellation, and backpressure

In the load lab and fleet auditor, use bounded workers, per-operation deadlines, cancellation, and explicit partial-failure reporting. Make goroutine ownership and shutdown behavior easy to follow. Prefer a small worker limit or semaphore over one goroutine per cluster/request with no bound.

## Streaming and JSON

Use `io.Reader`/`io.Writer` for logs and support bundles so memory use stays bounded. With Go 1.27, compare `encoding/json/v2` for OPA responses or configuration inputs and explore `encoding/json/jsontext` for streaming token-level processing. Check compatibility before changing a wire format: the original `encoding/json` API remains supported, while v2 has stricter defaults such as rejecting duplicate object names and invalid UTF-8. See the [Go 1.27 release notes](https://go.dev/doc/go1.27) and [JSON v2 migration guide](https://go.dev/doc/jsonv2-migration).

## Generics, methods, and iterators

Use Go 1.27 generic methods or the standard `iter.Seq` APIs for a small typed inventory or streaming adapter. Compare that API with a concrete implementation and keep the one that is easier to maintain. Do not add generic layers to ordinary Kubernetes resource handling just to use generics. See [generic methods](https://go.dev/blog/generic-methods) and the [`iter` package](https://pkg.go.dev/iter).

## Profiling, execution tracing, and PGO

Use `opa-loadlab` to create a repeatable workload, then inspect CPU, heap, goroutine, block, or mutex profiles and Go execution traces. Keep profiling endpoints off the public Service and account for profiling overhead. Consider profile-guided optimization only after collecting representative workload profiles; Go's PGO guidance recommends profiles that reflect real application behavior. See [Go diagnostics](https://go.dev/doc/diagnostics) and [PGO](https://go.dev/doc/pgo).

## Fuzzing, race detection, and dependency health

Fuzz parsers for configuration, JSON, and support-bundle metadata. Run concurrent paths under the race detector. Add `go vet` and `govulncheck` to a later CI/release workflow. These catch different classes of problems; none replaces careful review. See [Go security best practices](https://go.dev/doc/security/best-practices).

For Kubernetes controllers, use `envtest` to exercise API and reconcile behavior. It starts an API server and etcd but does not run kubelet or built-in controllers, so use a disposable full cluster for pod scheduling, Service routing, and end-to-end rollout behavior. See [Kubebuilder envtest](https://book.kubebuilder.io/reference/envtest).

## Code generation and release builds

Use `go:generate` or the relevant upstream builder for CRD/client code and a custom Collector distribution only when the project reaches that stage. Read and review generated output. Add a release workflow that records version/build metadata, builds Linux amd64 and arm64 images, and runs static and vulnerability checks.

## Niche features: learn in isolation, adopt by evidence

- **`sync.Pool`:** try it only if heap profiles show short-lived allocation pressure. It is for temporary reusable objects; the runtime may remove pooled items at any time, so it is not a cache. See [`sync.Pool`](https://pkg.go.dev/sync#Pool).
- **`cgo` and `unsafe`:** make a small separate experiment to understand native-library integration, memory ownership, and build constraints. Keep them out of core services unless an actual native dependency or measured bottleneck requires them. See the [cgo documentation](https://go.dev/cmd/cgo/) and [`unsafe`](https://pkg.go.dev/unsafe).
- **Cross-compilation:** build the same command for Linux amd64 and arm64, and document when a native dependency changes the build requirements. Keep normal service binaries pure Go when that meets the need.
