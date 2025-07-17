# Go Boilerplate Starter Kit

A comprehensive Go web application boilerplate with modern architecture, dependency injection, database migrations, internationalization, monitoring, and more.

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

## 📋 Prerequisites
- Go 1.23.4
- Docker and Docker Compose
- MySQL/PostgreSQL (optional, for database features)

## 🛠️ Installation

### 1. Clone the repository
```bash
git clone https://github.com/adityarifqyfauzan/go-boilerplate.git
cd go-boilerplate
```

### 2. Install dependencies
```bash
go mod download
```

### 3. Set up environment variables
```bash
cp .env.example .env
# Edit .env with your configuration
```

### 4. Run with Docker (Recommended)
```bash
# Build and start all services
docker-compose up --build

# Run in background
docker-compose up -d
```

### 5. Run locally
```bash
# Start the application
go run main.go

# Or build and run
go build -o main .
./main
```

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

## 🗄️ Database Setup

### Using Docker (Recommended)
The project includes a complete Docker setup with MySQL, Prometheus, and Grafana:

```bash
# Start with all services
docker-compose up

# Database will be available at:
# Host: mysql (from app container) or localhost:3306
# Database: go_starter_kit
# User: myuser
# Password: mypass
# Root Password: rootpass
```

**Important**: When using Docker, make sure your `.env` file has the correct database settings:
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

## 📊 Database Migrations

### Available Commands
```bash
# Run all migrations
make migrate-up

# Run migrations up to specific version
make migrate-up-to version=20250704025613

# Rollback all migrations
make migrate-down

# Rollback to specific version
make migrate-down-to version=20250704025613

# Check migration status
make migrate-status

# Create new migration
make migrate-create name=create_new_table

# Refresh migrations (down + up)
make migrate-refresh
```

### Using CLI directly
```bash
# Run migrations
go run main.go migrate:up

# Create new migration
go run main.go migrate:create create_users_table

# Check status
go run main.go migrate:status
```

### Current Migrations
- `20250704001636_create_user_statuses_table.go` - User statuses table
- `20250704001736_create_users_table.go` - Users table
- `20250704002839_create_roles_table.go` - Roles table
- `20250704023449_create_user_roles_table.go` - User roles pivot table
- `20250704025613_create_user_details_table.go` - User details table
- `20250704140231_create_user_status_histories_table.go` - User status history table

## 🌱 Database Seeders

### Run all seeders
```bash
make seeder
```

### Run specific seeder
```bash
make seeder-only name=UserSeeder
```

### Using CLI
```bash
# Run all seeders
go run main.go seeder

# Run specific seeder
go run main.go seeder --only UserSeeder
```

### Available Seeders
- `role.go` - Role data seeder
- `user_status.go` - User status data seeder

## 🧩 Module Generation

Generate new modules with all necessary files:

```bash
# Generate a new module
go run main.go make:module product

# This creates:
# - internal/module/product/model.go
# - internal/module/product/dto.go
# - internal/module/product/handler.go
# - internal/module/product/service.go
# - internal/module/product/route.go
# - internal/module/product/container.go
# - internal/module/product/local_repository.go
```

## 🌐 Internationalization (i18n)

The project supports multiple languages. Translation files are in `locales/`:

```bash
# Available languages
locales/
├── active.en.json    # English
├── active.id.json    # Indonesian
└── active.ja.json    # Japanese
```

### Using translations in code
```go
i18n := translator.NewTranslator(c.Value("localizer").(*i18n.Localizer))
message := i18n.T("hello", map[string]any{
    "Name": "World",
})
```

## 🔐 Authentication

JWT-based authentication is included with support for multiple user roles:

### Login
```bash
POST /v1/authentication/login
{
 email   : "user@example.com",
 password: "password"
}
```

### JWT Claims Structure
The JWT token contains the following claims:
```go
type Claims struct {
    UserID   int      `json:"user_id"`
    Email    string   `json:"email"`
    Username string   `json:"username` 
    Roles    []string `json:"roles"`  // Multiple roles support
}
```

### Protected Routes
Use the auth middleware for protected endpoints:
```go
userRoute.Use(middleware.AuthMiddleware())
```

### Role-Based Access Control
The middleware supports multiple roles and flexible role checking:

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

### Accessing User Information
Helper functions to retrieve user information from context:

```go
// Get user ID
userID, exists := middleware.GetUserID(c)

// Get user email
email, exists := middleware.GetUserEmail(c)

// Get user claims (full JWT claims)
claims, exists := middleware.GetUserClaims(c)

// Get user roles (returns []string)
roles, exists := middleware.GetUserRoles(c)
```

## 📈 Monitoring

### Prometheus Metrics
Metrics are available at `/metrics`:
- HTTP request count and duration
- Database connection stats
- Goroutine count

### Grafana Dashboard
Access Grafana at `http://localhost:3000`:
- Username: `admin`
- Password: `admin`

### Prometheus
Access Prometheus at `http://localhost:9090`

## 🐳 Docker Commands

### Build and run
```bash
# Build the application
docker build -t go-boilerplate .

# Run with docker-compose
docker-compose up --build

# Run in background
docker-compose up -d

# Stop services
docker-compose down
```

### Database operations in Docker
```bash
# Run migrations in container
docker-compose exec app go run main.go migrate:up

# Run seeders in container
docker-compose exec app go run main.go seeder
```

## 🧪 Testing

### Run tests
```bash
# Run all tests
go test ./...

# Run specific test
# (update this to your actual module, e.g. authentication)
go test ./internal/module/authentication

# Run with coverage
go test -cover ./...
```

## 📄 API Endpoints

### Health Check
```bash
GET /health
```

### Authentication
```bash
POST /v1/authentication/login    # Login
```

### Internationalization Example
```bash
GET /hello/World    # Returns localized greeting
```

## 🔧 Configuration

### Environment Variables
Create a `.env` file with:

```env
# Application
APP_NAME=go-starter-kit
APP_PORT=5001
APP_ENV=development

# Database (Required for application)
DB_HOST=localhost
DB_PORT=3306
DB_NAME=go_starter_kit
DB_USER=myuser
DB_PASS=mypass
DB_DRIVER=mysql

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRY=24h

# Monitoring
PROMETHEUS_ENABLED=true

# MongoDB (Optional)
MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_SECURITY=false
MONGO_USER=
MONGO_PASS=
```

## 🚀 Deployment

### Production Build
```bash
# Build for production
docker build -t go-boilerplate:prod .

# Run with production environment
docker run -p 5001:5001 go-boilerplate:prod
```

### Environment-specific builds
```bash
# Development
docker build -t go-boilerplate:dev .

# Production
docker build -t go-boilerplate:prod .
```

## 📚 Available Commands

### Make Commands
```bash
make migrate-up                         # Run migrations
make migrate-down                       # Rollback migrations
make migrate-down-to version={version}  # Rollback specific version
make migrate-status                     # Check migration status
make migrate-create name={name}         # Create new migration
make migrate-refresh                    # Refresh migrations
make seeder                             # Run all seeders
make seeder-only name={name}            # Run specific seeder
make module-create name={name}          # Generate module
```

### CLI Commands
```bash
go run main.go migrate:up
go run main.go migrate:create <name>
go run main.go seeder
go run main.go make:module <name>
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For support and questions:
- Create an issue in the repository
- Check the documentation
- Review the code examples

---

**Happy Coding! 🎉** 