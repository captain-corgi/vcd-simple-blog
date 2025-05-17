# Microservices Architecture

The VCD Simple Blog application implements a microservices architecture for the backend to enable independent development, deployment, and scaling of services.

## Overview

Our microservices architecture divides the backend into small, focused services that each handle a specific business domain. Each service has its own database and API, and communicates with other services through well-defined interfaces.

## Service Boundaries

We have defined the following microservices:

### 1. API Gateway

The API Gateway serves as the entry point for all client requests. It:
- Routes requests to the appropriate service
- Handles authentication and authorization
- Implements rate limiting and request validation
- Provides a unified API for the frontend

**Technologies**: Go, Echo framework

### 2. Auth Service

The Auth Service handles user authentication and authorization. It:
- Manages user registration and login
- Issues and validates JWT tokens
- Handles password reset and account verification
- Manages user sessions and logout

**Technologies**: Go, Echo framework, GORM, PostgreSQL

### 3. Blog Service

The Blog Service manages all blog-related functionality. It:
- Creates, reads, updates, and deletes blog posts
- Manages blog categories and tags
- Handles blog publishing workflow
- Manages comments and reactions

**Technologies**: Go, Echo framework, GORM, PostgreSQL

### 4. User Service

The User Service manages user profiles and preferences. It:
- Stores and retrieves user profile information
- Manages user preferences and settings
- Handles user relationships (following, etc.)
- Manages user activity and notifications

**Technologies**: Go, Echo framework, GORM, PostgreSQL

## Communication Patterns

### 1. Synchronous Communication (REST)

Services communicate with each other synchronously using REST APIs for:
- Read operations that need immediate responses
- Simple CRUD operations
- User-initiated actions that need immediate feedback

Example:
```go
// API Gateway calling Blog Service
resp, err := http.Get(h.blogServiceURL + "/blogs/" + id)
if err != nil {
    return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to blog service")
}
```

### 2. Synchronous Communication (gRPC)

For service-to-service communication that requires high performance, we use gRPC:
- Complex operations that involve multiple steps
- Operations that need to transfer large amounts of data
- Internal service-to-service communication

Example:
```go
// User Service calling Auth Service via gRPC
client := pb.NewAuthServiceClient(conn)
resp, err := client.ValidateToken(ctx, &pb.ValidateTokenRequest{
    Token: token,
})
```

### 3. Asynchronous Communication (Events)

For operations that don't need immediate responses, we use event-based communication:
- Background processing
- Notifications
- Data synchronization between services
- Operations that can be retried

## Database Per Service

Each microservice has its own database to ensure:
- Data isolation
- Independent scaling
- Freedom to choose the right database for each service's needs

Services never access another service's database directly. All data access is through the service's API.

## Service Discovery and Configuration

We use environment variables and DNS for service discovery in different environments:
- Local development: Docker Compose with service names
- Production: Kubernetes Service discovery or AWS ECS Service Discovery

Example:
```go
authServiceURL := os.Getenv("AUTH_SERVICE_URL")
if authServiceURL == "" {
    authServiceURL = "http://auth-service:8081"
}
```

## API Documentation

Each service maintains its own OpenAPI documentation to describe its API:
- Endpoints
- Request/response formats
- Authentication requirements
- Error responses

The API Gateway aggregates these into a unified API documentation for frontend developers.

## Benefits of Microservices Architecture

1. **Independent Development**: Teams can work on different services without affecting each other.
2. **Independent Deployment**: Services can be deployed independently, reducing risk and enabling continuous delivery.
3. **Independent Scaling**: Services can be scaled based on their specific resource needs.
4. **Technology Flexibility**: Each service can use the best technology for its specific requirements.
5. **Fault Isolation**: Failures in one service don't necessarily affect others.

## Challenges and Mitigations

1. **Distributed System Complexity**: We use proper logging, monitoring, and tracing to manage complexity.
2. **Data Consistency**: We implement eventual consistency patterns where appropriate.
3. **Service Coordination**: We use API Gateway for client requests and direct service-to-service communication for internal operations.
4. **Testing**: We implement comprehensive integration tests to verify service interactions.
