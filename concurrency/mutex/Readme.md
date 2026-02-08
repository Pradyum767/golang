## Mutexes in Go

### What and why
`sync.Mutex` is a mutual-exclusion lock that protects shared data so only one goroutine can access the critical section at a time. Pick mutexes for simple ownership of state; prefer channels when you want to model ownership/communication of the data itself.

### Types
- `sync.Mutex`: single reader/writer at a time (strict exclusion).
- `sync.RWMutex`: multiple concurrent readers allowed; writers still exclusive. Better for read-heavy workloads.

### Common methods
- `Lock()` / `Unlock()` for `Mutex` and writers of `RWMutex`.
- `RLock()` / `RUnlock()` for readers of `RWMutex`.

### Choosing between Mutex and RWMutex
| Feature       | sync.Mutex                    | sync.RWMutex                               |
| ------------- | ----------------------------- | ------------------------------------------ |
| Complexity    | Lower overhead                | Slightly higher (tracks reader count)      |
| Concurrency   | One goroutine at a time       | Many readers or one writer                 |
| Best for      | Short ops, frequent writes    | Read-heavy workloads, long reads           |
| Methods       | `Lock` / `Unlock`             | `RLock` / `RUnlock`, plus `Lock` / `Unlock` |

### Example
```go
var (
	mu  sync.RWMutex
	val int
)

func readVal() int {
	mu.RLock()
	defer mu.RUnlock()
	return val
}

func writeVal(v int) {
	mu.Lock()
	val = v
	mu.Unlock()
}
```