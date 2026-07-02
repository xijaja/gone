# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Gone is a Go + SvelteKit full-stack template that compiles into a single deployable binary: the SvelteKit frontend is built to static files and embedded into the Go binary via `go:embed` (`svelte/frontend.go`). Backend is Fiber + GORM + Postgres + Redis. Code comments, log messages, and commit messages are written in Chinese.

There are two branches: `main` (Postgres) and `with-sqlite` (embeds SQLite as well).

## Commands

### Frontend (run inside `svelte/`, uses pnpm)

```sh
pnpm install
pnpm run build        # REQUIRED before building/running the Go backend (go:embed needs svelte/build/)
pnpm run dev          # Vite dev server
pnpm run check        # svelte-check type checking
pnpm run lint         # prettier --check + eslint
pnpm run format       # prettier --write
```

### Backend (run from repo root, so godotenv finds the env file)

```sh
make pg               # start local Postgres in Docker (credentials match .env.dev)
make rs               # start local Redis in Docker
go run ./cmd/app      # dev mode: loads .env.dev, serves on :3030
go build -o app ./cmd/app
go test ./...         # no tests exist yet; config detects test mode and loads .env.dev from project root
```

Runtime flags (parsed in `internal/config/start.go`): `-s` production mode (loads `.env` instead of `.env.dev`, enables CSRF + log-to-DB middleware, serves embedded frontend), `-p` Fiber prefork, `-t` start cron tasks (`internal/auto`).

### Deployment

- `make deploy` — build frontend → Docker image → ship to remote host over SSH → run (see `makefile` for individual steps)
- `make fly` / `.github/workflows/fly-deploy.yml` — fly.io deploy; CI builds the frontend first because the embed requires it, and writes `.env` from GitHub secrets
- `make help` lists all targets

## Critical Setup Facts

- **`svelte/build/` must exist or the Go build fails** — `go:embed build/*` in `svelte/frontend.go` has nothing to embed otherwise. Always build the frontend first.
- **Postgres and Redis must be reachable to even start (or test) the backend.** `database/model/auto_migrate.go` has an `init()` that connects to Postgres, runs `AutoMigrate`, and seeds an `admin`/`123456` user at import time. `database/access/redis.go` similarly sets up the Redis pool. Anything importing `gone/database/model` pulls this in.
- Config is loaded and validated (all env vars required) in `internal/config`'s `init()`; missing vars are fatal. Dev config template is `.env.dev`; production copies it to `.env`.

## Architecture

### Startup and middleware order (`cmd/app/main.go`)

Order matters: CORS → request logging → (prod only: CSRF, log-to-DB) → `apis.Router` → static file middleware **last**, because the SPA filesystem middleware with `NotFoundFile: index.html` would otherwise shadow the API routes.

In dev mode static files are served from disk (`FRONTEND_STATIC_PATH`, default `./svelte/build`); in prod (`-s`) from the embedded FS. Port is `:3030`, with graceful shutdown handled by `config.GracefullyShuttingDown`.

### Request flow

- All API routes live under `/api`, registered in `apis/router.go`. The `/api/v1` group is protected by `middleware.JwtAuth` (Bearer token, HS256, `golang-jwt/v5`; claims defined in `apis/middleware/jwk.go`).
- Handlers (`apis/handler/`) are methods on empty struct pointers (`var u *handler.User; apiV1.Post("/user/login", u.Login)`).
- All responses use `internal/result`: `result.Success()/Error()/NoPermission()/SystemError()` returning `{code, msg, data}` with app-level codes 2000/4000/4001/5000. Attach payloads with `.WithData(fiber.Map{...})`. `internal/result/request.go` is an unrelated outbound-HTTP-client helper.

### Database layer

- `database/access/` — connection setup (GORM/Postgres with prod-dependent logging and prepared statements; redigo Redis pool).
- `database/model/` — active-record-style models: each struct holds its own CRUD methods using the package-level `db` handle. New models must be added to the `AutoMigrate` call in `auto_migrate.go`. A custom `sql` struct tag creates Postgres enum types before migration (see `sqlTagExecutor`). `local_time.go` provides the `LocalTime` type and `SomeTimesAt` embedded timestamps used by models.
- `database/cache/` — `NewRedis(key)` returns an `Rds` wrapper with typed get/set helpers around the pool.

### Frontend (`svelte/`)

SvelteKit with Svelte 5 runes, `adapter-static` with `fallback: index.html` (pure SPA — no SSR), TailwindCSS, and shadcn-svelte (bits-ui) components configured via `components.json`. `src/lib/dofetch.ts` centralizes API URLs and the fetch wrapper; its `baseUrl` switches between `/api` (production, same origin) and `http://127.0.0.1:3030/api` (dev, hence the CORS middleware).

### Other packages

- `internal/auto` — cron scheduled tasks (robfig/cron), enabled by `-t`.
- `pkg/utils` — logger init (lumberjack rotation to `logs/`), MD5 password hashing, validators, misc helpers.
- `pkg/llms` — DeepSeek API client.
- `cmd/cli` — placeholder second entry point.
