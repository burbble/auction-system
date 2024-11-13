# Auction System

## Technology Stack

- Go 1.22.3
- gRPC/Protocol Buffers
- PostgreSQL 14
- Docker & Docker Compose
- gRPC-Gateway for REST API

## Architecture

The project is built using Clean Architecture principles:

### Application Layers

- **API** (`api/proto/`) - Protocol Buffers definitions
- **Interfaces** (`internal/interfaces/`) - gRPC and REST handlers
- **Application** (`internal/application/`) - Business logic and use cases
- **Domain** (`internal/domain/`) - Business entities and interfaces
- **Infrastructure** (`internal/infrastructure/`) - Repository implementations

### Core Entities

- Users - System users
- Lots - Auction items
- Auctions - Auction events
- Bids - User bids

## Installation

### Prerequisites

- Docker and Docker Compose
- Go 1.22.3+
- Make
- Protocol Buffers compiler

### Installation Steps

1. Install dependencies for proto file generation:

```bash
make install-proto-deps
```
2. Generate code from proto files:

```bash
make gen-proto
```

3. Start the application:

```bash
make start
```

## Make Commands

```bash
make start        # Start the entire application
make up           # Start containers
make down         # Stop containers
make build        # Build the application
make migrate      # Apply migrations
make seed         # Load test data
make logs         # View logs for all services
make app-logs     # Application logs only
make test         # Run tests
make test-verbose # Run tests with verbose output
make db-logs      # View database logs
make reset-db     # Reset the database
make restart-app  # Restart the application
```

API Endpoints
User Service
1. Create a user

POST /api/v1/users

Request body:
```bash
{
    "username": "ivan_petrov",
    "email": "ivan@example.com"
}
```

2. Get a user

GET /api/v1/users/{id}

3. Update a user

PUT /api/v1/users/{id}

Request body:
```bash
{
    "username": "new_username",
    "email": "new_email@example.com"
}
```

4. Delete a user

DELETE /api/v1/users/{id}

5. List users

GET /api/v1/users?page_size=10&page_number=1

6. Update balance

POST /api/v1/users/{user_id}/balance

Request body:
```bash
{
    "amount": 1000.00
}
```

Lot Service
1. Create a lot

POST /api/v1/lots

Request body:
```bash
{
    "title": "Antique Clocks",
    "description": "Swiss clocks from the 19th century",
    "start_price": 5000.00,
    "creator_id": 1
}
```

2. Get a lot

GET /api/v1/lots/{id}

3. Update a lot

PUT /api/v1/lots/{id}

Request body:
```bash
{
    "title": "Updated Title",
    "description": "Updated Description",
    "start_price": 6000.00
}
```
4. Delete a lot

DELETE /api/v1/lots/{id}

5. List lots

GET /api/v1/lots?page_size=10&page_number=1

Auction Service
1. Create an auction

POST /api/v1/auctions

Request body:
```bash
{
    "lot_id": 1,
    "start_price": 5000.00,
    "min_step": 100.00,
    "start_time": "2024-03-20T10:00:00Z",
    "end_time": "2024-03-25T10:00:00Z"
}
```

2. Get an auction

GET /api/v1/auctions/{id}

3. Update an auction
PUT /api/v1/auctions/{id}

Request body:
```bash
{
    "start_price": 5500.00,
    "min_step": 150.00,
    "start_time": "2024-03-21T10:00:00Z",
    "end_time": "2024-03-26T10:00:00Z",
    "status": "ACTIVE"
}
```

4. Delete an auction

GET /api/v1/auctions?page_size=10&page_number=1&status=ACTIVE

Bid Service
1. Place a bid

POST /api/v1/bids

Request body:
```bash
{
    "auction_id": 1,
    "user_id": 1,
    "amount": 5100.00
}
```
2. Get a bid

GET /api/v1/bids/{id}

3. List bids

GET /api/v1/bids?auction_id=1&page_size=10&page_number=1
