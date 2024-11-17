
# Tender Management System

This project is a **Tender Management System** implemented in Go (Golang). It provides APIs for managing tenders, bids, and notifications. The system is designed to handle user and contractor interactions seamlessly.

---

## Features

### Tender Management
- **Create Tender**: Create new tenders.
- **Update Tender**: Update existing tenders.
- **Delete Tender**: Remove tenders.
- **List Tenders**: Retrieve a list of all tenders.
- **Award Tender**: Mark a tender as awarded.
- **List User Tenders**: Retrieve tenders associated with a specific user.

### Bid Management
- **Submit Bid**: Submit a bid for a tender.
- **List Bids**: Retrieve all bids.
- **Get All Bids by Tender ID**: Retrieve bids submitted for a specific tender.
- **List Contractor Bids**: Retrieve bids submitted by a specific contractor.
- **Get Bids by Tender ID**: Retrieve detailed bid information for a tender.

### Notification Management
- **Create Notification**: Add new notifications for system events.

---

## Project Structure

```
tender/
├── internal/
│   ├── pkg/
│   │   ├── genprotos/       # Generated protobuf files
│   │   ├── storage/         # Interfaces for database operations
│   ├── http/
│   │   ├── handlers/        # API handlers
│   │   ├── router.go        # HTTP router setup
├── migrations/              # Database migration files
├── Dockerfile               # Docker setup for the service
├── docker-compose.yml       # Docker Compose setup
└── README.md                # Project documentation
```

---

## Getting Started

### Prerequisites
Make sure you have the following installed:
- [Go 1.20+](https://golang.org/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Protobuf Compiler](https://grpc.io/docs/protoc-installation/)

---
