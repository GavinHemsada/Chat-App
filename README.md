# Chat App

<div align="center">

**A production-ready backend chat application built with Go**

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14+-316192?style=flat&logo=postgresql)](https://www.postgresql.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

[Features](#-features) • [Tech Stack](#-technology-stack) • [Getting Started](#-getting-started) • [API Documentation](#-api-documentation) • [Architecture](#-architecture)

</div>

---

## 📌 Overview

Chat App is a backend-focused real-time messaging application designed to demonstrate production-grade software engineering practices. Built with Go and PostgreSQL, this project showcases clean architecture, secure authentication, RESTful API design, and database management patterns suitable for modern web applications.

**Purpose**: This project serves as both a learning resource and a portfolio piece, demonstrating the ability to build scalable, maintainable backend systems from the ground up.

### Use Cases

- Real-time team collaboration
- Multi-room chat system
- Secure messaging platform
- Foundation for building chat features into larger applications

---

## ✨ Features

### Core Functionality
- 🔐 **Secure Authentication** - JWT-based auth with bcrypt password hashing
- 👥 **User Management** - Registration, login, and profile management
- 💬 **Room-based Chat** - Create and join multiple chat rooms
- 📨 **Message Persistence** - Reliable message storage and retrieval
- 🔒 **Protected Routes** - Middleware-based authorization
- 🗄️ **Database Migrations** - Version-controlled schema management

### Technical Highlights
- Clean architecture with clear separation of concerns
- Environment-based configuration
- Comprehensive error handling
- SQL injection protection via parameterized queries
- Scalable project structure
- Production-ready security practices

### Upcoming Features
- 🔌 Real-time WebSocket messaging
- 👑 Admin roles and permissions
- 📎 File attachments
- 🔍 Message search
- 📊 User analytics

---

## 🧰 Technology Stack

### Backend
| Technology | Purpose |
|------------|---------|
| **Go 1.20+** | Primary programming language |
| **Gorilla Mux** | HTTP routing and middleware |
| **PostgreSQL 14+** | Relational database |
| **JWT** | Stateless authentication |
| **bcrypt** | Password hashing |

### Tools & Libraries
- `golang-migrate` - Database migration management
- `lib/pq` - PostgreSQL driver
- `godotenv` - Environment variable management
- `jwt-go` - JWT token handling

### Development
- Git & GitHub for version control
- Docker (optional) for containerization
- Postman/cURL for API testing

---

## 🏗️ Architecture

The application follows **Clean Architecture** principles, ensuring testability, maintainability, and clear separation of concerns.

```
┌─────────────────────────────────────────┐
│           Client (HTTP/WS)              │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│          Handlers (HTTP Layer)          │
│  • Request parsing                      │
│  • Response formatting                  │
│  • Input validation                     │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│       Middleware (Auth, Logging)        │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│      Services (Business Logic)          │
│  • Authentication logic                 │
│  • Authorization checks                 │
│  • Business rules                       │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│     Repositories (Data Access)          │
│  • SQL queries                          │
│  • Database transactions                │
│  • Data mapping                         │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│         PostgreSQL Database             │
└─────────────────────────────────────────┘
```

### Layer Responsibilities

**Handlers**: Accept HTTP requests, validate input, call services, return responses  
**Middleware**: Cross-cutting concerns (auth, logging, CORS)  
**Services**: Implement business logic and orchestrate repositories  
**Repositories**: Handle all database operations  
**Models**: Define data structures and database schemas

---

## 📁 Project Structure

```
Chat-App/
│
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
│
├── internal/
│   ├── config/
│   │   └── config.go               # Environment configuration
│   │
│   ├── database/
│   │   ├── connection.go           # Database connection setup
│   │   └── migrations.go           # Migration runner
│   │
│   ├── handlers/
│   │   ├── auth_handler.go         # Authentication endpoints
│   │   ├── user_handler.go         # User management endpoints
│   │   ├── room_handler.go         # Chat room endpoints
│   │   └── message_handler.go      # Message endpoints
│   │
│   ├── middleware/
│   │   ├── auth.go                 # JWT authentication middleware
│   │   ├── logging.go              # Request logging
│   │   └── cors.go                 # CORS configuration
│   │
│   ├── models/
│   │   ├── user.go                 # User model
│   │   ├── room.go                 # Room model
│   │   └── message.go              # Message model
│   │
│   ├── repository/
│   │   ├── user_repository.go      # User data access
│   │   ├── room_repository.go      # Room data access
│   │   └── message_repository.go   # Message data access
│   │
│   ├── services/
│   │   ├── auth_service.go         # Authentication business logic
│   │   ├── user_service.go         # User business logic
│   │   ├── room_service.go         # Room business logic
│   │   └── message_service.go      # Message business logic
│   │
│   └── websocket/
│       ├── hub.go                  # WebSocket connection hub
│       └── client.go               # WebSocket client handling
│
├── migrations/
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   ├── 000002_create_rooms_table.up.sql
│   ├── 000002_create_rooms_table.down.sql
│   ├── 000003_create_messages_table.up.sql
│   └── 000003_create_messages_table.down.sql
│
├── .env.example                    # Example environment variables
├── .gitignore
├── go.mod                          # Go module definition
├── go.sum                          # Dependency checksums
├── LICENSE
└── README.md                       # This file
```

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed:

- **Go** `1.20` or higher → [Download](https://golang.org/dl/)
- **PostgreSQL** `14` or higher → [Download](https://www.postgresql.org/download/)
- **golang-migrate** → [Installation Guide](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- **Git** → [Download](https://git-scm.com/downloads)

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/yourusername/chat-app.git
cd chat-app
```

2. **Install dependencies**
```bash
go mod download
go mod tidy
```

3. **Set up environment variables**
```bash
cp .env.example .env
```

Edit `.env` with your configuration:
```env
# Server Configuration
APP_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=chat_app
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_change_this_in_production
JWT_EXPIRES_IN=24h

# Environment
ENVIRONMENT=development
```

4. **Create the database**
```bash
# Connect to PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE chat_app;

# Exit psql
\q
```

5. **Run database migrations**
```bash
migrate -path migrations \
  -database "postgres://postgres:your_password@localhost:5432/chat_app?sslmode=disable" \
  up
```

Or if migrations run automatically on startup, just proceed to the next step.

6. **Start the server**
```bash
go run ./cmd/server/main.go
```

The server will start at `http://localhost:8080`

### Health Check

Verify the server is running:
```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "OK",
  "timestamp": "2026-01-28T10:30:00Z"
}
```

---

## 📡 API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication Flow

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant DB
    
    Client->>API: POST /users/register
    API->>DB: Create user (hashed password)
    DB-->>API: User created
    API-->>Client: 201 Created
    
    Client->>API: POST /users/login
    API->>DB: Verify credentials
    DB-->>API: User found
    API->>API: Generate JWT
    API-->>Client: 200 OK + JWT token
    
    Client->>API: GET /users (with JWT)
    API->>API: Validate JWT
    API->>DB: Fetch users
    DB-->>API: Users data
    API-->>Client: 200 OK + Users
```

### Endpoints

#### 🔓 Public Endpoints (No Auth Required)

##### Register User
```http
POST /api/v1/users/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "SecureP@ssw0rd",
  "username": "johndoe"
}
```

**Response** `201 Created`
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "johndoe",
  "created_at": "2026-01-28T10:30:00Z"
}
```

##### Login User
```http
POST /api/v1/users/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "SecureP@ssw0rd"
}
```

**Response** `200 OK`
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "username": "johndoe"
  }
}
```

---

#### 🔐 Protected Endpoints (Auth Required)

**All protected endpoints require the JWT token in the Authorization header:**
```http
Authorization: Bearer <your_jwt_token>
```

##### Get All Users
```http
GET /api/v1/users
Authorization: Bearer <token>
```

**Response** `200 OK`
```json
[
  {
    "id": 1,
    "email": "user@example.com",
    "username": "johndoe",
    "created_at": "2026-01-28T10:30:00Z"
  },
  {
    "id": 2,
    "email": "jane@example.com",
    "username": "janedoe",
    "created_at": "2026-01-28T11:00:00Z"
  }
]
```

##### Get User by ID
```http
GET /api/v1/users/{id}
Authorization: Bearer <token>
```

**Response** `200 OK`
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "johndoe",
  "created_at": "2026-01-28T10:30:00Z"
}
```

##### Create Chat Room
```http
POST /api/v1/rooms
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "General Discussion",
  "description": "A place for general topics"
}
```

**Response** `201 Created`
```json
{
  "id": 1,
  "name": "General Discussion",
  "description": "A place for general topics",
  "created_by": 1,
  "created_at": "2026-01-28T12:00:00Z"
}
```

##### Get All Rooms
```http
GET /api/v1/rooms
Authorization: Bearer <token>
```

**Response** `200 OK`
```json
[
  {
    "id": 1,
    "name": "General Discussion",
    "description": "A place for general topics",
    "created_by": 1,
    "member_count": 5,
    "created_at": "2026-01-28T12:00:00Z"
  }
]
```

##### Join Room
```http
POST /api/v1/rooms/{id}/join
Authorization: Bearer <token>
```

**Response** `200 OK`
```json
{
  "message": "Successfully joined room",
  "room_id": 1
}
```

##### Send Message
```http
POST /api/v1/messages
Authorization: Bearer <token>
Content-Type: application/json

{
  "room_id": 1,
  "content": "Hello, everyone!"
}
```

**Response** `201 Created`
```json
{
  "id": 1,
  "room_id": 1,
  "user_id": 1,
  "username": "johndoe",
  "content": "Hello, everyone!",
  "created_at": "2026-01-28T13:00:00Z"
}
```

##### Get Room Messages
```http
GET /api/v1/rooms/{id}/messages?limit=50&offset=0
Authorization: Bearer <token>
```

**Response** `200 OK`
```json
[
  {
    "id": 1,
    "room_id": 1,
    "user_id": 1,
    "username": "johndoe",
    "content": "Hello, everyone!",
    "created_at": "2026-01-28T13:00:00Z"
  },
  {
    "id": 2,
    "room_id": 1,
    "user_id": 2,
    "username": "janedoe",
    "content": "Hi there!",
    "created_at": "2026-01-28T13:05:00Z"
  }
]
```

---

### Error Responses

All errors follow this format:

```json
{
  "error": "Error message description"
}
```

#### HTTP Status Codes

| Code | Meaning | When Used |
|------|---------|-----------|
| `200` | OK | Successful GET/PUT/DELETE request |
| `201` | Created | Successful POST request |
| `400` | Bad Request | Invalid input data |
| `401` | Unauthorized | Missing or invalid JWT token |
| `403` | Forbidden | Valid token but insufficient permissions |
| `404` | Not Found | Resource doesn't exist |
| `409` | Conflict | Resource already exists (e.g., duplicate email) |
| `500` | Internal Server Error | Server-side error |

#### Example Error Responses

**400 Bad Request**
```json
{
  "error": "Invalid email format"
}
```

**401 Unauthorized**
```json
{
  "error": "Missing or invalid authentication token"
}
```

**404 Not Found**
```json
{
  "error": "User not found"
}
```

**409 Conflict**
```json
{
  "error": "Email already registered"
}
```

---

## 🔌 WebSocket Support (Optional)

The application supports real-time messaging via WebSocket connections.

### Connection

```javascript
const token = "your_jwt_token";
const ws = new WebSocket(`ws://localhost:8080/ws?token=${token}`);

ws.onopen = () => {
  console.log("Connected to chat");
};

ws.onmessage = (event) => {
  const message = JSON.parse(event.data);
  console.log("New message:", message);
};

ws.onerror = (error) => {
  console.error("WebSocket error:", error);
};

ws.onclose = () => {
  console.log("Disconnected from chat");
};
```

### Events

#### Send Message
```json
{
  "type": "message:send",
  "payload": {
    "room_id": 1,
    "content": "Hello from WebSocket!"
  }
}
```

#### Receive Message
```json
{
  "type": "message:new",
  "payload": {
    "id": 123,
    "room_id": 1,
    "user_id": 5,
    "username": "johndoe",
    "content": "Hello from WebSocket!",
    "created_at": "2026-01-28T14:00:00Z"
  }
}
```

#### Join Room
```json
{
  "type": "room:join",
  "payload": {
    "room_id": 1
  }
}
```

---

## 🗄️ Database Schema

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Rooms Table
```sql
CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_by INTEGER REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Messages Table
```sql
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES rooms(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Room Members Table
```sql
CREATE TABLE room_members (
    id SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES rooms(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(room_id, user_id)
);
```

---

## 🔒 Security

### Authentication & Authorization
- **JWT Tokens**: Stateless authentication with configurable expiration
- **Password Hashing**: bcrypt with salt rounds for secure password storage
- **Protected Routes**: Middleware-based authorization on all sensitive endpoints

### Best Practices Implemented
- ✅ Passwords never stored in plain text
- ✅ JWT tokens signed with secret key
- ✅ Environment variables for sensitive data
- ✅ SQL injection prevention via parameterized queries
- ✅ HTTPS recommended for production
- ✅ CORS configuration for API security
- ✅ Input validation on all endpoints

### Production Security Checklist
- [ ] Use strong JWT_SECRET (min 32 characters)
- [ ] Enable HTTPS/TLS
- [ ] Set secure CORS policies
- [ ] Implement rate limiting
- [ ] Add request logging and monitoring
- [ ] Use database connection pooling
- [ ] Enable PostgreSQL SSL mode
- [ ] Implement password complexity requirements
- [ ] Add account lockout after failed login attempts
- [ ] Set up automated security scanning

---

## 🧪 Testing

### Run Tests
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run with verbose output
go test -v ./...
```

### Test Structure
```
internal/
├── handlers/
│   └── user_handler_test.go
├── services/
│   └── auth_service_test.go
└── repository/
    └── user_repository_test.go
```

---

## 🚀 Deployment

### Using Docker

1. **Build the Docker image**
```bash
docker build -t chat-app:latest .
```

2. **Run with Docker Compose**
```yaml
version: '3.8'

services:
  postgres:
    image: postgres:14-alpine
    environment:
      POSTGRES_DB: chat_app
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  app:
    build: .
    environment:
      DB_HOST: postgres
      DB_PORT: 5432
      DB_NAME: chat_app
      DB_USER: postgres
      DB_PASSWORD: postgres
      JWT_SECRET: your_production_secret
    ports:
      - "8080:8080"
    depends_on:
      - postgres

volumes:
  postgres_data:
```

```bash
docker-compose up -d
```

### Production Deployment Platforms
- **AWS**: EC2, RDS, Elastic Beanstalk
- **Heroku**: Easy deployment with PostgreSQL add-on
- **DigitalOcean**: App Platform or Droplets
- **Google Cloud**: Cloud Run, Cloud SQL
- **Railway**: Simple deployment with PostgreSQL

---

## 🛠️ Development

### Project Standards
- Follow Go best practices and idioms
- Write clean, readable, and maintainable code
- Add tests for new features
- Update documentation when adding features
- Use meaningful commit messages

### Code Style
```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Vet code
go vet ./...
```

### Adding New Features

1. Create feature branch
2. Implement changes following clean architecture
3. Add tests
4. Update documentation
5. Submit pull request

---

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Contribution Guidelines
- Follow the existing code style
- Write clear commit messages
- Add tests for new functionality
- Update documentation as needed
- Ensure all tests pass before submitting PR

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👤 Author

**Your Name**
- GitHub: [@yourusername](https://github.com/yourusername)
- LinkedIn: [Your Name](https://linkedin.com/in/yourprofile)
- Email: your.email@example.com

---

## 🙏 Acknowledgments

- [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router and URL matcher
- [golang-migrate](https://github.com/golang-migrate/migrate) - Database migrations
- [jwt-go](https://github.com/golang-jwt/jwt) - JWT implementation for Go
- Go community for excellent documentation and resources

---

## 📚 Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [JWT Introduction](https://jwt.io/introduction)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [REST API Best Practices](https://restfulapi.net/)

---

## 📊 Project Status

This project is actively maintained and open for contributions. Check the [Issues](https://github.com/yourusername/chat-app/issues) page for planned features and known bugs.

**Current Version**: 1.0.0  
**Last Updated**: January 2026

---

<div align="center">

**If you found this project helpful, please consider giving it a ⭐**

Made with ❤️ and Go

</div>
