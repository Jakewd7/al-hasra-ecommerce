backend/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── app/
│   │   └── app.go
│   │
│   ├── config/
│   │   ├── config.go
│   │   └── database.go
│   │
│   ├── middleware/
│   │   ├── auth.go
│   │   └── cors.go
│   │
│   ├── routes/
│   │   └── routes.go
│   │
│   └── modules/
│       ├── auth/
│       │   ├── handler.go
│       │   ├── service.go
│       │   ├── repository.go
│       │   ├── model.go
│       │   └── routes.go
│       │
│       ├── user/
│       │   ├── handler.go
│       │   ├── service.go
│       │   ├── repository.go
│       │   ├── model.go
│       │   └── routes.go
│       │
│       └── announcement/
│           ├── handler.go
│           ├── service.go
│           ├── repository.go
│           ├── model.go
│           └── routes.go
│
├── pkg/
│   ├── response/
│   │   └── response.go
│   └── utils/
│       └── jwt.go
│
├── go.mod
└── go.sum
