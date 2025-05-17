# Getting Started

This guide will help you set up your development environment for the VCD Simple Blog application.

## Prerequisites

Before you begin, make sure you have the following installed:

- **Node.js** (v18 or later)
- **pnpm** (v8 or later)
- **Go** (v1.21 or later)
- **Docker** and **Docker Compose**
- **Git**

## Clone the Repository

```bash
git clone https://github.com/yourusername/vcd-simple-blog.git
cd vcd-simple-blog
```

## Frontend Setup

1. Install dependencies:

```bash
pnpm install
```

2. Start the frontend development server:

```bash
pnpm --filter "./apps/frontend/shell" dev
```

This will start the shell application on http://localhost:3000.

## Backend Setup

1. Start the backend services using Docker Compose:

```bash
docker-compose up -d
```

This will start:
- PostgreSQL database
- API Gateway
- Auth Service
- Blog Service
- User Service
- Adminer (database management tool)

2. If you want to run a specific backend service locally (outside Docker):

```bash
cd apps/backend/blog-service
go run main.go
```

## Database Setup

The Docker Compose configuration automatically sets up the required databases. You can access them using:

- **Host**: localhost
- **Port**: 5432
- **Username**: postgres
- **Password**: postgres
- **Databases**: auth_db, blog_db, user_db

You can also use Adminer to manage the databases at http://localhost:8888.

## Environment Variables

The application uses environment variables for configuration. For local development, these are set in the Docker Compose file.

If you need to override any variables, you can create a `.env.local` file in the root directory.

## Development Tools

### Frontend

- **Storybook**: Run `pnpm --filter "./apps/frontend/shell" storybook` to start Storybook on http://localhost:6006
- **Tests**: Run `pnpm test` to run all frontend tests
- **Lint**: Run `pnpm lint` to lint all frontend code

### Backend

- **Tests**: Run `cd apps/backend && go test ./...` to run all backend tests
- **API Documentation**: Available at http://localhost:8080/swagger/index.html when running the API Gateway

## Common Tasks

### Creating a New Frontend Feature

1. Create a new directory in `apps/frontend` for your feature
2. Copy the basic structure from an existing feature
3. Update the package.json with the new name
4. Add the feature to the shell application in `apps/frontend/shell/src/App.tsx`

### Creating a New Backend Service

1. Create a new directory in `apps/backend` for your service
2. Create a new go.mod file
3. Implement the service following the Clean Architecture pattern
4. Add the service to the Docker Compose file
5. Add the service to the API Gateway configuration

## Troubleshooting

### Frontend Issues

- **Module not found errors**: Make sure you've run `pnpm install` and that the package is correctly listed in the dependencies
- **Hot reload not working**: Check that you're running the correct dev command and that your file is being watched

### Backend Issues

- **Connection refused errors**: Make sure the required services are running in Docker
- **Database errors**: Check that the database is running and that the connection details are correct

### Docker Issues

- **Port conflicts**: Make sure no other services are using the same ports
- **Container not starting**: Check the logs with `docker-compose logs <service-name>`

## Getting Help

If you're stuck, you can:
- Check the documentation in the `docs` directory
- Ask in the team chat
- Create an issue on GitHub
