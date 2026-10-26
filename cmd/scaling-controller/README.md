# Scaling policy controller exercise

A Kubernetes **controller** is a worker that compares desired state with observed state and works to close the gap. That repeated process is **reconciliation**.

This exercise explores managing a workload according to a scaling policy. The proposed [`ScalingPolicy` API](../../api/scalingpolicy/v1alpha1/README.md) would describe the policy's fields; this executable would interpret policy objects and observe or act on a workload. It can import those types, but does not run the API directory.

For example, a policy might allow a range of replica counts. The range alone does not say when to scale: deciding what additional information and behavior are needed is part of the exercise.

Questions to explore:

- What desired condition is this controller responsible for maintaining?
- How does it discover relevant changes, and what should cause it to check again?
- What should happen when the policy is invalid, observations are stale, or the program restarts?
- What happens when another controller also changes the replica count?
- How will you explain what it observed and why it chose an action?

Observation and status reporting provide a small starting point before enabling changes. Aim for **idempotent** actions: retrying the same work should not create extra effects.

Possible topics and tools: watches, status conditions, RBAC, shutdown, leader election, and controller-runtime/Kubebuilder. Choose tools as their purpose becomes clear. Compare this exercise with the [autoscaler](../autoscaler/README.md), which focuses on deriving a scaling recommendation from measurements; decide how their responsibilities would relate before running both against one workload.

See the [glossary](../../docs/glossary.md) for the terms used here.
