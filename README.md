# medcord-backend

Backend API for Medcord (Caelum) — a multi-tenant hospital management SaaS. Hospitals sign up, set up a workspace, and run patient management, EMR, labs, and asset tracking from one platform.

## Stack

- **Go 1.25** (auto-upgraded toolchain via `go.mod`)
- **Gin** — HTTP framework
- **MongoDB** via the official `go.mongodb.org/mongo-driver`
- **`go-playground/validator`** — request validation (via Gin's `binding` tags)
- **`golang-jwt/jwt/v5`** + **`bcrypt`** — auth
- **`log/slog`** — structured logging
- **`ulule/limiter`** — rate limiting
- **`gin-contrib/secure`** + **`gin-contrib/cors`** — security headers, CORS

Architecture details: [`../docs/backend-guide-go.md`](../docs/backend-guide-go.md).
Product spec: [`../docs/readme.md`](../docs/readme.md).

## Project layout

```
cmd/server/             entry point + graceful shutdown
internal/
  app/                  Gin engine + middleware wiring
  configs/              env loader
  controllers/          HTTP handlers
  deps/                 dependency wiring (built once at startup)
  middlewares/          recovery, request logger, rate limit, (auth — TBD)
  routes/               route registration
  shared/
    constants/          message keys + en/es/fr translations
    types/              ServiceResult[T], Page[T]
  utils/
    database/           Mongo connect + ping
    logger/             slog setup
    response/           response envelope helpers
```

Future modules add `services/`, `repositories/`, `models/`, `requests/`, and `migrations/` as needed.

## Getting started

### Prerequisites

- Go 1.23+ (toolchain will auto-upgrade if a dep needs newer)
- MongoDB running locally on `:27017` (or set `MONGODB_URI`)

### Setup

```bash
cp .env.example .env       # tweak JWT_SECRET, MONGODB_URI, etc.
go mod download
make run                   # or: go run ./cmd/server
```

For live reload during development:

```bash
go install github.com/air-verse/air@latest
make dev
```

## Make targets

| Target | What it does |
| --- | --- |
| `make run` | Run the server once |
| `make dev` | Run with `air` live reload |
| `make build` | Build to `bin/server` |
| `make test` | `go test ./...` |
| `make lint` | `golangci-lint run` |
| `make tidy` | `go mod tidy` |

## Environment variables

See `.env.example` for the full list. The server **hard-fails** on a missing `JWT_SECRET` — production safety, not a warning.

| Var | Default | Notes |
| --- | --- | --- |
| `APP_ENV` | `development` | Set to `production` to enable Gin release mode + JSON logs |
| `PORT` | `4000` | |
| `MONGODB_URI` | `mongodb://localhost:27017` | |
| `MONGODB_DB` | `medcord` | |
| `JWT_SECRET` | — | **Required.** No fallback. |
| `JWT_EXPIRES_IN` | `168h` | Go duration format |
| `RATE_LIMIT_WINDOW` | `15m` | |
| `RATE_LIMIT_MAX` | `100` | Per window, per IP |
| `CORS_ORIGINS` | `http://localhost:3000` | Comma-separated |

## API

All endpoints are prefixed with `/api`. Responses follow a single envelope:

```jsonc
// Success
{ "success": true, "data": {...}, "message": "..." }

// Error
{ "success": false, "error": "...", "details": {...} }
```

Localization: pass `?lang=es` or `Accept-Language: fr` to translate the `message`/`error` field. Supported: `en`, `es`, `fr`.

### Endpoints (current)

| Method | Path | Description |
| --- | --- | --- |
| GET | `/api/health` | Health check — returns service status, DB ping, uptime |

More routes land as feature modules ship.

## Conventions

- **Controllers are thin** — bind, validate, call service, format response.
- **Services never panic** — return `ServiceResult[T]` (carries the `error` and a `MessageKey`).
- **Repositories own Mongo** — services never import the Mongo driver.
- **`context.Context` everywhere** — first param on any blocking call; honor cancellation.
- **Constructor injection** — wire deps in `cmd/server/main.go` → `internal/deps`, no globals.

Full conventions, gotchas, and patterns are in [`../docs/backend-guide-go.md`](../docs/backend-guide-go.md). Read the **"Gin gotchas"** section near the top before writing handlers.

## License

Private. © Medcord.
