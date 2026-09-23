# AGENTS: scrollbacktamer (Go)

CLI tool for editing your scrollback buffer in Tmux. See `README.md` for
user-facing usage.

## Build & Test Commands

Prefer the `make` targets over invoking `go` directly.

- Build: `make build` (binary `scrollbacktamer`; embeds version/commit/date ldflags).
- Test: `make test` (`go test -v ./...`).
- Coverage: `make test-coverage` (generates `coverage.html`).
- Install: `make install` (`go install` into `$GOBIN`).
- Clean: `make clean`.
- Nix dev shell: `nix develop` (also `shell.nix`/`default.nix` for legacy nix).

## Layout

- `main.go` — entrypoint, command parsing, scrollback editing.
- `cli/` — CLI argument parsing.
- `lib/` — core logic: editor, settings, stream, tty integration.
- `lib/tty/` — tmux-specific TTY integration.
- Module path: `github.com/napisani/scrollbacktamer`.

## Code Style & Conventions

- **Language**: Go. Standard `gofmt` formatting.
- **Errors**: return explicit errors; no panics.
- **Testing**: unit tests in `lib/` with `editor_test.go`, `settings_test.go`,
  `stream_test.go`, `tty/tmux_test.go`.

## Release

- `.goreleaser.yaml` drives release builds; `.drone.yml` is the upstream CI (in
  the standalone repo).
- This subproject is auto-published from the monorepo to
  [`napisani/scrollbacktamer`](https://github.com/napisani/scrollbacktamer) —
  work in the monorepo, never commit to the published repo directly (see root
  `AGENTS.md`).