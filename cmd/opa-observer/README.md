# OPA doctor and support-bundle exercise

Evolve this observer into an operator CLI with a `doctor` command. Collect OPA health, Deployment and Pod conditions, Service endpoints, recent events, and the active image/configuration revision. Offer concise human output and stable JSON output with documented exit codes.

Add an optional support bundle that packages selected logs and diagnostics. Bound collection time and archive size, redact credentials and sensitive data, and make included data clear to the operator. Keep it read-only.

Learn: Kubernetes API observation, OPA health endpoints, service discovery, deadlines, streaming I/O, archive formats, redaction, CLI contracts, and safe diagnostic output. Treat policy decision evaluation as a separate optional feature.
