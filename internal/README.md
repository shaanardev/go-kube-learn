# Private Go packages

Add packages here as implementation needs emerge. Keep package APIs small and domain-focused. A package should have a clear owner and purpose; avoid adding empty layers or sharing types only because two files happen to look similar.

Potential future areas include OPA transport and decision contracts, Kubernetes clients, configuration, health reporting, and telemetry setup. Keep executable wiring in `cmd/` and keep policy decisions in Rego where possible.
