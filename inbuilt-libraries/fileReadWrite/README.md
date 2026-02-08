# File Read/Write Demo

Sample example showing JSON file read/write in Go. It loads `testRead.json` into a struct, prints the parsed data, then writes it back out as `testWrite.json`.

## What it does
- `processFileData` reads a JSON file with `os.ReadFile`.
- Unmarshals into `Payload`/`EngineersData` structs using `encoding/json`.
- Prints the result with `%+v` to include field names.
- Marshals the struct and writes `testWrite.json`.

## Run it

From this folder:
```bash
go run .
```

## File expectations
- Input: `testRead.json` in this directory.
- Output: `testWrite.json` will be created/overwritten. Use a reasonable mode (e.g., `0644`) if you adapt the code; on Windows the mode is mostly ignored.

## Structs
```go
type EngineersData struct {
    Name       string `json:"name"`
    Profession string `json:"profession"`
    Age        string `json:"age"`
}

type Payload struct {
    EngineersData EngineersData `json:"engineersData"`
}
```

## Tips
- If you hit permission errors (e.g., in OneDrive), try writing to a non-synced path or adjust file attributes.
- Swap `Age` to `int` in `EngineersData` if you want numeric age handling.
