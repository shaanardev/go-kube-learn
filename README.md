# Go + Kubernetes learning lab

A hands-on monorepo for learning Go and Kubernetes by building small tools and in-cluster services around Open Policy Agent (OPA) as a reference workload.

The repository is intentionally a scaffold. It contains no service implementation or generated controller code: you write those by hand as you work through the roadmap.

## Learning path

The sequence adapts the [50 Go project ideas](https://dev.to/kumarkalyan/top-50-project-ideas-to-become-a-golang-hero-98k) to a Kubernetes and in-cluster service focus:

1. Learn Go with a log parser, task CLI, and HTTP status checker.
2. Build a small HTTP service and a service health checker.
3. Observe Kubernetes pods and export useful telemetry.
4. Deploy OPA in a cluster, then build Go tooling to observe and manage that service.
5. Build a controller that rolls out OPA configuration and reports progress.
6. Measure OPA with a bounded load/soak tool, then build an autoscaler in stages.
7. Audit OPA versions and configuration across clusters.
8. Observe OPA and your Go services with OpenTelemetry; build a custom Collector component only if you find a concrete gap.

See [docs/learning-roadmap.md](docs/learning-roadmap.md) for the skills and milestones for each stage.
See [docs/go-feature-labs.md](docs/go-feature-labs.md) for advanced Go features and production workflows to practice through these projects.
See the [short glossary](docs/glossary.md) for unfamiliar Go and Kubernetes terms

## OPA's role in this lab

OPA is an in-cluster service to deploy, observe, configure, and use as a realistic workload while learning Kubernetes. The Go projects can inspect its health, watch its pods and configuration, report its telemetry, and eventually manage or scale it. You do not need to use OPA as a policy decision point for these projects.

Useful references:

- [OPA integration options](https://www.openpolicyagent.org/docs/integration)
- [OPA REST API](https://www.openpolicyagent.org/docs/rest-api)
- [OPA operations and readiness](https://www.openpolicyagent.org/docs/operations)

## Repository shape

```text
cmd/                         Executable-specific learning briefs
internal/                    Planned private Go packages; add code when needed
api/                         Planned Kubernetes API and CRD design
deploy/kubernetes/           Deployment notes and hardening checklist
policies/                    Optional Rego learning area
docs/                        Architecture and learning roadmap
```

`cmd/` groups runnable programs, including background services. `api/` describes planned Kubernetes resource types, rather than HTTP request handlers. A controller program may import those types; it does not run the `api/` directory. These directory names express project conventions, not Go language requirements. See [cmd/README.md](cmd/README.md) and [api/README.md](api/README.md) for more context.

This starts as one Go module. Keep services in the same module while they share a learning lifecycle and release process; split into multiple modules only if independent versioning or dependency boundaries become a real need. Add a package only when you have a concrete implementation to put there.

## Working style

- Write the implementation yourself. Use the briefs as prompts, not as hidden specifications.
- Treat suggested tools and milestones as starting points for exploration. Explain your choices; the briefs are not a required implementation recipe.
- For each project, sketch inputs, outputs, failure cases, and the smallest useful interface before coding.
- Prefer the standard library until a project has a clear need for an external dependency.
- Keep context cancellation, bounded timeouts, explicit errors, and graceful shutdown in mind from the first server.
- Write unit and integration tests as part of each learning project; keep Kubernetes integration work separate from fast local checks.
- Record what you learned and the next small step in your own notes.

See [docs/go-quality.md](docs/go-quality.md) for the advisory Go lint and static-analysis checks used on pull requests.

## Start here

1. Install a supported Go toolchain and read [the Go installation guide](https://go.dev/doc/install).
2. Read [docs/learning-roadmap.md](docs/learning-roadmap.md), then start with the log parser.
3. Before adding Kubernetes dependencies, choose the Kubernetes version you will run locally and check compatibility for the libraries you choose, such as `client-go` or `controller-runtime`.
4. Read [docs/architecture.md](docs/architecture.md) before building a controller or tool that changes the OPA workload.
