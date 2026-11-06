# Square POS Integration — Go

A **QLUB knowledge-transfer exercise** implementing a Go API around Square sandbox orders and payments. Gorilla Mux exposes local endpoints, JWT middleware identifies users, and GORM/PostgreSQL stores users and order records.

## Run locally

Prerequisites: Go 1.23.5 or a compatible newer toolchain, PostgreSQL, and a Square sandbox application with an access token and location ID.

```bash
git clone https://github.com/paranietharan/square-pos.git
cd square-pos
go mod download
```

Create a local PostgreSQL database named `square_pos`. Configure `.env` in the repository root:

```dotenv
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_local_user
DB_PASSWORD=your_local_password
DB_NAME=square_pos
DB_SSLMODE=disable
JWT_SECRET=replace-with-a-local-random-secret
JWT_EXPIRATION_IN_SECONDS=3600
ACCESS_TOKEN=your_square_sandbox_access_token
LOCATION_ID=your_square_sandbox_location_id
CREATE_ORDER_API_URL=https://connect.squareupsandbox.com/v2/orders
```

The repository already tracks a `.env` file. Use your own local values and do not commit credentials. The database startup code requires the file to exist and runs GORM auto-migration for users and orders.

```bash
go run ./cmd
```

The server listens on `http://localhost:8080`. That port is hardcoded in the router; setting `PORT` alone does not change it. Although the generic config contains MySQL-like defaults, the actual connection in [pkg/database/db.go](pkg/database/db.go) uses PostgreSQL.

## Endpoints

All paths have the prefix `/api/v1`.

| Method | Path | Purpose |
| --- | --- | --- |
| POST | `/register` | Register a local user |
| POST | `/login` | Obtain a JWT |
| GET | `/admin` | Authenticated greeting |
| POST | `/create-order` | Create a Square sandbox order |
| GET | `/order/{id}` | Retrieve an order |
| GET | `/order/table/{id}` | Retrieve orders associated with a table |
| POST | `/submit-payment` | Submit a Square sandbox payment |

The order/payment routes require an Authorization bearer token from the local login flow. The `/admin` name does not itself establish an administrator-only authorization policy. Request shapes are defined in [pkg/dto](pkg/dto) and [pkg/types](pkg/types).

## Source guide

- [pkg/router/api.go](pkg/router/api.go): Route groups and server startup.
- [pkg/service/user/routes.go](pkg/service/user/routes.go): Registration/login handlers.
- [pkg/service/pos/square_impl.go](pkg/service/pos/square_impl.go): Square requests and sandbox URLs.
- [pkg/service/auth/jwt.go](pkg/service/auth/jwt.go): Token creation and middleware.
- [docs](docs): Supporting project material.

The integration sends a hardcoded Square API version header (`2025-01-23`). Retrieval and payment URLs target the sandbox. Use sandbox credentials and test payment data when exploring the implementation.

## Development checks and limitations

```bash
go build ./...
go test ./...
```

No `*_test.go` files are currently included. The JWT middleware logs extracted tokens, and external API error handling is simplified. This repository preserves an integration exercise; payment reliability and authentication behavior have not been established by an automated test suite.
