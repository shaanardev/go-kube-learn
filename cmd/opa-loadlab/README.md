# OPA load and soak lab

Build a command that measures an OPA service in a disposable learning cluster. Require an explicit target and keep duration, request rate, and concurrency bounded. Report request counts, latency distribution, HTTP outcomes, and timeouts; support clean cancellation and a summary when interrupted.

Add scenarios for startup, slow responses, and service unavailability using controlled lab conditions. Do not turn this into an unrestricted load generator.

Learn: bounded worker pools, context cancellation, backpressure, reusable HTTP transports, streaming metrics, profiling, and runtime execution traces. Use its results to inform the autoscaler exercise rather than guessing at thresholds.
