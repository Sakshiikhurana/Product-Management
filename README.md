# Product Management System

This is a product management system built using Go, Redis, PostgreSQL, RabbitMQ, and Docker. It provides a full-stack solution for managing products, including features for asynchronous image processing, caching, and API endpoints for CRUD operations on product data.

## Features

- **Product Management API**: Allows you to create, read, update, and delete products.
- **Asynchronous Image Processing**: Processes images in the background using RabbitMQ.
- **Caching with Redis**: Caches product data to reduce load on the database.
- **RabbitMQ Integration**: Handles image processing tasks asynchronously.
- **Comprehensive Logging**: Uses `logrus` for structured logging across all services.
- **Unit & Integration Tests**: Includes over 90% code coverage with comprehensive tests.

## Architecture

- **API Layer**: Exposes HTTP endpoints for managing products.
- **Database Layer**: Uses PostgreSQL to store product information.
- **Image Processing Microservice**: A separate service that processes product images.
- **Caching Layer**: Redis is used for caching frequently accessed product data.
- **Message Queue**: RabbitMQ handles asynchronous tasks such as image processing.

## Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/product-management.git
   cd product-management
