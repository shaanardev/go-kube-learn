# Kubernetes integration package

Add shared Kubernetes client helpers only after a second use case appears. Keep RBAC assumptions visible and prefer typed, narrow operations over a generic client framework. Ensure API calls and watches stop when their context is cancelled.
