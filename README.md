# anyway

A small visual novel engine for Go, built on [Ebitengine](https://ebitengine.org).
Scripts are plain YAML: a list of commands the engine runs one by one.

```go
scene, err := anyway.Load("prologue.yaml")
if err != nil {
	log.Fatal(err)
}
anyway.Run(scene, anyway.Config{Title: "prologue"})
```

```yaml
scene_id: 0
scene_name: "prologue"
commands:
  - type: "TEXT"
    text: "Hello."
  - type: "CLEAR"
  - type: "TEXT"
    text: "Anyway."
```

Click, Enter or Space finishes the current line, then continues.

## Commands

| type    | fields  | status          |
|---------|---------|-----------------|
| `TEXT`  | `text`  | works           |
| `CLEAR` |         | works           |
| `BG`    | `image` | parsed, not run |
| `BGM`   | `audio` | parsed, not run |
| `SFX`   | `audio` | parsed, not run |

Unknown types and missing fields are rejected when the script loads.

## Try it

```
go run ./examples/prologue
go test ./...
```

## Status

Early. The API will change.

## License

MIT
