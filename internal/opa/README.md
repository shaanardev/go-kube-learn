# Optional OPA API integration package

Add a client only when a Go tool needs to inspect OPA health, status, or another API. Keep transport details, deadlines, and response decoding small and explicit. Decision evaluation is optional; preserve the distinction between a policy denial, undefined result, invalid response, and request failure if you add it.

Do not log credentials, full decision input, or sensitive policy data. Keep operation-specific fallback behavior with the caller.
