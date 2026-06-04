# pets

Pet management REST API. Gin + GORM + PostgreSQL, clean layered architecture.

## Layout

```
cmd/api            # entrypoint
internal/
  config           # env config
  database         # GORM connection
  domain           # entities + repository interfaces
  repository       # GORM repository implementations
  service          # business logic
  handler          # HTTP handlers (Gin)
  server           # router wiring
```

## Setup

```bash
cp .env.example .env   # edit credentials
go mod tidy
make run               # starts on :8080
```

## Endpoints

| Method | Path              | Description   |
|--------|-------------------|---------------|
| GET    | /health           | health check  |
| POST   | /api/v1/pets      | create pet    |
| GET    | /api/v1/pets      | list pets     |
| GET    | /api/v1/pets/:id  | get pet       |
| PUT    | /api/v1/pets/:id  | update pet    |
| DELETE | /api/v1/pets/:id  | delete pet    |
