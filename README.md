# Go Networking Projects

A collection of hands-on networking projects built while working through **Black Hat Go** — chapter by chapter. Each sub-project maps to a chapter's core concepts and is implemented progressively from a naive approach to a production-quality design.

---

## Projects

### `tcp-scanners-proxies`

Covers TCP fundamentals: port scanning with increasing concurrency models, raw TCP proxy forwarding, and Netcat-style remote command execution.

#### Port Scanner — 5 progressive iterations

| Step | Approach | Key Concept |
|------|----------|-------------|
| 1 | Sequential scan (ports 1–1024) | `net.Dial` basics |
| 2 | Naive goroutine per port | Goroutine spawning (race condition) |
| 3 | Synchronized goroutines | `sync.WaitGroup` |
| 4 | Worker pool | Buffered channels + WaitGroup |
| 5 | Multichannel worker pool *(active)* | Dual channels, sorted output |

#### TCP Proxy

- Custom `io.Reader` / `io.Writer` implementations
- Echo server using `bufio` and `net.Conn`
- Port-forwarding proxy via bidirectional `io.Copy`
- Netcat-style bind shell (`/bin/sh -i`) with flushed stdout via a custom `Flusher` wrapper

**Run:**
```bash
cd tcp-scanners-proxies
go mod init tcp-scanners-proxies
go run cmd/portScanner/main.go
go run cmd/tcpProxy/main.go
```

---

### `HTTP-CLIENTS-REMOTE-INTERACTIONS`

*(In progress — HTTP client interactions and remote API calls.)*

---

## Setup

Each sub-project is a standalone Go module. No external dependencies are required for current chapters — the standard library (`net`, `sync`, `io`, `os/exec`) covers everything.

```bash
cd <sub-project>
go mod init <sub-project-name>
go run cmd/<entry>/main.go
```
