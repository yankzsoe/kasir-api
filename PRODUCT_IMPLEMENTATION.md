# Product Controller, Service, and Repository Implementation

## Overview
This document outlines the complete product feature implementation that mirrors the category controller pattern.

## Files Created

### 1. Controller: [controllers/products.controller.go](controllers/products.controller.go)
- **CreateProduct**: Handle POST request to create a new product
- **GetAllProducts**: Handle GET request to retrieve all products
- **GetProductByID**: Handle GET request to retrieve a specific product by ID
- **UpdateProduct**: Handle PUT request to update an existing product
- **DeleteProduct**: Handle DELETE request to soft delete a product

Features:
- Request validation using DTOs
- Error handling with proper HTTP status codes
- DTO to model and model to DTO conversion
- Swagger documentation comments for all endpoints

### 2. Service: [services/product.service.go](services/product.service.go)
- **CreateProduct**: Business logic for creating products with validation
- **GetAllProducts**: Retrieve all products from repository
- **GetProductByID**: Retrieve a specific product by ID
- **UpdateProduct**: Update product with validation
- **DeleteProduct**: Soft delete a product

Validations included:
- Product name is required
- Product price must be greater than 0 (on create)
- Product price cannot be negative (on update)
- Product stock cannot be negative

### 3. Repository: [repositories/product.repository.go](repositories/product.repository.go)
- **Create**: Insert a new product into the database
- **FindByID**: Find a product by ID
- **FindAll**: Retrieve all products
- **Update**: Update an existing product
- **Delete**: Soft delete a product (using gorm DeletedAt)

Implements the `IProductRepository` interface for dependency injection.

### 4. Data Transfer Objects: [dtos/product.go](dtos/product.go)
- **ProductCreateRequest**: DTO for creating products
- **ProductUpdateRequest**: DTO for updating products
- **ProductUriRequest**: DTO for URI parameters
- **ProductResponse**: Response DTO for products
- **CreateProductSuccessResponse**: Success response for create
- **GetProductsSuccessResponse**: Success response for list
- **GetProductSuccessResponse**: Success response for single product
- **UpdateProductSuccessResponse**: Success response for update
- **DeleteProductSuccessResponse**: Success response for delete

## Updated Files

### 1. [repositories/repository.go](repositories/repository.go)
- Added `Product IProductRepository` field to Repository struct
- Initialize ProductRepository in NewRepository

### 2. [main.go](main.go)
- Create ProductService instance from repository
- Set ProductService in controller via SetProductService

### 3. [routers/routes.go](routers/routes.go)
- Added setupProductRoutes function to configure product routes
- Call setupProductRoutes in SetupRoutes function
- Registered endpoints:
  - `POST /api/v1/products` - Create product
  - `GET /api/v1/products` - Get all products
  - `GET /api/v1/products/:id` - Get product by ID
  - `PUT /api/v1/products/:id` - Update product
  - `DELETE /api/v1/products/:id` - Delete product

## Architecture Pattern

The implementation follows the same clean architecture pattern as the Category feature:

```
Request → Controller → Service → Repository → Database
                ↓
              DTOs
```

- **Controller**: Handles HTTP requests/responses and validation
- **Service**: Contains business logic and validation rules
- **Repository**: Abstracts database operations using GORM
- **DTOs**: Separate request and response models from domain models

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/products` | Create a new product |
| GET | `/api/v1/products` | Get all products |
| GET | `/api/v1/products/:id` | Get product by ID |
| PUT | `/api/v1/products/:id` | Update product |
| DELETE | `/api/v1/products/:id` | Delete product |

## Request/Response Examples

### Create Product
**Request:**
```json
{
  "nama": "Laptop",
  "harga": 10000000,
  "stok": 5,
  "is_active": true,
  "category_id": 1
}
```

**Response (201):**
```json
{
  "message": "Product created successfully",
  "data": {
    "id": 1,
    "nama": "Laptop",
    "harga": 10000000,
    "stok": 5,
    "is_active": true,
    "category_id": 1,
    "created_at": "2026-01-31T12:00:00Z",
    "updated_at": "2026-01-31T12:00:00Z"
  }
}
```

## Testing

To test the implementation, you can use:
- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- **Postman**: Import the Swagger JSON specification
- **cURL**: Direct HTTP requests to the endpoints

## Notes

- The product feature uses the same soft delete pattern as categories
- Product model fields follow the existing Produk struct naming convention
- All validations are consistent with the category implementation
- Swagger documentation is included for all endpoints
