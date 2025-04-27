## Go Merchants Backend API

---

## 📌 Technical Requirements

### Tech Stack

| Komponen                     | Keterangan                                        |
| :--------------------------- | :------------------------------------------------ |
| Go (Golang) 1.21+            | Bahasa utama backend                              |
| net/http                     | Core package untuk HTTP server                    |
| encoding/json                | Parsing dan encode/decode JSON                    |
| github.com/joho/godotenv     | Untuk load .env file ke environment variables     |
| github.com/golang-jwt/jwt/v5 | Untuk implementasi JWT Authentication             |
| github.com/bxcodec/faker/v4  | Untuk generate data dummy customers dan merchants |
| github.com/google/uuid       | Untuk generate unique ID untuk History            |

### Folder Structure

| Folder            | Isi                                                                 |
| :---------------- | :------------------------------------------------------------------ |
| `main.go`         | File utama `main.go` untuk run server                               |
| `src/controller/` | Semua HTTP handler logic (Login, Logout, Payment)                   |
| `src/service/`    | Business Logic                                                      |
| `src/repository/` | Baca/tulis file JSON                                                |
| `src/model/`      | Semua model struct (Customer, Merchant, History, Request, Response) |
| `src/routes/`     | Routing API endpoint per module                                     |
| `src/middleware/` | Middleware (JWT Authorization, CORS, Rate Limiting, IP Whitelist)   |
| `src/utils/`      | Helper functions (Encrypt/Decrypt, HashPassword, JWT utils)         |
| `data/`           | File JSON data (customers.json, merchants.json, history.json)       |

---

## 📌 Cara Menjalankan Aplikasi (Running Instructions)

### 1. Siapkan Environment

Buat file `.env` di root project:

```env
JWT_SECRET_KEY=supersecretjwtkey
ENCRYPTION_SECRET_KEY=12345678901234567890123456789012
```

> Note: ENCRYPTION_SECRET_KEY wajib 32 karakter.

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Jalankan Server

```bash
go run main.go -mode=server
```

Server jalan di:

```
http://localhost:8080
```

Semua API lewat prefix:

```
http://localhost:8080/api/v1/
```

---

## 📌 API Documentation

### 🔒 Login

- **POST** `/api/v1/login`

**Request Body:**

```json
{
  "username": "your_username",
  "password": "password123"
}
```

**Response Success:**

```json
{
  "message": "Login successful",
  "token": "jwt_token_here",
  "customer": {
    "id": "customer-001",
    "username": "your_username",
    "balance": 50000
  }
}
```

### 🔒 Logout

- **POST** `/api/v1/logout`
- Headers: `Authorization: Bearer <jwt_token_here>`

**Response Success:**

```json
{
  "message": "Logout successful"
}
```

### 💳 Payment

- **POST** `/api/v1/payment`
- Headers: `Authorization: Bearer <jwt_token_here>`

**Request Body:**

```json
{
  "merchant_id": "merchant-001",
  "amount": 10000
}
```

**Response Success:**

```json
{
  "message": "Payment successful"
}
```

**Possible Errors:**

| Case                 | Response                              |
| :------------------- | :------------------------------------ |
| Invalid Merchant     | `{ "error": "merchant not found" }`   |
| Insufficient Balance | `{ "error": "insufficient balance" }` |
| Unauthorized         | `{ "error": "Unauthorized" }`         |

### 📓 Get Customers (Optional Testing)

- **GET** `/api/v1/customers`

### 📓 Get Merchants

- **GET** `/api/v1/merchants`

---

## 🚀 Security Implemented

| Feature                                     | Status |
| :------------------------------------------ | :----- |
| JWT Authentication                          | ✅     |
| CORS Middleware                             | ✅     |
| Rate Limit Middleware (1 request/2s per IP) | ✅     |
| IP Whitelist Middleware                     | ✅     |
| AES Encryption Merchant Account Number      | ✅     |

---

## 📚 Good Luck & Happy Hacking! 🚀
