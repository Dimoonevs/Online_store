# Online Shop Website

This is an online shop website where users can authenticate using the credentials username: admin and password: admin123. Additionally, users can create, retrieve, update, and delete sellers, as well as create, retrieve, update, and delete products.

## Features

- Authentication with username: admin and password: admin123
- CRUD operations for sellers
- CRUD operations for products
- Dockerized application with Dockerfile and Docker Compose
- RESTful API using Go
- PostgreSQL database

## Getting Started

To run the application, make sure you have Docker installed. Then, you can start the application using the following command:

```bash
make up_build
```

This command will build and start the application in Docker containers.

## API Endpoints

Below are the main API endpoints for this application:

- **GET /seller/all**: Retrieve sellers.
- **GET /seller**: Retrieve seller by ID.
  - Request Header Example:
  ```json
  {"id": 1}
  ```
- **POST /seller**: Create a new seller.
  - Request Body Example:
  ```json
  { "name": "John", "phone": "0966962985" }
  ```
- **PUT /seller**: Update seller by ID.
  - Request Header Example:
  ```json
  {"id": 1}
  ```
  - Request Body Example:
  ```json
  { "name": "John Doe", "phone": "+380966962985" }
  ```
- **DELETE /seller**: Delete seller by ID.
  - Request Header Example:
  ```json
  {"id": 1}
  ```
- **GET /product**: Retrieve product by ID.
  - Request Header Example:
  ```json
  {"id": 1}
  ``` 
- **GET /product/seller**: Retrieve product by ID seller.
  - Request Header ID seller Example:
   ```json
    {"id": 1}
   ``` 
- **POST /product**: Create a new product.
  - Request Header ID seller Example:
  ```json
   {"id": 1}
  ```
  - Request Body Example:
  ```json
   { "name": "Laptop", "price": 999.99 }
  ```
- **PUT /product**: Update product by ID.
  - Request Header ID seller Example:
  ```json
  {"id": 1}
  ```
  - Request Body Example:
  ```json
  { "name": "Desktop", "price": 899.99 }
  ```
- **DELETE /product**: Delete product by ID.
  - Request Header Example:
  ```json
  {"id": 1}
  ```
- **GET /product/all**: Get all product.

