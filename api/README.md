# Kubernetes API design

Here, **API** means the structure of objects that users and programs exchange with Kubernetes. This directory is intended for Go definitions of custom resource types. It is not an HTTP server or a collection of ASP.NET-style controllers. Think of the definitions as versioned data contracts.

For example, these are three separate things:

| Part | Role |
| --- | --- |
| A `ScalingPolicy` type defined here | Describes the fields a policy object can contain |
| A particular `ScalingPolicy` object stored in Kubernetes | Contains someone's requested settings and reported status |
| A controller program started from `cmd/scaling-controller/` | Reads objects and works to bring observed state toward the desired state |

The controller may import Go types from here; it does not run this directory. A **CustomResourceDefinition (CRD)** is the schema and registration that teaches the Kubernetes API server about a new resource type. The API server stores and serves its objects; your controller supplies the behavior. See the [glossary](../docs/glossary.md).

## Design exercise

A custom resource is one way to express desired state. Decide whether you need one, and what existing resources or simpler configuration sources could already express. Before choosing generation tools or writing API types, consider:

- the user-facing fields and validation rules
- defaulting and versioning expectations
- status conditions and what each condition means
- which controller owns each status field
- which Kubernetes resources the controller will read or write

The first candidate is a `ScalingPolicy` resource with explicit bounds and a pause/dry-run concept. The [versioned exercise brief](scalingpolicy/v1alpha1/README.md) suggests questions to explore. Kubebuilder is an available scaffolding tool when you reach this stage; learn what its generated files do before relying on them.
