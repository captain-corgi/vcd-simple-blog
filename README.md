# VCD Simple Blog - Modern Web Application Monorepo

This repository contains a comprehensive monorepo for a modern web application following Clean Architecture and Domain-Driven Design principles.

## Architecture Overview

This project implements:
- Monorepo structure with clearly separated concerns
- Clean Architecture principles throughout the codebase
- Domain-Driven Design (DDD) for business logic organization
- Microservices architecture for the backend
- Micro Frontend architecture for the frontend

## Technical Stack

### Frontend
- React 18+ with TypeScript 5.0+
- Micro Frontend architecture with shell and feature applications
- pnpm workspaces for package management
- Vite as the build tool
- Tailwind CSS and Shadcn UI
- TanStack Query and TanStack Table
- React Hook Form
- Mock Service Worker (MSW)
- Storybook
- Vitest for testing
- Playwright for end-to-end testing

### Backend
- Go 1.21+
- Echo framework for REST APIs
- Clean Architecture with explicit layers
- Domain-Driven Design with bounded contexts
- Go workspace for monorepo management
- gRPC with Protocol Buffers
- OpenAPI 3.0+ documentation
- GORM with PostgreSQL
- JWT and OAuth2 authentication
- Comprehensive testing

### Infrastructure
- Docker and Docker Compose
- Terraform for infrastructure as code
- GitHub Actions for CI/CD
- AWS deployment
- Multi-environment support
- Security scanning
- Monitoring with Prometheus, Grafana, and Datadog

## Getting Started

### Prerequisites
- pnpm 8.0+
- Go 1.21+
- Docker and Docker Compose
- Node.js 18+

### Development Setup
```bash
# Clone the repository
git clone https://github.com/yourusername/vcd-simple-blog.git
cd vcd-simple-blog

# Install dependencies
pnpm install

# Start development environment
docker-compose up -d
pnpm dev
```

## Documentation
See the `/docs` directory for detailed documentation on architecture, development workflows, and deployment procedures.

## License
MIT
