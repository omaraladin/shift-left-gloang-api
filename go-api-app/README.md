# Go API Application

This is a basic API application built with Go (Golang). The application provides endpoints for health checks and user management.

## Project Structure

```
go-api-app
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── server
│   │   └── server.go        # Server implementation and route definitions
│   └── handlers
│       ├── health.go        # Health check endpoint
│       └── users.go         # User management endpoints
├── pkg
│   └── config
│       └── config.go        # Configuration management
├── api
│   └── openapi.yaml         # OpenAPI specification for the API
├── scripts
│   └── migrate.sh           # Database migration script
├── go.mod                   # Module definition and dependencies
├── .gitignore               # Git ignore file
└── README.md                # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd go-api-app
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run cmd/server/main.go
   ```

4. **Access the API:**
   - Health Check: `GET /health`
   - Users: `GET /users`

## Usage Examples

- To check the health of the API, send a GET request to `/health`.
- To retrieve a list of users, send a GET request to `/users`.

## License

This project is licensed under the MIT License.