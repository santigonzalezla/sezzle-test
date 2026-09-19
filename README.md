# Sezzle Calculator

Full-stack calculator: Go REST backend + React/Vite frontend. Operations: add, subtract, multiply, divide, power, square root, percentage.

Full technical/API documentation (all endpoints, request/response schemas) is generated with Swagger and available at https://backend-production-d763.up.railway.app/docs/ (or `http://localhost:8080/docs/` when running locally).

## Structure

```
backend/    Go REST API (standard net/http, Swagger at /docs/)
frontend/   React 19 + Vite + TypeScript
```

## Requirements

- Go 1.22+
- Node 18+ and pnpm
- Docker (optional, to run everything together)

## 1. Clone

```bash
git clone https://github.com/santigonzalezla/sezzle-test.git
cd sezzle-test
```

## 2. Run locally

### Backend

```bash
cd backend
cp .env.example .env
go run ./cmd/api
```

- API: `http://localhost:8080`
- Swagger: `http://localhost:8080/docs/`

### Frontend (in another terminal)

```bash
cd frontend
cp .env.example .env
pnpm install
pnpm run dev
```

- App: `http://localhost:5173`

## 3. Run with Docker (backend + frontend together)

```bash
docker compose up --build
```

- Backend: `http://localhost:8080`
- Frontend: `http://localhost:8081`

## 4. Tests

```bash
# Backend
cd backend && go test ./... -cover

# Frontend
cd frontend && pnpm run test:coverage
```

## API examples

```bash
curl -X POST http://localhost:8080/api/calculate/add -d '{"a":2,"b":3}'
# -> {"operation":"add","a":2,"b":3,"result":5}

curl -X POST http://localhost:8080/api/calculate/divide -d '{"a":10,"b":0}'
# -> {"error":{"code":"CALCULATION_ERROR","message":"division by zero","path":"/api/calculate/divide","timestamp":"..."}}

curl -X POST http://localhost:8080/api/calculate/sqrt -d '{"a":9}'
# -> {"operation":"sqrt","a":9,"b":0,"result":3}
```

Available operations: `add` `subtract` `multiply` `divide` `power` `sqrt` `percentage`

## Design decisions

- Backend uses standard `net/http` (no framework), Go 1.22+ routing.
- All arithmetic logic lives in the backend; the frontend only consumes the API.
- API docs generated with Swagger (`swaggo`) from code comments.
- Standardized JSON error envelope (`{"error": {...}}`) across the API.
- Frontend uses CSS Modules (no external styling libraries).

## Assumptions

- `sqrt` only needs operand `a`; `b` is accepted but ignored.
- No operator chaining without pressing `=` (e.g. `5 + 3 + 2 =` is not supported); pressing an operator after a result reuses that result as the next `a`.
- `CORS_ORIGINS` is a comma-separated list of allowed origins.

## Deploying to Railway

Two services from the same GitHub repo, each built from its existing `Dockerfile`, each with its own public domain:

**Backend**
- Root Directory: `backend`
- Public port: `8080`
- Variables: `CORS_ORIGINS=https://<frontend-public-domain>`

**Frontend**
- Root Directory: `frontend`
- Public port: `80` (nginx)
- Build variable: `VITE_API_BASE_URL=https://<backend-public-domain>` — must be set as a **build-time** variable, since Vite bakes it into the JS bundle at build time, not at runtime.

Order: deploy the backend first to get its domain, set it as `VITE_API_BASE_URL` on the frontend, then set the frontend's domain as `CORS_ORIGINS` on the backend. Every `git push` to `main` auto-redeploys both services.

**Gotcha**: `CORS_ORIGINS` must match the frontend's origin exactly, including the scheme — `http://` vs `https://` is a different origin and silently breaks CORS (browser shows `Failed to fetch`; backend logs only show the failed `OPTIONS` preflight, never the real request).

## Live demo

- Frontend: https://frontend-production-72be1.up.railway.app
- Backend: https://backend-production-d763.up.railway.app
- Swagger: https://backend-production-d763.up.railway.app/docs/
