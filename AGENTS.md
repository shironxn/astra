# AGENTS.md

## Project

**Astra** — A REST API (authors + books). No framework — stdlib `net/http` with Go 1.22+ method+path routing (e.g. `"GET /authors/{id}"`).

## Entrypoint

`cmd/main.go` — manual dependency injection wiring repo → service → handler.

## Architecture (clean)

```
handler/ → service/ → repository/ → MySQL
```

- Interfaces live in `internal/domain/` alongside models. Each layer depends on the domain interface, not the concrete type.
- `internal/config/server.go` registers routes on `http.NewServeMux()` and starts the server.

## Database

- Raw SQL (no ORM). Driver: `github.com/go-sql-driver/mysql`.
- Migrations via **goose**, stored in `internal/config/db/migrations/`.
- DSN uses `?parseTime=true` so `TIMESTAMP` columns scan into `time.Time`.

## Commands

| Command | Notes |
|---|---|
| `make run` | `go mod tidy` → build → execute `bin/main` |
| `make dev` | Air hot-reload (no `.air.toml`, flags inlined in Makefile) |
| `make build` | `go build -o bin/main ./cmd` |
| `make test` | `go test -v ./...` — no tests exist yet |
| `make migration-up/down/reset/status` | Requires `goose` CLI installed |


## Known bugs / quirks

- `internal/handler/book.go:64,69` — `GetById` and `GetAll` are missing `return` after `http.Error()`; execution falls through.
- `internal/config/db/mysql.go:11` — struct name `Databse` (typo).
- `internal/handler/book.go:35` — `BookHandler.Create` returns `200 OK`, should be `201 Created` like `AuthorHandler`.
- No test files exist anywhere in the repo.

## Env vars (`.env`)

| Var | Default | Note |
|---|---|---|
| `APP_HOST` | `localhost` | |
| `APP_PORT` | `8080` | |
| `DB_USER` | `root` | Must have CREATE rights for migrations |
| `DB_PASS` | | |
| `DB_HOST` | `localhost` | |
| `DB_PORT` | `3306` | |
| `DB_NAME` | `bookstore` | Must exist before running migrations |

## Prerequisites

- Go 1.22.4+
- MySQL running
- `goose` CLI (`go install github.com/pressly/goose/v3/cmd/goose@latest`)
- `air` CLI (`go install github.com/cosmtrek/air@latest`) — optional, for `make dev`
