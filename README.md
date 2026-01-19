# Go Blog Engine

A modular, full-stack blog application built with Go, following Clean Architecture principles. This project demonstrates a robust backend structure serving a server-side rendered frontend.

## 🚀 Getting Started

### Prerequisites
*   **Go** (1.18+)
*   **PostgreSQL**

### Installation & Setup

1.  **Clone the repository**
    ```bash
    git clone https://github.com/realwebdev/blog.git
    cd blog
    ```

2.  **Database Setup**
    Ensure PostgreSQL is running and your `config.yml` is updated with the correct credentials.

3.  **Run Migrations**
    Initialize the database schema using Goose.
    ```bash
    go run main.go migrate up
    ```

4.  **Seed Database** (Optional)
    Populate the database with users and sample articles.
    *   *Note: Sets default admin password to `password123` unless `ADMIN_PASSWORD` env var is set.*
    ```bash
    go run main.go seed
    ```

5.  **Start the Server**
    Launch the development server.
    ```bash
    go run main.go serve
    ```
    Visit `http://localhost:8080` (or your configured port) to view the blog.

## 🛠️ Tech Stack

*   **Language**: Go (Golang)
*   **Architecture**: Modular Service-Repository Pattern
*   **Database**: PostgreSQL, Goose (Migrations)
*   **CLI**: Cobra
*   **Frontend**: Go Templates (HTML), Bootstrap 4, jQuery