# OPA configuration rollout controller exercise

A Kubernetes **controller** is a worker that keeps comparing what should be true with what it observes, then acts on any difference. This is **reconciliation**, rather than handling an HTTP request.

Here, the condition concerns the OPA image or configuration you want running. For example: you request revision B, but the workload still uses revision A. Your program could update the OPA Deployment and observe whether revision B becomes ready. Kubernetes' existing Deployment machinery handles replacing pods; your program would decide which revision is wanted and report progress toward it.

This directory is for the executable's startup. It does not imply all implementation belongs here, and it does not run anything in `api/`. A custom Kubernetes resource is one possible source of desired configuration; choosing how intent is supplied is part of the exercise.

Questions to explore:

- Where does the desired image or configuration come from, and how is it validated?
- How can you tell the difference between a requested revision and one that is actually serving?
- What should happen if the desired revision changes during a rollout?
- What happens after a failed update, a repeated notification, or a controller restart?
- Which changes belong to your program, and which belong to Kubernetes' existing controllers?

Observation is a useful first step. A preview of proposed changes can help you understand behavior before enabling writes. Aim for **idempotent** actions: repeating an attempt should not trigger unnecessary changes.

Possible topics and tools: custom resources, controller-runtime, revision identifiers, readiness, status conditions, API conflicts, retries, RBAC, and leader election. Let the behavior you want to understand guide the choice of tools. `envtest` can help explore API interactions; a disposable cluster lets you observe actual pod and rollout behavior. OPA policy evaluation remains an optional separate exercise.

See the [glossary](../../docs/glossary.md) for the terms used here.
