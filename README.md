# Senior Fullstack 2025 - Go Backend

This project is a backend application built with **Go (Golang)** and **Fiber**, designed to demonstrate authentication using **Redis** and vector search capabilities using **Qdrant**.

## 🚀 Tech Stack

- **Language:** Go 1.25+
- **Framework:** [Fiber v2](https://gofiber.io/)
- **Database (Auth):** Redis
- **Vector Database:** Qdrant
- **Authentication:** JWT (JSON Web Tokens)

## 📋 Prerequisites

Before running the application, ensure you have the following installed:

- [Go](https://go.dev/dl/) (version 1.21 or higher recommended)
- [Redis](https://redis.io/) (or a cloud instance)
- [Qdrant](https://qdrant.tech/) (or a cloud instance)

## 🛠️ Installation & Setup

1.  **Clone the repository:**

    ```bash
    git clone <repository-url>
    cd senior-fullstack-2025
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Configure Environment Variables:**
    Create a `.env` file in the root directory based on `.env.example` (if available) or use the following template:

    ```dotenv
    JWT_SECRET=your_super_secret_key

    # Redis Configuration
    REDIS_ADDR=your-redis-host:port
    REDIS_PASSWORD=your-redis-password

    # Qdrant Configuration
    QDRANT_URI=your-qdrant-host (e.g., xyz.eu-central-1-0.aws.cloud.qdrant.io)
    QDRANT_API_KEY=your-qdrant-api-key
    QDRANT_PORT=6334
    ```

## 🏃‍♂️ Running the Application

Start the server using the following command:

```bash
go run cmd/main.go
```

The server will start on `http://localhost:3000`.

## 🔌 API Endpoints

### Authentication (`/auth`)

| Method | Endpoint         | Description           | Body                                                                        |
| :----- | :--------------- | :-------------------- | :-------------------------------------------------------------------------- |
| `POST` | `/auth/register` | Register a new user   | `{"username": "...", "realname": "...", "email": "...", "password": "..."}` |
| `POST` | `/auth/login`    | Login and receive JWT | `{"username": "...", "password": "..."}`                                    |

### Vector Search (`/vector`)

**Note:** All vector endpoints require `Authorization: Bearer <token>` header.

| Method | Endpoint         | Description                                | Body |
| :----- | :--------------- | :----------------------------------------- | :--- |
| `POST` | `/vector/init`   | Initialize collection & insert sample data | None |
| `POST` | `/vector/search` | Search for nearest neighbor article        | None |
| `POST` | `/vector/seed`   | Bulk insert 1000 vectors (Concurrent)      | None |

## 🧪 Testing with Postman

A Postman collection is included in the root directory: `postman_collection.json`.

1.  Import the file into Postman.
2.  Use the **Auth > Register** request to create a user.
3.  Use the **Auth > Login** request to get a token. The collection is scripted to automatically save the token to the environment.
4.  Use the **Vector** requests to test the Qdrant integration.
