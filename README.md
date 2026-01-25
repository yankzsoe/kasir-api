# Kasir API

A complete Point of Sale (POS) API built with Go, Gin, and GORM.

## Features

- SQLite database with GORM ORM
- RESTful API with Gin framework
- Category management (CRUD operations)
- Swagger API documentation
- Environment-based configuration
- Clean architecture with service layer

## Project Structure

```
kasir-api/
├── configs/           # Configuration and database setup
│   ├── config.go      # Environment configuration loader
│   └── database.go    # Database initialization
├── controllers/       # HTTP request handlers
│   └── categories.controller.go
├── models/           # Data models
│   └── category.go
├── services/         # Business logic layer
│   └── category.service.go
├── routers/          # Route definitions
│   └── routes.go
├── docs/             # Swagger documentation
├── .env              # Environment variables
├── .env.example      # Example environment file
├── main.go           # Application entry point
└── go.mod            # Go module dependencies
```

## Prerequisites

- Go 1.16 or higher
- SQLite3

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd kasir-api
```

2. Install dependencies:
```bash
go mod download
```

3. Copy `.env.example` to `.env` and configure:
```bash
cp .env.example .env
```

4. Generate Swagger documentation:
```bash
swag init
```

## Configuration

All configuration is managed via environment variables in the `.env` file:

### Database Configuration
- `DB_DRIVER`: Database driver (default: sqlite)
- `DB_PATH`: Path to SQLite database file (default: kasir.db)

### Server Configuration
- `PORT`: Server port (default: 8080)
- `GIN_MODE`: Gin mode - debug, release, test (default: debug)

### API Configuration
- `API_VERSION`: API version (default: v1)
- `API_TITLE`: API title (default: Kasir API)
- `API_DESCRIPTION`: API description

### Swagger Configuration
- `SWAGGER_HOST`: Swagger documentation host
- `SWAGGER_BASEPATH`: API base path
- `SWAGGER_SCHEMES`: HTTP schemes (http, https)

## Running the Application

```bash
go run main.go
```

The API will start on `http://localhost:8080`

## API Documentation

Once the server is running, visit:
```
http://localhost:8080/swagger/index.html
```

## API Endpoints

### Health Check
- `GET /api/v1/health` - Check API health status

### Categories
- `GET /api/v1/categories` - Get all categories
- `POST /api/v1/categories` - Create new category
- `GET /api/v1/categories/:id` - Get category by ID
- `PUT /api/v1/categories/:id` - Update category
- `DELETE /api/v1/categories/:id` - Delete category

## Example Requests

### Create Category
```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electronics",
    "description": "Electronic items",
    "is_active": true
  }'
```

### Get All Categories
```bash
curl http://localhost:8080/api/v1/categories
```

### Get Category by ID
```bash
curl http://localhost:8080/api/v1/categories/1
```

### Update Category
```bash
curl -X PUT http://localhost:8080/api/v1/categories/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electronics Updated",
    "description": "Updated description"
  }'
```

### Delete Category
```bash
curl -X DELETE http://localhost:8080/api/v1/categories/1
```

## Development

### Generate Swagger Docs
After making changes to endpoints, regenerate Swagger documentation:
```bash
swag init
```

### Database Migration
The application automatically migrates the database schema on startup.

## License

MIT License

## Support

For issues and questions, please create an issue in the repository.
