# 💧 WPS Backend Service

This is the backend for the **Water Purification System (WPS)** platform — an e-commerce and service booking system focused on selling residential water filtration products and managing mobile water cleaning services.

Built with **Go (Golang)** using a layered monorepo architecture, the system provides scalable foundations for managing users, products, orders, service bookings, and payments.

---

## 🏗️ Project Structure

wps-backend/
│
├── cmd/ # Application entry point
│ └── main.go
├── config/ # Config loader and YAML configs
├── internal/ # Core logic
│ ├── controller/ # HTTP handlers
│ ├── middleware/ # Auth, logging
│ ├── model/ # Entities, DTOs
│ ├── repository/ # DB abstractions
│ ├── routes/ # API routing
│ └── service/ # Business logic
├── repos/ # Postgres & Redis implementations
├── migrations/ # SQL migrations
├── pkg/ # Utilities (logger, db, cache, etc.)
└── frontend/ # Frontend (placeholder)


## 🚀 Getting Started

### Prerequisites
- Go 1.20+
- PostgreSQL 14+
- Redis 6+
- Docker (optional, for local dev)

### 1. Clone the repository
```bash
1. Configure the environment
Edit your config/config.yaml file as needed.
server:

database:

redis:

2. Run migrations

3. Start the server

go run cmd/main.go


### 📦 Features
✅ Modular monorepo architecture

✅ REST API with chi

✅ PostgreSQL + Redis support

✅ JWT Authentication (middleware in progress)

✅ Configurable with YAML

### 🧩 Booking & Payment modules coming soon

### 🧠 Tech Stack
Go (Golang)

PostgreSQL via pgx

Redis via go-redis

JWT Auth via jwtauth

Chi Router for HTTP handling

YAML Config for flexibility

📄 License
This project is licensed under the MIT License.

🤝 Contributing
Pull requests and suggestions are welcome! For major changes, please open an issue first to discuss what you’d like to change.

📫 Contact
Muhammadjoni Rahimzod – LinkedIn
