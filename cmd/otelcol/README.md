# Custom OpenTelemetry Collector exercise

First deploy and configure the standard Collector. If a specific gap remains, use the OpenTelemetry Collector Builder to assemble a small custom distribution with one component. A useful candidate is a receiver that polls OPA's health endpoint and emits readiness and availability telemetry.

Learn: receivers, processors, exporters, pipelines, backpressure, component lifecycle, configuration validation, and telemetry cardinality. Avoid maintaining a custom distribution without a concrete requirement.
