# 🎓 Students REST API (Go)

A simple and modular **RESTful API** built in **Go (Golang)** for managing student data.  
This project follows a clean architecture approach — separating configuration, HTTP handlers, storage, and utilities — making it easy to maintain and extend.

---

## 🚀 Features

- CRUD operations for Students (Create, Read, Update, Delete)
- SQLite database integration
- YAML-based configuration support
- Modular and scalable folder structure
- Easy to extend with new routes or database providers

---

## 🧱 Tech Stack

- **Language:** Go (Golang)
- **Database:** SQLite
- **Configuration:** YAML
- **Architecture:** Modular / Clean Architecture
- **HTTP Routing:** Go’s standard `net/http`

---

## 📂 Project Structure

```

Go-API
├── cmd
│   └── students-api
│       └── students-api.go
├── config
│   └── local.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── http
│   │   └── handlers
│   │       └── student
│   │           └── student.go
│   ├── storage
│   │   ├── sqlite
│   │   │   └── sqlite.go
│   │   └── storage.go
│   ├── types
│   │   └── types.go
│   └── utils
│       └── response
│           └── response.go
├── README.md
└── storage
    └── storage.db

```


---

## ⚙️ Configuration

The `config/local.yaml` file contains local settings for your project.  
Example:

```yaml
server:
  port: 8080

database:
  driver: sqlite
  path: "./storage/storage.db"

```

---
🏃‍♂️ How to Run Locally
1. Clone the repository
git clone https://github.com/your-username/Go-API.git
cd Go-API

2. Install dependencies
go mod tidy

3. Run the application
go run ./cmd/students-api


Server will start on:
➡️ http://localhost:8082

---
```
📡 Example API Endpoints
Method	Endpoint	Description
GET	/students	Get all students
GET	/students/{id}	Get a student by ID
POST	/students	Add a new student
PUT	/students/{id}	Update an existing student
DELETE	/students/{id}	Delete a student
Example Request (POST /students)
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-d '{
  "name": "John Doe",
  "age": 22,
  "email": "john@example.com"
}'
```
---

🧩 Future Improvements

JWT authentication

Swagger API documentation

Docker container support

Migration support for multiple databases

🤝 Contributing

Contributions are welcome!
If you’d like to improve this project, feel free to fork and submit a pull request.


