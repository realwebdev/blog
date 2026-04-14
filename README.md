https://goreportcard.com/report/github.com/realwebdev/blog
[![Go Report Card](https://goreportcard.com)]([https://goreportcard.com](https://goreportcard.com/report/github.com/realwebdev/blog))
# 📝 Go Blog Engine

A modular, full stack blog application built in **Go**, following **Clean Architecture principles**.

This project demonstrates a well structured backend serving a **server side rendered frontend**, with clear separation of concerns, CLI tooling, and database migrations.

---

## ✨ Features

* User authentication & authorization
* Admin panel for content management
* Article CRUD operations
* Server-side rendered HTML using Go templates
* PostgreSQL-backed persistence
* Database migrations with Goose
* CLI powered application commands
* Modular architecture with clear boundaries

---

## 🧱 Architecture Overview

The application follows a layered, modular structure inspired by Clean Architecture:

```
├── cmd/                # CLI commands (serve, migrate, seed)
├── internal/
│   ├── domain/         # Core business entities
│   ├── repository/     # Data access layer
│   ├── service/        # Business logic
│   ├── handler/        # HTTP handlers
│   └── middleware/     # HTTP middleware
├── templates/          # Server-side HTML templates
├── static/             # CSS, JS, assets
├── migrations/         # Goose migration files
└── config/             # Configuration management
```

### Design Principles

* Clear separation between domain, service, and infrastructure
* Dependency direction pointing inward (core independent from frameworks)
* CLI based application lifecycle management
* Explicit database migrations
* Minimal global state

---

## 🛠 Tech Stack

* **Language**: Go (1.18+)
* **Architecture**: Modular Service–Repository pattern
* **Database**: PostgreSQL
* **Migrations**: Goose
* **CLI Framework**: Cobra
* **Frontend**: Go HTML Templates
* **Styling**: Bootstrap 4
* **Client-side Enhancements**: jQuery

---

## 🚀 Getting Started

### Prerequisites

* Go 1.18+
* PostgreSQL (running locally or via Docker)

---

## 🔧 Installation & Setup

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/realwebdev/blog.git
cd blog
```

---

### 2️⃣ Configure Database

Ensure PostgreSQL is running.

Update `config.yml` with your database credentials:

```yaml
database:
  host: localhost
  port: 5432
  user: postgres
  password: postgres
  name: blog
```

---

### 3️⃣ Run Database Migrations

Initialize the database schema:

```bash
go run main.go migrate up
```

---

### 4️⃣ Seed Database (Optional)

Populate the database with sample users and articles:

```bash
go run main.go seed
```

> ⚠️ Default admin password is `password123` unless `ADMIN_PASSWORD` environment variable is set.

---

### 5️⃣ Start the Server

```bash
go run main.go serve
```

Visit:

```
http://localhost:8080
```

(or the port defined in your configuration)

---

## 🔐 Authentication

* Role based access control
* Admin only content management
* Secure password handling (bcrypt recommended if implemented)

---

## 📦 CLI Commands

| Command        | Description                        |
| -------------- | ---------------------------------- |
| `serve`        | Start HTTP server                  |
| `migrate up`   | Apply database migrations          |
| `migrate down` | Roll back migrations               |
| `seed`         | Populate database with sample data |

---

## 🧪 Testing

If applicable:

```bash
go test ./...
```

(Add this section only if you have test coverage — otherwise omit or clarify scope.)

---

## 📌 Design Decisions

* Server-side rendering chosen for simplicity and SEO-friendliness.
* Goose for migration version control.
* Cobra to separate application lifecycle commands cleanly.
* Clean Architecture to ensure framework-independent core logic.

---

## 🚧 Future Improvements

* API versioning
* REST/JSON API alongside SSR
* Integration tests
* Structured logging
* Dockerized deployment
* CI/CD pipeline
* Pagination & search
* Caching layer (Redis)

---

## 👨‍💻 Purpose

This project showcases:

* Backend architecture design in Go
* Database migration management
* Clean separation of business logic
* Full stack Go application development


