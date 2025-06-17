# Broker Platform Backend – Liquide Assignment

##  Tech Stack
- Go (Gin)
- PostgreSQL with GORM
- JWT Auth (access + refresh token)
- Unit testing with Testify

##  Project Setup

### 1.  Clone the Repository or Unzip
```bash
cd broker-platform
```

### 2. 🛠 Create a `.env` File
```bash
cp .env.example .env
```

Edit `.env`:
```
DATABASE_DSN=host=localhost user=postgres password=yourpassword dbname=broker port=5432 sslmode=disable
JWT_SECRET=supersecret
```

### 3. 🗄️ PostgreSQL Table Setup

Run this SQL to create the `users` table:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);
```

### 4.  Run the Application
```bash
go run cmd/main.go
```

App will start on: http://localhost:8080

### 5.  Authentication Flow
- `POST /signup` → Register user
- `POST /login` → Get access + refresh token
- Use access token as: `Authorization: Bearer <token>`
- `POST /refresh-token` → Get new access token

### 6. 📘 API Endpoints

| Method | Endpoint         | Auth | Description                  |
|--------|------------------|------|------------------------------|
| POST   | /signup          | ❌   | User signup                  |
| POST   | /login           | ❌   | User login                   |
| POST   | /refresh-token   | ❌   | Get new access token         |
| GET    | /holdings        | ✅   | Get holdings (mocked)        |
| GET    | /orderbook       | ✅   | Get orderbook with PNL       |
| GET    | /positions       | ✅   | Get positions with PNL card  |
| GET    | /health          | ❌   | Health check                 |
