# Go Boilerplate Starter Kit

A comprehensive Go web application boilerplate with modern architecture, dependency injection, database migrations, internationalization, monitoring, and more.

---

## 🚀 Features

### Core Features
- **Gin Web Framework** - Fast HTTP web framework
- **Dependency Injection** - Using Uber's Dig for clean architecture
- **Database Support** - MySQL, PostgreSQL, and MongoDB support
- **Database Migrations** - Using Goose for schema management
- **Internationalization (i18n)** - Multi-language support
- **Validation** - Request validation with custom error messages
- **JWT Authentication** - JWT token-based authentication
- **Monitoring** - Prometheus metrics and Grafana dashboards
- **Docker Support** - Complete Docker setup with multi-stage builds

### Development Features
- **CLI Commands** - Built-in commands for migrations, seeding, and module generation
- **Code Generation** - Auto-generate modules, repositories, and containers
- **Testing Support** - Built-in testing structure
- **API Response Helpers** - Standardized API responses
- **Pagination** - Built-in pagination support
- **Middleware** - Authentication and i18n middleware

---

## 📋 Prerequisites
- Go 1.23.4
- Docker and Docker Compose
- MySQL/PostgreSQL (optional, for database features)

---

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/adityarifqyfauzan/go-boilerplate.git
   cd go-boilerplate
   ```
2. **Install dependencies**
   ```bash
   go mod download
   ```
3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```
4. **Run with Docker (Recommended)**
   ```bash
   docker-compose up --build
   # Or run in background
   docker-compose up -d
   ```
5. **Run locally**
   ```bash
   go run main.go
   # Or build and run
   go build -o main .
   ./main
   ```

---

## 🏗️ Project Structure

```
go-boilerplate/
├── config/                 # Configuration files
├── internal/               # Internal application code
│   ├── bootstrap/          # Application bootstrap
│   ├── command/            # CLI commands
│   ├── container/          # Dependency injection containers
│   ├── database/           # Database migrations and seeders
│   │   ├── migrations/     # Database migration files
│   │   └── seeders/        # Database seeder files
│   ├── helper/             # Utility helpers
│   ├── model/              # Data models
│   ├── module/             # Feature modules
│   │   ├── authentication/ # Authentication module
│   ├── repository/         # Data access layer
│   └── routes/             # Route definitions
├── locales/                # Internationalization files
├── mysql/                  # MySQL-specific files
├── pkg/                    # Public packages
│   ├── apm/                # Application performance monitoring
│   ├── jwt/                # JWT utilities
│   ├── middleware/         # HTTP middleware
│   ├── opentelemetry/      # OpenTelemetry utilities
│   ├── translator/         # Translation utilities
│   └── validator/          # Validation utilities
├── docker-compose.yml      # Docker services configuration
├── Dockerfile              # Multi-stage Docker build
├── go.mod                  # Go module definition
├── main.go                 # Application entry point
├── Makefile                # Build and development commands
├── dbconfig.yml            # Database configuration
└── prometheus.yml          # Prometheus configuration
```

---

## 🗄️ Database Setup

### Using Docker (Recommended)
```bash
docker-compose up
```
- Database available at:
  - Host: `mysql` (from app container) or `localhost:3306`
  - Database: `go_starter_kit`
  - User: `myuser`
  - Password: `mypass`
  - Root Password: `rootpass`

**.env example:**
```env
DB_HOST=mysql
DB_PORT=3306
DB_NAME=go_starter_kit
DB_USER=myuser
DB_PASS=mypass
DB_DRIVER=mysql
```

### Manual Database Setup
1. Create a MySQL/PostgreSQL database
2. Run migrations:
   ```bash
   make migrate-up
   ```

---

## 📊 Database Migrations

### Available Commands
```bash
make migrate-up                         # Run migrations
make migrate-up-to version={version}    # Run up to specific version
make migrate-down                       # Rollback all migrations
make migrate-down-to version={version}  # Rollback to specific version
make migrate-status                     # Check migration status
make migrate-create name={name}         # Create new migration
make migrate-refresh                    # Refresh migrations (down + up)
```

### CLI Usage
```bash
go run main.go migrate:up
go run main.go migrate:create <name>
go run main.go migrate:status
```

### Current Migrations
- `20250704001636_create_user_statuses_table.go` - User statuses table
- `20250704001736_create_users_table.go` - Users table
- `20250704002839_create_roles_table.go` - Roles table
- `20250704023449_create_user_roles_table.go` - User roles pivot table
- `20250704025613_create_user_details_table.go` - User details table
- `20250704140231_create_user_status_histories_table.go` - User status history table

---

## 🌱 Database Seeders

### Run all seeders
```bash
make seeder
```

### Run specific seeder
```bash
make seeder-only name=UserSeeder
```

### CLI Usage
```bash
go run main.go seeder
go run main.go seeder --only UserSeeder
```

### Available Seeders
- `role.go` - Role data seeder
- `user_status.go` - User status data seeder
- `user.go` - User data seeder

---

## 🧩 Module Generation

Generate new modules with all necessary files:
```bash
go run main.go make:module product
```
Creates:
- `internal/module/product/model.go`
- `internal/module/product/dto.go`
- `internal/module/product/handler.go`
- `internal/module/product/service.go`
- `internal/module/product/route.go`
- `internal/module/product/container.go`
- `internal/module/product/local_repository.go`

---

## 🌐 Internationalization (i18n)

Translation files are in `locales/`:
- `active.en.json` (English)
- `active.id.json` (Indonesian)
- `active.ja.json` (Japanese)

**Usage in code:**
```go
localizer := translator.NewLocalizer("en")
i18n := translator.NewTranslator(localizer)
message := i18n.T("hello", map[string]any{"Name": "World"})
```

---

## 🔐 Authentication & User Endpoints

### Endpoints
- `POST /api/v1/authentication/login` — User login
- `POST /api/v1/authentication/register` — User registration
- `POST /api/v1/authentication/forgot-password` — Forgot password (not implemented yet)
- `POST /api/v1/authentication/refresh-token` — Refresh JWT access token
- `GET /api/v1/authentication/me` — Get current user info (requires authentication)

### Example Requests

#### Login
```bash
POST /api/v1/authentication/login
{
  "email": "user@example.com",
  "password": "password"
}
```

#### Register
```bash
POST /api/v1/authentication/register
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password",
  "password_confirmation": "password"
}
```

#### Refresh Token
```bash
POST /api/v1/authentication/refresh-token
{
  "refresh_token": "<refresh_token>"
}
```

#### Me
```bash
GET /api/v1/authentication/me
# Requires Authorization: Bearer <token>
```

#### Forgot Password
```bash
POST /api/v1/authentication/forgot-password
# Not implemented yet
```

### JWT Claims Structure
```go
type Claims struct {
    UserID   int      `json:"user_id"`
    Email    string   `json:"email"`
    Username string   `json:"username"`
    Roles    []string `json:"roles"`
}
```

### Role-Based Access Control
```go
// Require admin role
userRoute.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
// Require user role
userRoute.Use(middleware.AuthMiddleware(), middleware.UserMiddleware())
// Require any of multiple roles
userRoute.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin","moderator"))
// Require specific role
userRoute.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("editor"))
```

---

## 📈 Monitoring

- **Prometheus Metrics:** `/metrics`
- **Grafana Dashboard:** [http://localhost:3000](http://localhost:3000) (admin/admin)
- **Prometheus:** [http://localhost:9090](http://localhost:9090)

---

## 🐳 Docker Commands

```bash
# Build the application
docker build -t go-boilerplate .
# Run with docker-compose
docker-compose up --build
# Run in background
docker-compose up -d
# Stop services
docker-compose down
# Run migrations in container
docker-compose exec app go run main.go migrate:up
# Run seeders in container
docker-compose exec app go run main.go seeder
```

---

## 🧪 Testing

```bash
# Run all tests
go test ./...
# Run specific test
go test ./internal/module/authentication
# Run with coverage
go test -cover ./...
```

---

## 📄 API Endpoints

### Health Check
```bash
GET /health
```

### Authentication
```bash
POST /v1/authentication/login
POST /v1/authentication/register
POST /v1/authentication/forgot-password
POST /v1/authentication/refresh-token
GET /v1/authentication/me
```

### Internationalization Example
```bash
GET /hello/World    # Returns localized greeting
```

---

## 🔧 Configuration

Create a `.env` file with:
```env
# Application
APP_NAME=go-starter-kit
APP_PORT=5001
APP_ENV=development
# Database
DB_HOST=localhost
DB_PORT=3306
DB_NAME=go_starter_kit
DB_USER=myuser
DB_PASS=mypass
DB_DRIVER=mysql
# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRY=60
# MongoDB (Optional)
MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_SECURITY=false
MONGO_USER=
MONGO_PASS=
```

---

## 🚀 Deployment

### Production Build
```bash
docker build -t go-boilerplate:prod .
docker run -p 5001:5001 go-boilerplate:prod
```

### Environment-specific builds
```bash
# Development
docker build -t go-boilerplate:dev .
# Production
docker build -t go-boilerplate:prod .
```

---

## 📚 Available Commands

### Make Commands
```bash
make migrate-up
make migrate-down
make migrate-down-to version={version}
make migrate-status
make migrate-create name={name}
make migrate-refresh
make seeder
make seeder-only name={name}
make module-create name={name}
```

### CLI Commands
```bash
go run main.go migrate:up
go run main.go migrate:create <name>
go run main.go seeder
go run main.go make:module <name>
```

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

---

## 📄 License

This project is licensed under the MIT License.

---

## 🆘 Support

For support and questions:
- Create an issue in the repository
- Check the documentation
- Review the code examples

---

**Happy Coding! 🎉** 