# anyway

Small visual novel engine for Go on Ebitengine. Scripts are YAML command lists.
Module: `github.com/Err-Nil-Lab/anyway`. Public repo, MIT.

## Layout

- `script.go`: `Scene`/`Command`, `Load`, `Parse`, validation.
- `textbox.go`: `wrapText` and the typewriter `textbox` (rune-based).
- `player.go`: runs commands; no renderer dependency, so it is testable.
- `game.go`: Ebitengine `Run`, `Update`, `Draw`, `Layout`. The only file that touches the window.
- `examples/prologue/`: runnable example, script embedded with `go:embed`.

## Rules

- Package lives at the repo root. No `engine/` subfolder.
- Everything in the repo (code, comments, scripts, tests, commit messages) is English.
- Keep the repo minimal: no teaching docs or notes here.
- Run `gofmt -l .` and `go test ./...` before committing.
- Do not push or create remote resources without the owner's explicit go-ahead.

## Status

`TEXT` and `CLEAR` work. `BG`, `BGM`, `SFX` are parsed and validated but not executed yet;
implementing them is the next step (start with `BG`).
