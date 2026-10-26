# Short glossary

These meanings apply to this learning lab. Some words, especially **controller** and **API**, have different meanings in other frameworks.

## Programs and directories

| Term | Meaning here |
| --- | --- |
| **Executable / command** | A runnable program. It may finish after a task or keep running as a service. |
| **`cmd/`** | A conventional Go directory for executable entry points. Similar in purpose to keeping each .NET application's `Program.cs` together with its startup code. It does not mean only interactive CLI tools. |
| **Package** | Go source files in one directory compiled together. A program imports supporting packages; it does not run their directories. |
| **`internal/`** | Go packages with language-tool-enforced import restrictions: code outside the directory tree rooted at `internal`'s parent cannot import them. Here, this keeps implementation packages private to this project. |
| **`api/`** | This project's location for planned Kubernetes resource type definitions, roughly like versioned data contracts. The name itself has no special meaning to Go. |

## Kubernetes objects and behavior

| Term | Meaning here |
| --- | --- |
| **Kubernetes API / API server** | The interface programs use to read and change Kubernetes objects, and the server that exposes it. A controller is usually a client of this API. |
| **Resource / object** | A resource is a kind of thing exposed by the API, such as Deployments. An object is a particular instance, such as the Deployment named `opa`. |
| **Custom resource / CRD** | A custom resource adds a resource type of your own. A CustomResourceDefinition registers that type and its schema with Kubernetes. Registration lets the API server store and serve objects; it does not implement a controller's behavior. |
| **`spec` / desired state** | The settings or intent an object declares: what should be true. A replica count is one example. |
| **`status` / observed state** | Reported observations about an object: what the responsible components have seen. Reports can lag behind reality. |
| **Controller** | A worker that repeatedly compares desired state with observed state and acts to bring them closer. Think of a .NET `BackgroundService` with that job. An ASP.NET controller instead handles HTTP requests and responses. |
| **Reconcile / reconciliation** | An attempt to bring observed state toward desired state. It may involve reading, changing something, reporting status, or finding nothing to do. Further attempts happen as needed. |
| **Idempotent** | Repeating the same operation has the same intended effect as doing it once. A repeated reconciliation should not create duplicate resources or unnecessary rollouts. |
| **Watch** | An API mechanism for receiving notifications of changes to resources. A notification is a reason to check state again. |
| **Workload / Pod / Deployment** | A workload is an application running in the cluster. A Pod runs one or more containers together. A Deployment declares and manages replicated Pods through Kubernetes' existing controllers. |
| **Rollout / readiness** | A rollout moves a workload to a requested version or configuration. Readiness indicates whether an instance can currently serve its intended traffic. An accepted update does not mean the rollout has finished. |
| **Status condition** | A named status report, often with a true, false, or unknown value and a reason, such as whether a workload is ready. |
| **RBAC** | Role-based access control: rules determining which Kubernetes API actions an identity may perform. |

## This lab's names

| Term | Meaning here |
| --- | --- |
| **OPA** | Open Policy Agent, used here as a service to deploy, observe, and manage. Writing policy rules is optional. |
| **`ScalingPolicy` / `v1alpha1`** | A proposed custom resource type and an early version of its API contract. A policy object contains data; a separate program gives that data operational meaning. |
| **Observer / controller / autoscaler** | An observer reports what it sees. A controller works toward desired state. An autoscaler derives scaling decisions from signals and may apply them. These roles can overlap; separate exercise directories do not require separate production services. |
| **Dry-run** | A way to explore a proposed change without applying it. A program printing its plan and the Kubernetes API validating a request without persisting it are different forms; define what your exercise means. |
| **controller-runtime / Kubebuilder** | Optional Go tooling: controller-runtime provides libraries for controllers; Kubebuilder scaffolds projects and API-related files. Neither tool defines what your controller should do. |

For more detail, see the [Kubernetes controller explanation](https://kubernetes.io/docs/concepts/architecture/controller/) and [Go module layout guidance](https://go.dev/doc/modules/layout). Start with the [executable overview](../cmd/README.md) or [API overview](../api/README.md) to connect the terms to this repository.
