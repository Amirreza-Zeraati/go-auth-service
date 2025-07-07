# Go Auth Service

A simple JWT-based authentication service built with Go and the Gin web framework.

## 🔐 Features

- User Signup
- User Login with JWT generation
- Protected route using middleware

## 🛠 Tech Stack

- Go (Golang)
- Gin framework
- JWT
- PostgreSQL (or your preferred database)
- Environment configuration via `.env`

---

## 🚀 Endpoints

| Method | Endpoint    | Description                 |
|--------|-------------|-----------------------------|
| POST   | /signup     | Register a new user         |
| POST   | /login      | Authenticate and get a JWT  |
| GET    | /private    | Access protected resource   |

---

## ⚙️ Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL
- `.env` file in the root directory

### Example `.env`
```bash
DB_URL=postgres://user:password@localhost:5432/dbname
SECRET=your_jwt_secret_key
```

---

### Run the Server

```bash
go run main.go
```

## 🧪 Testing
Use Postman, Httpie, curl, or any REST client to test the endpoints. Include the JWT token in the Authorization header for the /private route

