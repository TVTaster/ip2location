# **IP2Location Service**

## Overview

`IP2Location Service` is a high-performance API service written in Go. It is designed with modularity, scalability, and
clean architecture principles in mind. The project resolves [main purpose, e.g., "IP-to-location mapping"] efficiently
by using structured components and a layered design.
This service enables functionalities such as:

- Fast IP lookups using a CSV-based IP database.
- Rate-limiting controls for API endpoints.
- Easy integration into both local and production environments via Docker.

## **Key Features**

- **IP Geolocation**: Maps IP addresses to geographical information (e.g., country, region, city).
- **Rate-limiting**: Configurable API rate limits for throttling requests and preventing abuse.
- **Containerization**: Fully-ready for deployment with Docker.
- **Clean Architecture**: Decoupled components for maintainability and for implementing future features without breaking
  core functionality.
- **Extensibility**: Easily extendable repository and service layers to integrate with other databases or APIs.
- **Robust Testing Suite**: Comprehensive test coverage for critical components to ensure system reliability.

## **Project Structure**

This repository is structured to follow standard clean architecture principles. Here's a breakdown of the major
directories and their roles:

``` 
├── cmd/
│   └── server/           # Entry point for the application/server
│       └── main.go       # Main application logic to boot up the server
├── internal/
│   ├── config/           # configuration (environments variables)
│   ├── mocks/            # Mocks entities (e.g., Datastore, Rate Limiter, etc.)
│   ├── models/           # Entity definitions (e.g., Datastore, Rate Limiter, etc.)
│   ├── repositories/     # Interfaces and database implementations
│   │   └── csv_repository/     # Example: CSV datastore implementation
│   ├── services/         # Business logic and service layer
├── go.mod                # Module dependencies configuration
└── README.md             # Documentation for the project
```

Each directory encapsulates specific responsibilities to ensure modularity and maintainability.

## **Installation and Setup**

### **Prerequisites**

To run this project, ensure the following prerequisites are met:

- **Go**: >= 1.19
- **Docker** (optional, for containerized deployment)

### **Steps to Run Locally**

1. Clone the repository:

``` bash
   git clone [repository_url]
   cd [project_directory]
```

1. Install dependencies:

``` bash
   go mod tidy
```

1. Set up the environment variables: Create a `.env` file in the root directory with the following:

example:
``` 
   RATE_LIMIT_PER_SEC=5
   CSV_DATA_SOURCE_PATH=./data/ip_database.csv
   ..
```
Or use the preconfigured in the repo

1. Run the application:

``` bash
   go run ./cmd/server/main.go
```

### **Run with Docker**

1. Build the Docker image:

``` bash
   docker build -t ip2location-service .
```

1. Run the service:

``` bash
   docker run -e RATE_LIMIT_PER_SEC=5 -e DB_TYPE=csv CSV_DATA_SOURCE_PATH=./data/ip_database.csv -p 8080:8080 ip2location-service
```

1. Access the API: Open `http://localhost:8080` in your browser or use a tool like `curl`/Postman to interact with the
   endpoints.

## **Usage**

This section briefly outlines the API usage.

### **Endpoints**

1. **GET /location/:ip**
    - **Description**: Retrieves the geolocation data for a given IP address.
    - **Example Request**:

``` bash
     curl http://localhost:8080/v1/find-country?ip=8.8.8.8
```

- **Response**:

``` json
     {
       "country": "United States",
       "city": "Mountain View"
     }
```

1. **GET /health**
    - **Description**: Health check endpoint to verify the service is running.
    - **Response**:

``` text
     Hello, this is the IP2Location server!
```

### **Rate Limiting**

- The rate limit for the application can be configured using the `RATE_LIMIT_PER_SEC` environment variable.
- Example:

``` bash
  curl http://localhost:8080/v1/find-country?ip=8.8.8.8
  # If exceeded:
  {
     "error": "Too many requests, retry later."
  }
```

## **Technologies Used**

- **Programming Language**: Go (Golang)
- **Database**: CSV (or any supported DB with repository integration)
- **Frameworks**:
    - HTTP Routing: `net/http`
      - Dependency Injection: Lightweight custom implementation

- **Containerization**: Docker
- **Environment Management**: `godotenv`

## **Testing**

The project includes unit tests and integration tests for key components.

### **Run Tests:**

1. Install `go test` dependencies:

``` bash
   go install
```

1. Run the test suite:

``` bash
   go test ./... -v
```

The testing suite ensures the stability of:

- Services (business logic)
- Repository interactions
- Handlers and HTTP endpoints

I need to add design discussion about the packages and project structure