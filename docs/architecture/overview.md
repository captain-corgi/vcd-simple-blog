# Architecture Overview

## System Architecture

The VCD Simple Blog application is built using a modern, cloud-native architecture that follows industry best practices for scalability, maintainability, and reliability.

### Key Architectural Patterns

1. **Monorepo Structure**
   - Single repository containing all code for frontend, backend, and infrastructure
   - Clear separation of concerns with dedicated directories
   - Shared code and libraries in packages directory

2. **Clean Architecture**
   - Separation of concerns with distinct layers
   - Domain-centric design with business logic isolated from infrastructure
   - Dependency rule: inner layers don't depend on outer layers

3. **Domain-Driven Design (DDD)**
   - Business logic organized around domain concepts
   - Explicit bounded contexts with clear boundaries
   - Ubiquitous language shared across the codebase

4. **Microservices**
   - Backend organized as independent microservices
   - Each service has its own database and API
   - Services communicate via well-defined interfaces

5. **Micro Frontends**
   - Frontend organized as independent feature applications
   - Shell application that loads and orchestrates features
   - Shared UI components and utilities

## High-Level Component Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                        Client Browsers                           │
└───────────────────────────────┬─────────────────────────────────┘
                                │
┌───────────────────────────────▼─────────────────────────────────┐
│                         CloudFront CDN                           │
└───────────────────────────────┬─────────────────────────────────┘
                                │
┌───────────────────────────────▼─────────────────────────────────┐
│                    Application Load Balancer                     │
└─┬─────────────────────────────┬─────────────────────────────────┘
  │                             │
┌─▼─────────────────┐ ┌─────────▼─────────┐
│  Frontend Shell   │ │    API Gateway    │
│  (React/Vite)     │ │    (Go/Echo)      │
└─┬─────────────────┘ └─┬───────┬─────────┘
  │                     │       │
┌─▼─────────────────┐   │       │
│ Feature Modules   │   │       │
│ - Blog            │   │       │
│ - Auth            │   │       │
│ - Profile         │   │       │
│ - Admin           │   │       │
└───────────────────┘   │       │
                        │       │
┌────────────────────┐ ┌▼───────▼─────────┐ ┌─────────────────────┐
│   Auth Service     │ │   Blog Service    │ │   User Service      │
│   (Go/Echo/GORM)   │ │   (Go/Echo/GORM)  │ │   (Go/Echo/GORM)    │
└─┬──────────────────┘ └─┬─────────────────┘ └─┬───────────────────┘
  │                      │                     │
┌─▼──────────────────┐ ┌─▼─────────────────┐ ┌─▼───────────────────┐
│   Auth Database    │ │   Blog Database    │ │   User Database     │
│   (PostgreSQL)     │ │   (PostgreSQL)     │ │   (PostgreSQL)      │
└────────────────────┘ └───────────────────┘ └─────────────────────┘
```

## Technology Stack

### Frontend
- React 18+ with TypeScript 5.0+
- Vite for build tooling
- Tailwind CSS and Shadcn UI
- TanStack Query and TanStack Table
- React Hook Form
- Storybook for component development
- Vitest and Playwright for testing

### Backend
- Go 1.21+ with Echo framework
- Clean Architecture with explicit layers
- Domain-Driven Design principles
- gRPC for service-to-service communication
- GORM with PostgreSQL
- JWT and OAuth2 for authentication

### Infrastructure
- Docker and Docker Compose for local development
- Terraform for infrastructure as code
- AWS as the deployment platform
- GitHub Actions for CI/CD
- Prometheus, Grafana, and Datadog for monitoring
