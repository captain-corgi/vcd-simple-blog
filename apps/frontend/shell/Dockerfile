# Build stage
FROM node:18-alpine AS build

# Install pnpm
RUN npm install -g pnpm

# Set working directory
WORKDIR /app

# Copy root package.json and pnpm-workspace.yaml
COPY package.json pnpm-workspace.yaml ./

# Copy all package.json files
COPY apps/frontend/shell/package.json ./apps/frontend/shell/
COPY apps/frontend/blog/package.json ./apps/frontend/blog/
COPY apps/frontend/auth/package.json ./apps/frontend/auth/
COPY packages/ui/package.json ./packages/ui/
COPY packages/utils/package.json ./packages/utils/

# Install dependencies
RUN pnpm install --frozen-lockfile

# Copy the rest of the application
COPY . .

# Build the application
RUN pnpm --filter "@vcd-simple-blog/shell" build

# Production stage
FROM nginx:alpine

# Copy the build output
COPY --from=build /app/apps/frontend/shell/dist /usr/share/nginx/html

# Copy nginx configuration
COPY apps/frontend/shell/nginx.conf /etc/nginx/conf.d/default.conf

# Expose port
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
