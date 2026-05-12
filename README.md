# Astra

[![Go](https://img.shields.io/badge/Go-1.22.4-%2300ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

A REST API for managing authors and books, built with Go following Clean Architecture principles. No framework — uses stdlib `net/http` with Go 1.22+ method+path routing.

<img src="assets/thumbnail.gif" alt="Astra Thumbnail" width="100%">

## Tech Stack

- **Go** 1.22.4+ — stdlib `net/http`, `database/sql`
- **MySQL** — via `go-sql-driver/mysql`
- **godotenv** — environment variable
- **goose** — database migrations
- **air** — live-reload (optional, for development)

## Prerequisites

- Go 1.22.4+
- MySQL running
- `goose` CLI: `go install github.com/pressly/goose/v3/cmd/goose@latest`
- `air` CLI (optional): `go install github.com/cosmtrek/air@latest`

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/shironxn/astra
   cd go-clean-arch
   ```

2. Set up environment variables:

   ```bash
   cp .env.example .env
   ```

3. Create the database:

   ```sql
   CREATE DATABASE bookstore;
   ```

4. Run migrations:

   ```bash
   make migration-up
   ```

5. Run the application:

   ```bash
   make run
   ```

## Environment Variables

| Variable   | Default     | Description                      |
| ---------- | ----------- | -------------------------------- |
| `APP_HOST` | `localhost` | Server host                      |
| `APP_PORT` | `8080`      | Server port                      |
| `DB_USER`  | `root`      | Database user (must have CREATE) |
| `DB_PASS`  |             | Database password                |
| `DB_HOST`  | `localhost` | Database host                    |
| `DB_PORT`  | `3306`      | Database port                    |
| `DB_NAME`  | `bookstore` | Database name (must exist)       |

## Commands

| Command                 | Description                          |
| ----------------------- | ------------------------------------ |
| `make run`              | Tidy, build, and run the application |
| `make dev`              | Run with Air (hot-reload)            |
| `make build`            | Build binary to `bin/main`           |
| `make test`             | Run all tests (`go test -v ./...`)   |
| `make tidy`             | Run `go mod tidy`                    |
| `make clean`            | Remove `bin/` and `tmp/` directories |
| `make migration-up`     | Apply pending migrations             |
| `make migration-down`   | Roll back the last migration         |
| `make migration-reset`  | Roll back all migrations             |
| `make migration-status` | Show migration status                |
| `make help`             | Display all available commands       |

## Project Structure

```
.
├── cmd/
│   └── main.go              # Entrypoint — dependency injection wiring
├── internal/
│   ├── config/
│   │   ├── app.go           # Env loading & config structs
│   │   ├── server.go        # Router setup, route registration, server start
│   │   └── db/
│   │       ├── mysql.go     # MySQL connection
│   │       └── migrations/  # Goose migration files
│   ├── domain/
│   │   ├── author.go        # Author model + interfaces (repo, service, handler)
│   │   └── book.go          # Book model + interfaces
│   ├── handler/
│   │   ├── author.go        # HTTP handlers for Author
│   │   └── book.go          # HTTP handlers for Book
│   ├── repository/
│   │   ├── author.go        # Author data access (raw SQL)
│   │   └── book.go          # Book data access (raw SQL)
│   └── service/
│       ├── author.go        # Author business logic
│       └── book.go          # Book business logic
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Architecture

The application follows Clean Architecture with four layers, each depending on interfaces defined in `internal/domain/`:

```
┌──────────────────────────────────────────────────┐
│                   Handler                         │
│            (HTTP — net/http)                      │
│         internal/handler/                         │
├──────────────────────────────────────────────────┤
│                   Service                         │
│         (Business Logic)                          │
│         internal/service/                         │
├──────────────────────────────────────────────────┤
│                 Repository                        │
│         (Data Access — raw SQL)                   │
│         internal/repository/                      │
├──────────────────────────────────────────────────┤
│                   MySQL                           │
└──────────────────────────────────────────────────┘
```

- **Handler** — parses HTTP requests, calls service, writes JSON responses
- **Service** — business logic (currently pass-through to repository)
- **Repository** — raw SQL queries via `database/sql`
- **Domain** — shared models and interface contracts

## API Reference

### Authors

| Method | URL             | Description            |
| ------ | --------------- | ---------------------- |
| POST   | `/authors`      | Create a new author    |
| GET    | `/authors`      | Get all authors        |
| GET    | `/authors/{id}` | Get an author by ID    |
| PUT    | `/authors/{id}` | Update an author by ID |
| DELETE | `/authors/{id}` | Delete an author by ID |

#### `POST /authors`

```json
// Request
{
  "name": "John Doe",
  "bio": "An author."
}

// Response — 201 Created
```

#### `GET /authors`

```json
// Response — 200 OK
[
  {
    "id": 1,
    "name": "John Doe",
    "bio": "An author.",
    "created_at": "2025-06-01T12:00:00Z",
    "updated_at": "2025-06-01T12:00:00Z"
  }
]
```

#### `GET /authors/{id}`

```json
// Response — 200 OK
{
  "id": 1,
  "name": "John Doe",
  "bio": "An author.",
  "created_at": "2025-06-01T12:00:00Z",
  "updated_at": "2025-06-01T12:00:00Z"
}
```

#### `PUT /authors/{id}`

```json
// Request
{
  "name": "Jane Doe",
  "bio": "Updated bio."
}

// Response — 200 OK
```

#### `DELETE /authors/{id}`

```
// Response — 200 OK
```

### Books

| Method | URL           | Description         |
| ------ | ------------- | ------------------- |
| POST   | `/books`      | Create a new book   |
| GET    | `/books`      | Get all books       |
| GET    | `/books/{id}` | Get a book by ID    |
| PUT    | `/books/{id}` | Update a book by ID |
| DELETE | `/books/{id}` | Delete a book by ID |

#### `POST /books`

```json
// Request
{
  "title": "The Go Programming Language",
  "author_id": 1,
  "genre": "Programming",
  "synopsis": "A comprehensive guide to Go."
}

// Response — 200 OK
```

#### `GET /books`

```json
// Response — 200 OK
[
  {
    "id": 1,
    "title": "The Go Programming Language",
    "author_id": 1,
    "genre": "Programming",
    "synopsis": "A comprehensive guide to Go.",
    "created_at": "2025-06-01T12:00:00Z",
    "updated_at": "2025-06-01T12:00:00Z"
  }
]
```

#### `GET /books/{id}`

```json
// Response — 200 OK
{
  "id": 1,
  "title": "The Go Programming Language",
  "author_id": 1,
  "genre": "Programming",
  "synopsis": "A comprehensive guide to Go.",
  "created_at": "2025-06-01T12:00:00Z",
  "updated_at": "2025-06-01T12:00:00Z"
}
```

#### `PUT /books/{id}`

```json
// Request
{
  "title": "Updated Title",
  "genre": "Technology",
  "synopsis": "Updated synopsis."
}

// Response — 200 OK
```

#### `DELETE /books/{id}`

```
// Response — 200 OK
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feat/my-feature`
3. Commit your changes: `git commit -m "feat: add my feature"`
4. Push the branch: `git push origin feat/my-feature`
5. Open a pull request

## License

[MIT](LICENSE)
