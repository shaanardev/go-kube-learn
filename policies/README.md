# Optional OPA policy exercise

OPA's policy features are available for a separate Rego learning track, but they are not required for the Kubernetes projects in this repo. If you choose that track, write Rego policies here by hand, organize modules by decision domain, and document the input/output contract alongside them.

Suggested first exercise: define one named decision, list the exact input fields it needs, decide whether an undefined result is an error, and try allow, deny, malformed-input, and missing-input cases.

For service integration, see [docs/architecture.md](../docs/architecture.md). OPA's [Data API](https://www.openpolicyagent.org/docs/rest-api) evaluates named decisions over HTTP; OPA's [integration guide](https://www.openpolicyagent.org/docs/integration) also describes the Go SDK and other integration choices.
