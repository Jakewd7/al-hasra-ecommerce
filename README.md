<p align="center">
  <img src="https://readme-typing-svg.demolab.com?font=Fira+Code&size=24&pause=1200&color=00C853&center=true&vCenter=true&width=800&lines=Gin+Backend+Templater;Modular+%7C+Fast+%7C+Secure;Built+with+Go+%2B+Gin" />
</p>

<p align="center">
  <img src="https://img.shields.io/badge/go-1.24-blue" />
  <img src="https://img.shields.io/badge/framework-gin-green" />
  <img src="https://img.shields.io/badge/database-postgresql-blue" />
  <img src="https://img.shields.io/badge/auth-JWT%20RSA-brightgreen" />
  <img src="https://img.shields.io/badge/license-MIT-purple" />
</p>

# Jake's templates
A brief description of what this project does and who it's for
# Gin Backend Templater
This project is a modular backend templater built with Go and Gin, designed to accelerate backend development by providing a clean architecture, reusable module generator, and built-in authentication system.
It supports rapid module scaffolding, JWT key generation, PostgreSQL integration, and API documentation via Swagger.
## ✨ Key Features
- Modular architecture using Gin Framework
- Automatic module generator (handler, service, repository, model, routes)
- Built-in Authentication & Authorization
  - Cookie-based authentication
  - Bearer token authentication
  - Refresh token support
- Users & Roles management included by default
- PostgreSQL database support
- Centralized response handling
- Swagger API documentation
- JWT key pair generator (RSA)
## 🛠️ Available Commands
### 1. Generate New Module
This command automatically scaffolds a new module inside internal/modules/.
go run cmd/template/main.go


Generated structure per module:
```bash
module_name/
├── handler.go
├── service.go
├── repository.go
├── model.go
└── routes.go
```

This ensures consistency and separation of concerns across all modules.
## 2. Generate JWT RSA Keys
This command generates RSA private and public keys for JWT signing and verification.
go run cmd/keygen/main.go


These keys are used by the authentication system for:
- Access token signing
- Refresh token validation
## 📁 Project Structure
```
backend/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
│
├── internal/
│   ├── app/
│   │   └── app.go           # Application bootstrap
│   │
│   ├── config/
│   │   ├── config.go        # Environment & app configuration
│   │   └── database.go      # PostgreSQL connection setup
│   │
│   ├── middleware/
│   │   ├── auth.go          # Authentication & authorization middleware
│   │   └── ngecor_euy.go    # CORS configuration
│   │
│   ├── routes/
│   │   └── routes.go        # Global route registration
│   │
│   └── modules/
│       ├── auth/            # Authentication module
│       ├── user/            # User management module
│       └── announcement/    # Announcement module
│
├── pkg/
│   ├── response/
│   │   └── response.go      # Standardized API responses
│   └── utils/
│       └── jwt.go           # JWT utilities
│
├── go.mod
└── go.sum
```

🧩 Module Architecture
Each module follows a clean and consistent structure:
- handler.go
Handles HTTP requests and responses.
- service.go
Contains business logic.
- repository.go
Handles database operations.
- model.go
Defines database models and entities.
- routes.go
Registers module-specific routes.
This architecture improves maintainability, scalability, and testability.
## 🗄️ Database

- **Database**: PostgreSQL  
- **Connection**: Managed via `internal/config/database.go`  
- Fully compatible with enterprise-scale database schemas  
- Built-in tables include:
  - Users
  - Roles
  - Authentication tokens



## 🔐 Authentication System

This project includes a complete authentication system with the following features:

- Cookie-based authentication
- Bearer token authentication
- Refresh token mechanism
- Role-Based Access Control (RBAC)
- JWT signed using RSA key pairs

Supports both **web-based** and **API-based** authentication workflows.

## 📘 API Documentation (Swagger)

Swagger UI is available once the server is running.

This provides:
- Interactive API exploration
- Request/response schema visualization
- Authentication testing

## 🚀 Getting Started

1. Configure environment variables  
2. Set up the PostgreSQL database  
3. Generate JWT RSA keys  
4. Run the server

```bash
go run cmd/server/main.go
```
