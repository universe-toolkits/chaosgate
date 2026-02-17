<img align="center" alt="ChaosGate" src="https://github.com/user-attachments/assets/480e642c-1042-415e-9cd6-31ed4f1c1b1e" />


# chaosgate
ChaosGate is a programmable egress proxy for staging and QA environments that enables controlled chaos engineering and full-featured HTTP mocking against external dependencies.

---
## Features

It sits between your services and third-party providers (via DNS/extra_hosts override) and allows you to:

- Intercept outbound HTTP/HTTPS traffic
- Route requests dynamically based on host, path, method, headers, or body
- Inject failures (503, 401, timeouts, latency, connection drops, percentage-based errors)
- Simulate network instability and provider outages
- Return fully custom mock responses (not just errors)
- Mutate or override real upstream responses
- Control everything at runtime through a web UI
- ChaosGate combines:
- Reverse proxy
- Fault injection engine
- HTTP mock server
- Runtime rule engine

It is designed for real-world staging environments where QA engineers and developers need to test system resilience against unreliable or misbehaving third-party services — without modifying application code.


Use ChaosGate to validate retries, fallback logic, circuit breakers, schema handling, and overall robustness before production.
