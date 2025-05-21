# Go Blog API

## About the Project

This project is a backend API for a blogging platform, built with Go. It provides functionalities for managing blog posts and handles user authentication using JWTs. The application connects to a MongoDB database to store and retrieve data.

## Features

*   **Blog Post Management:**
    *   Create new blog posts.
    *   Retrieve a list of all blog posts.
    *   Retrieve a single blog post by its ID.
*   **User Authentication:**
    *   User registration and login (inferred from project structure: `AuthHandler`, `User` model with password hashing).
    *   JWT (JSON Web Token) based authentication for securing API endpoints.
*   **API Endpoints:**
    *   `/api/auth`: For authentication-related operations.
    *   `/api/blog`: For creating and listing blog posts.
    *   `/api/blog/{id}`: For retrieving a specific blog post.
    *   `/api/protected`: An example of a protected endpoint requiring JWT authentication.

## Technologies Used

*   **Programming Language:** [Go](https://golang.org/)
*   **Database:** [MongoDB](https://www.mongodb.com/) (currently configured to connect to `mongodb://localhost:27017`)
*   **Authentication:** JWT (JSON Web Tokens)
*   **Dependency Management:** Go Modules (`go.mod`, `go.sum`)
*   **Environment Variables:** Uses a `.env` file for configuration (e.g., `JWT_SECRET`).

## Project Structure

The project follows a layered architecture for separation of concerns:

*   `cmd/main.go`: Main application entry point, server initialization.
*   `internal/database/`: Database connection setup (e.g., `mongo.go`).
*   `internal/handler/`: HTTP request handlers (e.g., `blog.go` for blog API, `auth.go` for authentication).
*   `internal/middleware/`: Middleware functions (e.g., `jwt.go` for JWT authentication).
*   `internal/model/`: Data structure definitions (e.g., `post.go`, `user.go`).
*   `internal/service/`: Business logic layer (e.g., `post.go`).
*   `internal/repository/`: Data access layer, interacting with the database.
*   `internal/pkg/auth`: Utility package for authentication related functions.

## Getting Started

### Prerequisites

*   Go (version X.X.X - *Specify version if known, otherwise can be general*)
*   MongoDB (running instance)
*   A `.env` file in the root directory with necessary environment variables, e.g.:
    ```
    JWT_SECRET=your_jwt_secret_key_here
    ```

### Installation & Running

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd <repository-directory>
    ```
2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```
3.  **Run the application:**
    ```bash
    go run cmd/main.go
    ```
    The server will start on `http://localhost:8080`.

## API Endpoints

*(This section can be expanded with more details on request/response formats if available)*

*   **Authentication:**
    *   `POST /api/auth/register` (Example - *Actual endpoint might vary*)
    *   `POST /api/auth/login` (Example - *Actual endpoint might vary*)
*   **Blog Posts:**
    *   `GET /api/blog`: Get all blog posts.
    *   `POST /api/blog`: Create a new blog post.
        *   Request Body: JSON with post details (e.g., `{"title": "My First Post", "content": "Hello world!"}`)
    *   `GET /api/blog/{id}`: Get a blog post by its ID.
*   **Protected Route:**
    *   `GET /api/protected`: Access a protected resource (requires JWT in Authorization header).

---

*This README was generated based on an understanding of the project's Go codebase.*
