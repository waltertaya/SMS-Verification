# 📲 SMS Verification API using Gin + Twilio

This project is a backend authentication API built using [Gin Web Framework](https://github.com/gin-gonic/gin) with Two-Factor Authentication (2FA) via **Twilio SMS API**.

---

## 🚀 Features

- User registration with hashed password storage
- Secure user login with verification check
- 2FA SMS code request and verification via Twilio
- API routes documented at `/api/v2/docs`
- CORS enabled for cross-origin frontend support
- 404 handler for all undefined routes with UI feedback

---

## 🧰 Technologies Used

- Go (Golang)
- Gin Gonic Web Framework
- Twilio Verify API
- MySQL (via GORM)
- HTML templates for documentation and 404 UI

---

## 📦 Installation & Setup

1. **Clone the Repository**

```bash
git clone https://github.com/waltertaya/SMS-Verification.git
cd SMS-Verification
```

2. **Create a `.env` file**

Add your environment variables:

```env
DB_URL=your_mysql_url
TWILIO_VERIFY_SERVICE_SID=your_twilio_verify_sid
TWILIO_ACCOUNT_SID=your_account_sid
TWILIO_AUTH_TOKEN=your_auth_token
```

3. **Run the Server**

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

---

## 🔐 API Endpoints

All routes prefixed with `/api/v2`

| Method | Endpoint                 | Description                     |
|--------|--------------------------|---------------------------------|
| GET    | `/api`                   | Welcome message                 |
| POST   | `/auth/register`         | Register new user               |
| POST   | `/auth/login`            | Login existing user             |
| POST   | `/auth/2fa/code`         | Request 2FA code via Twilio     |
| POST   | `/auth/2fa/verify`       | Verify 2FA code and login       |

📄 Visit [http://localhost:8080/api/v2/docs](http://localhost:8080/api/v2/docs) for full API documentation.

---

## 🛠 Directory Structure

```
SMS-Verification/       

.
├── controllers                   # API handlers
│   ├── userHandler.go
│   └── verifyUsers.go
├── db                            # DB connection setup
│   └── connectDB.go
├── go.mod
├── go.sum
├── initializers                  # Environment loader
│   └── loadEnv.go
├── main.go                       # Entry point
├── migrate
│   └── migrate.go
├── models                        # User model
│   └── userModel.go
├── README.md
├── templates                     # index.html, 404.html
│   ├── 404.html
│   └── index.html
└── utils                         # Utility functions
    ├── logErrors.go
    ├── passwordHashing.go
    └── uuidGen.go

8 directories, 15 files
```

---

## ❌ 404 Handler

Any undefined route will render a `404.html` page showing the incorrect URL and a reminder about using Twilio SMS for verification.

---

## 📞 Twilio Integration

This app uses [Twilio Verify API](https://www.twilio.com/docs/verify/api) to send verification codes to users’ phone numbers via SMS.

---

## 🧪 Sample Request - Code Send

```http
POST /api/v2/auth/2fa/code
Content-Type: application/json

{
  "phone": "+254712345678"
}
```

## ✅ Sample Request - Code Verify

```http
POST /api/v2/auth/2fa/verify
Content-Type: application/json

{
  "phone": "+254712345678",
  "code": "123456",
  "id": "user-unique-id"
}
```

---

## 👨‍💻 Author

**Walter Onyango**  
GitHub: [@waltertaya](https://github.com/waltertaya)

---

## 📝 License

MIT License
