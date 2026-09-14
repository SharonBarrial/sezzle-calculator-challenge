# Sezzle Calculator Challenge

Full-stack calculator: React + TypeScript frontend, Go backend REST API.

## Stack

- **Backend:** Go 1.22 (standard library only — no external dependencies)
- **Frontend:** React 19 + TypeScript, built with Vite
- **Tests:** Go's built-in `testing` package (backend), Vitest + React Testing Library (frontend)
- **Deployment:** Docker + Docker Compose (optional)

## Project structure
sezzle-calculator/
├── backend/
│ ├── calculator/ # pure arithmetic logic (unit tested, no HTTP)
│ ├── handlers/ # HTTP layer: JSON parsing, validation, routing targets
│ ├── main.go # router + CORS middleware + server bootstrap
│ └── Dockerfile
├── frontend/
│ ├── src/
│ │ ├── api/ # calculatorApi.ts — the only place that talks to the backend
│ │ ├── components/ # Calculator.tsx + Calculator.css
│ │ └── App.tsx
│ └── Dockerfile
├── docker-compose.yml
└── PROMPTS.md # AI prompts used while building this

## Running locally (without Docker)

### Backend

```bash
cd backend
go run .
# server listening on :8080
```

### Frontend

```bash
cd frontend
npm install
npm run dev
# app available at http://localhost:5173
```

By default the frontend calls `http://localhost:8080`. To change this, create a `frontend/.env`:

```bash
VITE_API_BASE_URL=http://localhost:8080
```

## Running with Docker Compose

```bash
docker compose up --build
```

- Frontend: http://localhost:3000
- Backend: http://localhost:8080

To run in the background: `docker compose up --build -d`, then `docker compose logs -f` to follow logs, `docker compose down` to stop.

## Running tests

### Backend

```bash
cd backend
go test ./... -v -cover
```

Coverage: 100% on `calculator/` (pure logic), ~84% on `handlers/` (HTTP layer — the uncovered lines are internal error-encoding edge cases that are impractical to trigger in a test).

Generate an HTML coverage report:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

### Frontend

```bash
cd frontend
npm test                 # run once
npm run test:coverage    # with coverage report
```

11 tests covering the API client (`calculatorApi.test.ts`) and the calculator component (`Calculator.test.tsx`), including validation, successful calls, and backend error handling.

## API reference

Base URL: `http://localhost:8080`

All operation endpoints accept `POST` with a JSON body and return `{ "result": number }` on success, or `{ "error": string }` on failure.

| Endpoint | Body | Notes |
|---|---|---|
| `POST /api/add` | `{ "a": number, "b": number }` | |
| `POST /api/subtract` | `{ "a": number, "b": number }` | |
| `POST /api/multiply` | `{ "a": number, "b": number }` | |
| `POST /api/divide` | `{ "a": number, "b": number }` | 422 if `b` is 0 |
| `POST /api/power` | `{ "a": number, "b": number }` | `a` raised to `b` |
| `POST /api/sqrt` | `{ "a": number }` | 422 if `a` is negative |
| `POST /api/percentage` | `{ "a": number, "b": number }` | What % `a` is of `b`. 422 if `b` is 0 |
| `GET /api/health` | — | Used by Docker healthcheck |

### Example

```bash
curl -X POST http://localhost:8080/api/add \
  -H "Content-Type: application/json" \
  -d '{"a": 2, "b": 3}'
# {"result":5}

curl -X POST http://localhost:8080/api/divide \
  -H "Content-Type: application/json" \
  -d '{"a": 5, "b": 0}'
# HTTP 422 — {"error":"division by zero"}
```

## Design decisions & assumptions

- **Pure logic separated from HTTP (`calculator/` vs `handlers/`).** Makes the arithmetic 100% unit-testable without spinning up a server, and keeps the transport layer swappable.
- **One REST endpoint per operation** rather than a single generic `/calculate` endpoint with the operation name in the body — more RESTful and self-documenting.
- **400 vs 422 status codes.** 400 means the request body itself is malformed (bad JSON, wrong types, unknown fields). 422 means the JSON is well-formed but the *values* are invalid for the operation (divide by zero, negative sqrt). This distinction is intentional and semantically correct HTTP.
- **`DisallowUnknownFields()` on JSON decoding.** Rejects requests with unexpected extra fields instead of silently ignoring them — catches API misuse early.
- **No external Go dependencies.** The standard library's `net/http` (with Go 1.22's method-based routing) was sufficient for this scope, avoiding unnecessary third-party risk/maintenance.
- **CORS is wide open (`Access-Control-Allow-Origin: *`).** Acceptable for this take-home; in a real production deployment this would be locked to the frontend's actual origin.
- **`percentage(a, b)` means "what % is `a` of `b`"** (i.e. `(a/b)*100`), not "increase `a` by `b`%" — documented here since the assignment doesn't specify.
- **Frontend validates input client-side before calling the API** (non-numeric input never reaches the network) but still handles backend-returned errors (e.g. division by zero) distinctly, since some validation can only happen server-side.
- **`VITE_API_BASE_URL` is baked in at frontend build time**, not read at container runtime — this is a Vite/static-site constraint, documented in the Docker section below.

## Docker notes

- Both services use multi-stage builds — the final images contain only the compiled binary (backend) or static files + Nginx (frontend), not the full build toolchain.
- The backend's `/api/health` endpoint backs a Docker healthcheck; the frontend container waits for the backend to be healthy before starting.
- Because Vite bakes environment variables into the JS bundle at build time, `VITE_API_BASE_URL` is passed as a Docker build `ARG`, not a runtime `ENV` — it must point to a URL reachable from the *browser* (e.g. `http://localhost:8080`), not the internal Docker network service name.

## AI tooling disclosure

This project was built with AI assistance (Claude). See [`PROMPTS.md`](./PROMPTS.md) for the prompts used.