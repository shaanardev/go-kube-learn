# ScalingPolicy API design exercise

This directory is reserved for Go types describing a possible `ScalingPolicy` Kubernetes resource. `v1alpha1` names an early API version; it is not a program or a version of the controller. There is no implementation here yet.

Questions to explore before writing types:

- What does a policy describe, and what would be better left to existing Kubernetes resources?
- Which fields express intent (`spec`), and which report observations (`status`)?
- What makes a policy invalid? What should an omitted field mean?
- How could the definition change without confusing existing users or controllers?
- Which program, if any, should act on the policy and own its status?

Write a small example object and explain its meaning before choosing tooling. If you use Kubebuilder markers and generators, inspect the resulting types and CRD schema. The directory name does not prescribe a particular scaling algorithm or controller library.
