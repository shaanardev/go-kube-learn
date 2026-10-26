# HTTP health checker exercise

Accept a list of HTTP endpoints and report their status within a fixed overall time budget. Bound concurrency and per-request timeouts. Do not make one slow endpoint stall all the others.

Learn: `net/http`, contexts, cancellation, goroutine ownership, result aggregation, and useful exit codes.
