---

## 🔐 Auth Service

this service handles:

- User signup (`/signup`)
- User login with JWT generation (`/login`)
- Protected route requiring JWT (`/private`)

### 📦 Dependencies
- `github.com/gin-gonic/gin`
- `https://github.com/githubnemo/CompileDaemon`
- Gorm | sqlite | postgres
- JWT package (e.g., `github.com/golang-jwt/jwt`)
- `.env` configuration for secrets and DB connection

### 🔧 Endpoints

```bash
POST   /signup    # Create a new user
POST   /login     # Login and receive a JWT token
GET    /private   # Access protected resource (JWT required)
