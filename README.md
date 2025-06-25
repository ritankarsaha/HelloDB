# 🚀 NLP to SQL - Natural Language Query Processor

A Go-based application that converts natural language queries into SQL statements, featuring PostgreSQL integration, token-based authentication, and AI model integration with Ollama.

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
- [Configuration](#-configuration)
- [Development](#-development)
- [API Documentation](#-api-documentation)
- [Database Management](#-database-management)
- [Testing](#-testing)
- [Deployment](#-deployment)
- [Contributing](#-contributing)

## ✨ Features

- **Natural Language Processing**: Convert natural language queries to SQL
- **AI Integration**: Powered by Ollama with Llama3 model
- **PostgreSQL Support**: Full database integration with migrations
- **Token Authentication**: Secure JWT-based authentication system
- **Cron Jobs**: Scheduled batch processing capabilities
- **SQL Generation**: Automated SQL code generation with SQLC
- **Docker Support**: Containerized deployment ready
- **Vulnerability Scanning**: Built-in security scanning with govulncheck
- **Comprehensive Testing**: Unit tests with coverage reporting
- **Mock Generation**: Automated mock generation for testing

## 🛠 Tech Stack

### Backend
- **Language**: Go 1.21+
- **Database**: PostgreSQL
- **ORM/Query Builder**: SQLC
- **Migration Tool**: Goose
- **AI Model**: Ollama (Llama3)
- **Authentication**: JWT tokens
- **Cron Jobs**: Built-in scheduler
- **Testing**: Go testing framework with Mockgen

### DevOps & Tools
- **Containerization**: Docker
- **Database**: PostgreSQL 12
- **Code Generation**: SQLC
- **Mocking**: Mockgen
- **Security**: Govulncheck
- **Build Automation**: Make

## 📁 Project Structure

Based on your Makefile and configuration, your project likely follows this structure:

```
go-chat-with-db/
├── bin/                        # Compiled binaries
│   └── nlptosql               # Main application binary
├── internal/                   # Private application code
│   └── database/              # Database layer
│       └── mock/              # Generated mocks
│           └── store.go
├── sql/                       # Database files
│   └── schemas/               # Migration files
├── logs/                      # Application logs
│   └── app.log
├── test.env                   # Environment configuration
├── Makefile                   # Build automation
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
├── main.go                    # Application entry point
├── sqlc.yaml                  # SQLC configuration
├── Dockerfile                 # Container definition
└── README.md                  # This file
```

## 🚀 Getting Started

### Prerequisites

- **Go**: Version 1.21 or higher
- **PostgreSQL**: Version 12 or higher
- **Docker**: For database and deployment
- **Make**: For build automation
- **Goose**: For database migrations
- **SQLC**: For SQL code generation

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-chat-with-db
   ```

2. **Set up environment variables**
   ```bash
   cp test.env .env
   # Edit .env with your specific configuration
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Set up PostgreSQL with Docker**
   ```bash
   docker run --name postgres12 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=dbchat -p 5432:5432 -d postgres:12
   ```

5. **Create database**
   ```bash
   make createdb
   ```

6. **Run database migrations**
   ```bash
   make gooseup
   ```

7. **Generate SQL code**
   ```bash
   make sqlc
   ```

8. **Build the application**
   ```bash
   make build
   ```

9. **Run the application**
   ```bash
   make run
   ```

## ⚙️ Configuration

The application uses environment variables defined in `test.env`:

```env
# Server Configuration
PORT=:8080

# Database Configuration
DB_DRIVER=postgres
DB_URL=postgresql://user:password@localhost:5432/dbchat?sslmode=disable
ENVIRONMENT=development

# Cron Job Configuration
CRON_SCHEDULE=@every10m        # Run every 10 minutes
CRON_BATCH_SIZE=100           # Process 100 items per batch

# Authentication Configuration
TOKEN_SECRET_KEY=809bbbb5225c50433e287fc78d22c0e8
TOKEN_SYMMETRIC_KEY=809bbbb5225c50433e287fc78d22c0e8
ACCESS_TOKEN_DURATION=15m     # Token expires in 15 minutes

# AI Model Configuration
API_KEY=809bbbb5225c50433e287fc78d22c0e8
ORG_ID=ollama
PROJECT_ID=001
MODEL=llama3                  # AI model for NLP processing
TEMP=0.7                      # Model temperature for creativity

# Logging Configuration
LOG_PATH=./logs/app.log       # Application log file path
```

## 🛠 Development

### Available Make Commands

#### Build & Run
```bash
make build          # Format code and build binary to bin/nlptosql
make run            # Run the compiled application
```

#### Database Management
```bash
make createdb       # Create PostgreSQL database
make dropdb         # Drop PostgreSQL database
make gooseup        # Run database migrations up
make goosedown      # Rollback database migrations
```

#### Code Generation
```bash
make sqlc           # Generate Go code from SQL
make sqlc-docker    # Generate using Docker (if SQLC not installed locally)
make mock           # Generate mocks for testing
```

#### Testing & Quality
```bash
make test           # Run tests with coverage (short mode)
make runvulnscan    # Run vulnerability scan and generate reports
```

#### Docker
```bash
make buildimage     # Build Docker image as nlqtosql:latest
```

### Development Workflow

1. **Make code changes**
2. **Generate SQL code** (if database queries changed):
   ```bash
   make sqlc
   ```
3. **Run migrations** (if schema changed):
   ```bash
   make gooseup
   ```
4. **Generate mocks** (if interfaces changed):
   ```bash
   make mock
   ```
5. **Run tests**:
   ```bash
   make test
   ```
6. **Build and test**:
   ```bash
   make build
   make run
   ```

## 📚 API Documentation

The application runs on port 8080 by default. Based on the NLP-to-SQL functionality, likely endpoints include:

### Query Processing
- `POST /api/query` - Convert natural language to SQL
- `GET /api/query/history` - Get query history
- `POST /api/execute` - Execute generated SQL

### Authentication
- `POST /api/auth/login` - User authentication
- `POST /api/auth/refresh` - Refresh JWT token

*Note: Specific API documentation should be generated based on your actual endpoints.*

## 🗄️ Database Management

### Migration Commands

```bash
# Create new migration
goose -dir sql/schemas create migration_name sql

# Run migrations
make gooseup

# Rollback migrations
make goosedown

# Check migration status
goose -dir sql/schemas postgres "postgres://user:password@localhost:5432/dbchat?sslmode=disable" status
```

### SQLC Code Generation

SQLC generates type-safe Go code from SQL queries. After modifying SQL files:

```bash
make sqlc          # Generate locally
make sqlc-docker   # Generate using Docker
```

## 🧪 Testing

### Running Tests

```bash
# Run all tests with coverage
make test

# Run specific test package
go test -v ./internal/...

# Run tests with detailed coverage
go test -v -cover ./...
```

### Mock Generation

Generate mocks for testing database interfaces:

```bash
make mock
```

This generates mocks in `internal/database/mock/store.go` for the Store interface.

### Security Scanning

Run vulnerability scans:

```bash
make runvulnscan
```

This generates:
- `vuln.json` - Detailed JSON report
- `vulnsum.txt` - Summary report

## 🚀 Deployment

### Docker Deployment

1. **Build Docker image**
   ```bash
   make buildimage
   ```

2. **Run with Docker**
   ```bash
   docker run -p 8080:8080 --env-file test.env nlqtosql:latest
   ```

### Production Deployment

1. **Set environment variables for production**
2. **Build the application**
   ```bash
   make build
   ```
3. **Run database migrations**
   ```bash
   make gooseup
   ```
4. **Start the application**
   ```bash
   ./bin/nlptosql
   ```

## 🤝 Contributing

1. **Fork the repository**
2. **Create a feature branch**
3. **Make your changes**
4. **Run tests and vulnerability scan**
   ```bash
   make test
   make runvulnscan
   ```
5. **Format and build**
   ```bash
   make build
   ```
6. **Commit your changes**
7. **Push and create a Pull Request**

### Code Quality

- Code is automatically formatted during build (`gofmt -l -s -w .`)
- Run vulnerability scans before committing
- Ensure all tests pass
- Generate mocks when interfaces change
- Update migrations for schema changes

## 📞 Support

- **Author**: Ritankar Saha
- **Email**: ritankar.saha786@gmail.com
- **GitHub**: https://github.com/ritankarsaha

## 🔧 Tools & Dependencies

### Required Tools
- **Go**: 1.21+
- **PostgreSQL**: 12+
- **Goose**: Database migrations
- **SQLC**: SQL code generation
- **Mockgen**: Mock generation
- **Govulncheck**: Vulnerability scanning

### Key Dependencies
Based on the configuration, your project likely uses:
- Database drivers for PostgreSQL
- JWT libraries for authentication
- Cron scheduling libraries
- HTTP routing framework
- Ollama client for AI integration

---

**NLP to SQL Converter - Transforming natural language into structured queries** 🚀

Made with ❤️ by [Ritankar](https://github.com/ritankarsaha)