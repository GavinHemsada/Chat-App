# Chat App

A real-time chat application built to practice backend development, authentication, database design, and scalable API architecture.  
This project supports user registration, secure login, chat rooms, and message handling, with a clean and extensible structure.

---

## 🚀 Features

- User registration and login
- JWT-based authentication
- Role-ready architecture (easy to extend)
- Chat rooms (public / group / direct)
- Message sending and retrieval
- Secure protected APIs
- PostgreSQL database integration
- Clean backend folder structure
- Ready for WebSocket real-time chat extension

---

## 🧰 Tech Stack

### Backend
- Go (Golang)
- Gorilla Mux (Router)
- PostgreSQL
- JWT Authentication
- bcrypt (password hashing)

### Tools
- golang-migrate (database migrations)
- Docker (optional)
- Git

---

## 📁 Project Structure

Chat-App/
│
├── cmd/
│ └── server/
│ └── main.go
│
├── internal/
│ ├── config/ # Environment configuration
│ ├── database/ # Database connection & migrations
│ ├── handlers/ # HTTP handlers (controllers)
│ ├── middleware/ # JWT & auth middleware
│ ├── models/ # Database models
│ ├── repository/ # DB queries
│ ├── services/ # Business logic
│ └── websocket/ # WebSocket logic (if enabled)
│
├── migrations/ # SQL migration files
├── .env.example
├── go.mod
├── go.sum
└── README.md

---

## ✅ Requirements

- Go **1.20+**
- PostgreSQL **14+**
- Git

---

## ⚙️ Environment Variables

Create a `.env` file in the project root.

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_NAME=chat_app
DB_USER=postgres
DB_PASSWORD=postgres

JWT_SECRET=super_secret_key
JWT_EXPIRES_IN=24h
