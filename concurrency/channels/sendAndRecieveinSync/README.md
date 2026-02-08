# Channel Send/Receive in Sync

Demonstrates two goroutines coordinating via an unbuffered channel and a WaitGroup. 
`printHello` sends integers 0–9 while printing "Hello"; 
`printWorld` receives those integers, prints "World <n>", and echoes them back to release the sender. The channel is closed by `printHello` after the loop so `printWorld` exits.

## How it works
- `printHello` loop (`for i := range 10`) runs `i = 0..9` (requires Go 1.22+ for `range` over ints), sending each `i` and waiting for an echo before continuing.
- `printWorld` ranges over the channel, prints the received `i`, then sends it back. When the channel is closed, the range ends and the goroutine returns.
- `sync.WaitGroup` waits for both goroutines to finish before `main` exits.

## Run
From this directory (module root must contain `go.mod`), run:

```bash
go run .
```

On older Go versions (<1.22), replace `for i := range 10` with a classic loop:

```go
for i := 0; i < 10; i++ {
    // ...
}
```

## Key points
- Unbuffered channels block the sender until the receiver receives (and vice versa), giving you synchronization for free.
- Always close the channel on the sending side when you are done to stop receiver ranges cleanly.
- Use the same `WaitGroup` (passed by pointer) so `Wait` sees `Done` calls from both goroutines.
