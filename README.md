# Auth System

![Go](https://img.shields.io/badge/Go-1.25.4-00ADD8?style=flat-square&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14+-336791?style=flat-square&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Latest-2496ED?style=flat-square&logo=docker)
![Redis](https://img.shields.io/badge/Redis-Alpine-DC382D?style=flat-square&logo=redis)
![JWT](https://img.shields.io/badge/JWT-Latest-000000?style=flat-square&logo=jsonwebtokens)
![GORM](https://img.shields.io/badge/GORM-v1.31.1-FF6B6B?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)

A robust Go backend authentication system with user registration, password hashing, JWT token generation, and PostgreSQL database integration. Built with modern Go practices and containerized with Docker.

## Overview

This is a production-ready authentication backend service that provides secure user registration and management capabilities. It features secure password hashing with bcrypt, JWT token generation for session management, and a clean REST API architecture.

## Features

- **User Registration**: Create new user accounts with validation
- **Secure Password Hashing**: Uses bcrypt with configurable cost factor (10)
- **JWT Token Generation**: 24-hour token expiration for secure authentication
- **Email Uniqueness**: Enforces unique email addresses at the database level
- **Address Management**: Support for embedded address information (street, city, state, zip code)
- **Avatar Support**: Optional user avatar URLs with validation
- **PostgreSQL Integration**: Persistent data storage with automatic schema migration via GORM
- **Redis Support**: Ready for caching and session management
- **Docker Deployment**: Full Docker and Docker Compose support for easy deployment
- **Environment-based Configuration**: Configurable via environment variables

## Tech Stack

- **Language**: Go 1.25.4
- **Database**: PostgreSQL
- **ORM**: GORM v1.31.1
- **Password Hashing**: bcrypt (golang.org/x/crypto)
- **JWT**: github.com/golang-jwt/jwt/v5
- **Caching**: Redis (via go-redis/v9)
- **Containerization**: Docker & Docker Compose

## Project Structure

```
auth-system/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── database/
│   │   └── db.go               # PostgreSQL connection and schema migration
│   ├── handlers/
│   │   └── auth.go             # HTTP handlers for registration and auth
│   ├── middleware/              # (Extensible for auth middleware)
│   ├── models/
│   │   └── user.go             # User and Address data models
│   └── services/                # (Extensible for business logic)
├── pkg/
│   └── utils/
│       ├── jwt.go              # JWT token generation
│       └── password.go         # Password hashing and validation
├── docker-compose.yml          # Multi-container orchestration
├── Dockerfile                  # Multi-stage Go build
├── go.mod                      # Go module dependencies
└── README.md
```

## Prerequisites

- **Go** 1.25.4 or higher
- **Docker** and **Docker Compose** (for containerized setup)
- **PostgreSQL** 14+ (if running without Docker)
- **Redis** (optional, for caching features)

## Installation

### Option 1: Using Docker Compose (Recommended)

Clone the repository and set up your environment:

```bash
git clone <repository>
cd auth-system
cp .env.example .env  # Create environment file
```

Create a `.env` file in the root directory:

```env
PORT=8080
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=your_secure_password
DB_NAME=authdb
DB_PORT=5432
REDIS_ADDR=redis:6379
REDIS_PORT=6379
JWT_SECRET=your_jwt_secret_key_here
```

Start the services:

```bash
docker-compose up --build
```

The API will be accessible at `http://localhost:8080`

### Option 2: Local Development

Install dependencies:

```bash
go mod download
```

Set up environment variables:

```bash
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=your_password
export DB_NAME=authdb
export DB_PORT=5432
export PORT=8080
export JWT_SECRET=your_jwt_secret
```

Run the application:

```bash
go run ./cmd/api
```

## Configuration

All configuration is managed via environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | 8080 |
| `DB_HOST` | PostgreSQL host | localhost |
| `DB_USER` | Database user | postgres |
| `DB_PASSWORD` | Database password | (required) |
| `DB_NAME` | Database name | authdb |
| `DB_PORT` | Database port | 5432 |
| `JWT_SECRET` | Secret key for JWT signing | (required) |
| `REDIS_ADDR` | Redis address | localhost:6379 |
| `REDIS_PORT` | Redis port | 6379 |

## API Endpoints

### User Registration

**Endpoint**: `POST /api/v1/register`

**Request Body**:
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "password": "securepassword123",
  "avatar_url": "https://example.com/avatar.jpg",
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "zip_code": "10001"
  }
}
```

**Success Response** (201 Created):
```json
{
  "message": "User registerd successfully!"
}
```

**Error Responses**:
- `400 Bad Request`: Invalid JSON format or missing required fields
- `409 Conflict`: Email already exists

## Data Models

### User

```go
type User struct {
    ID        uint          `json:"id"`
    FirstName string        `json:"first_name"`
    LastName  string        `json:"last_name"`
    Email     string        `json:"email"`
    Password  string        `json:"password"` // hashed
    Address   Address       `json:"address"`
    AvatarURL string        `json:"avatar_url"`
    CreatedAt time.Time     `json:"created_at"`
    UpdatedAt time.Time     `json:"updated_at"`
}
```

### Address

```go
type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    State   string `json:"state"`
    ZipCode string `json:"zip_code"`
}
```

## Security Features

- **Password Hashing**: Uses bcrypt with cost factor of 10 for computational complexity
- **Email Uniqueness**: Database constraint prevents duplicate emails
- **JWT Tokens**: 24-hour expiration time
- **Input Validation**: Email format and password length validation
- **Environment-based Secrets**: Sensitive data stored in environment variables
- **Error Handling**: Generic error messages to prevent information leakage

## Development

### Building the Docker Image

```bash
docker build -t auth-system .
```

### Running Tests

```bash
go test ./...
```

### Code Structure

- **Handlers** (`internal/handlers/`): HTTP request processing and validation
- **Models** (`internal/models/`): Data structures and GORM tags
- **Database** (`internal/database/`): Connection management and migrations
- **Utils** (`pkg/utils/`): Shared utilities (JWT, password hashing)

## Database Schema

The application automatically creates and migrates the `users` table on startup:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    street VARCHAR,
    city VARCHAR,
    state VARCHAR,
    zip_code VARCHAR,
    avatar_url VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Troubleshooting

### Database Connection Failed

Ensure all database environment variables are correctly set and PostgreSQL is running:

```bash
# Check environment variables
env | grep DB_
```

### Port Already in Use

Change the `PORT` environment variable or kill the process using the port:

```bash
lsof -i :8080
kill -9 <PID>
```

### Docker Build Issues

Clean up Docker and rebuild:

```bash
docker-compose down -v
docker-compose up --build
```

## Future Enhancements

- [ ] User login endpoint with JWT authentication
- [ ] Email verification flow
- [ ] Password reset functionality
- [ ] Role-based access control (RBAC)
- [ ] OAuth2 integration
- [ ] Rate limiting and DDoS protection
- [ ] API documentation (Swagger/OpenAPI)
- [ ] User profile update endpoint
- [ ] Session management with Redis

## License

MIT License - feel free to use this in your projects.

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests.
