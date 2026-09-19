# Sezzle Calculator

Full-stack calculator: Go REST backend + React/Vite frontend. Operations: add, subtract, multiply, divide, power, square root, percentage.

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
curl -X POST http://localhost:8080/api/calculate/divide -d '{"a":10,"b":0}'
curl -X POST http://localhost:8080/api/calculate/sqrt -d '{"a":9}'
```

Available operations: `add` `subtract` `multiply` `divide` `power` `sqrt` `percentage`

## Design decisions

- Backend uses standard `net/http` (no framework), Go 1.22+ routing.
- All arithmetic logic lives in the backend; the frontend only consumes the API.
- API docs generated with Swagger (`swaggo`) from code comments.
- Standardized JSON error envelope (`{"error": {...}}`) across the API.
- Frontend uses CSS Modules (no external styling libraries).
