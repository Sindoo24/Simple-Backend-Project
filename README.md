# User Management API

A robust RESTful API built with Go, featuring JWT authentication, role-based access control, and comprehensive user management capabilities.

## 🚀 Features

- **Authentication & Authorization**
  - User signup with email validation
  - Secure login with JWT tokens
  - Password strength validation
  - Bcrypt password hashing
  - HTTP-only secure cookies
  - Role-based access control (user/admin)

- **User Management**
  - CRUD operations for users
  - Dynamic age calculation from date of birth
  - Pagination support
  - Input validation

- **Developer Experience**
  - Comprehensive unit tests (14 tests, 100% passing)
  - Standardized error handling with request IDs
  - Structured logging with Zap
  - Docker containerization
  - Database migrations
  - SQLC for type-safe SQL queries

## 📋 Prerequisites

- Docker & Docker Compose
- Go 1.24+ (for local development)
- PostgreSQL 15 (handled by Docker)

## 🛠️ Tech Stack

- **Framework**: [Fiber](https://gofiber.io/) v2.52.10
- **Database**: PostgreSQL 15
- **Authentication**: JWT (golang-jwt/jwt)
- **Password Hashing**: bcrypt
- **Logging**: Zap
- **Validation**: go-playground/validator
- **SQL**: SQLC for type-safe queries
- **Testing**: Go testing package

## 📦 Installation

### Using Docker (Recommended)

1. Clone the repository:
```bash
git clone https://github.com/Sindoo24/Simple-Backend-Project.git
cd Simple-Backend-Project
```

2. Create environment file:
```bash
cp .env.example .env
```

3. Update `.env` with your configuration:
```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/userdb?sslmode=disable
SERVER_PORT=8080
LOG_LEVEL=info
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRY_HOURS=24
COOKIE_SECURE=true
```

4. Start the application:
```bash
docker-compose up -d
```

5. Verify the application is running:
```bash
docker-compose logs -f api
```

The API will be available at `http://localhost:8080`

### Local Development

1. Install dependencies:
```bash
go mod download
```

2. Install SQLC:
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

3. Generate SQL code:
```bash
sqlc generate
```

4. Run the application:
```bash
go run cmd/server/main.go
```

## 🧪 Testing

Run all tests:
```bash
go test -v ./...
```

Run specific test package:
```bash
go test -v ./internal/handler
go test -v ./internal/service
```

Run tests with coverage:
```bash
go test -v -cover ./...
```

## 📚 API Documentation

### Base URL
```
http://localhost:8080
```

### Authentication Endpoints

#### Signup
Create a new user account.

**Endpoint**: `POST /auth/signup`

**Request Body**:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "SecurePass123!",
  "dob": "1990-01-15"
}
```

**Response** (201 Created):
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "created_at": "2026-01-07T10:00:00Z"
}
```

**Password Requirements**:
- Minimum 8 characters
- At least one uppercase letter
- At least one lowercase letter
- At least one digit
- At least one special character (!@#$%^&*()_+-=[]{}|;:,.<>?)

#### Login
Authenticate and receive a JWT token.

**Endpoint**: `POST /auth/login`

**Request Body**:
```json
{
  "email": "john@example.com",
  "password": "SecurePass123!"
}
```

**Response** (200 OK):
```json
{
  "message": "Login successful",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user"
  }
}
```

**Note**: JWT token is set as an HTTP-only secure cookie.

### User Endpoints (Protected)

All user endpoints require authentication via JWT token in the `Authorization` header:
```
Authorization: Bearer <your-jwt-token>
```

#### Get All Users
**Endpoint**: `GET /users`

**Query Parameters**:
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10, max: 100)

**Response** (200 OK):
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "dob": "1990-01-15",
    "age": 36
  }
]
```

#### Get User by ID
**Endpoint**: `GET /users/:id`

**Response** (200 OK):
```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "1990-01-15",
  "age": 36
}
```

#### Create User
**Endpoint**: `POST /users`

**Request Body**:
```json
{
  "name": "Jane Smith",
  "dob": "1995-05-20"
}
```

#### Update User
**Endpoint**: `PUT /users/:id`

**Request Body**:
```json
{
  "name": "Jane Doe",
  "dob": "1995-05-20"
}
```

#### Delete User
**Endpoint**: `DELETE /users/:id`

**Response** (204 No Content)

### Admin Endpoints (Admin Only)

#### Get All Users (Admin)
**Endpoint**: `GET /admin/users`

Requires admin role. Returns all users with additional details.

### Error Responses

All errors follow a standardized format:

```json
{
  "error": {
    "message": "Error description",
    "code": "ERROR_CODE",
    "request_id": "unique-request-id"
  }
}
```

**Error Codes**:
- `VALIDATION_FAILED`: Input validation error
- `INVALID_CREDENTIALS`: Invalid email or password
- `INVALID_TOKEN`: Invalid or expired JWT token
- `UNAUTHORIZED`: Missing authentication
- `FORBIDDEN`: Insufficient permissions
- `NOT_FOUND`: Resource not found
- `ALREADY_EXISTS`: Duplicate resource
- `INTERNAL_ERROR`: Server error

## 🏗️ Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── config/
│   └── config.go                # Configuration management
├── db/
│   ├── migrations/              # Database migrations
│   └── sqlc/                    # SQLC queries and generated code
├── internal/
│   ├── handler/                 # HTTP handlers
│   │   ├── auth_handler.go
│   │   ├── auth_handler_test.go
│   │   ├── user_handler.go
│   │   └── admin_handler.go
│   ├── middleware/              # HTTP middleware
│   │   ├── auth.go
│   │   └── logger.go
│   ├── models/                  # Data models
│   │   ├── auth.go
│   │   ├── errors.go
│   │   ├── response.go
│   │   └── user.go
│   ├── repository/              # Data access layer
│   │   └── user_repository.go
│   ├── routes/                  # Route definitions
│   │   └── routes.go
│   └── service/                 # Business logic
│       ├── auth_service.go
│       └── auth_service_test.go
├── docker-compose.yml           # Docker composition
├── Dockerfile                   # Container definition
├── go.mod                       # Go dependencies
└── README.md                    # This file
```

## 🔒 Security Features

- **Password Security**: Bcrypt hashing with cost factor 12
- **JWT Tokens**: Secure token-based authentication
- **HTTP-only Cookies**: Prevents XSS attacks
- **Secure Cookies**: HTTPS-only in production
- **SameSite Strict**: CSRF protection
- **Input Validation**: Comprehensive request validation
- **SQL Injection Prevention**: Parameterized queries via SQLC

## 🐳 Docker Commands

```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop services
docker-compose down

# Rebuild and start
docker-compose up -d --build

# View container status
docker-compose ps
```

## 📊 Database Schema

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License.

## 👥 Authors

- **Sindoo24** - [GitHub](https://github.com/Sindoo24)

## 🙏 Acknowledgments

- Fiber framework for the excellent web framework
- SQLC for type-safe SQL
- The Go community for amazing tools and libraries
