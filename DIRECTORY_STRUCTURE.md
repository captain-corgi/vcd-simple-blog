# VCD Simple Blog - Directory Structure

```
vcd-simple-blog/
│
├── .github/                                # GitHub configuration
│   └── workflows/                          # GitHub Actions workflows
│       ├── ci.yml                          # CI workflow
│       └── cd.yml                          # CD workflow
│
├── apps/                                   # Application code
│   ├── frontend/                           # Frontend applications
│   │   ├── shell/                          # Shell application (host)
│   │   │   ├── public/                     # Public assets
│   │   │   ├── src/                        # Source code
│   │   │   │   ├── components/             # UI components
│   │   │   │   │   ├── layout/             # Layout components
│   │   │   │   │   └── ui/                 # UI components
│   │   │   │   ├── hooks/                  # Custom React hooks
│   │   │   │   ├── lib/                    # Utilities and helpers
│   │   │   │   ├── modules/                # Module integration
│   │   │   │   ├── pages/                  # Page components
│   │   │   │   ├── App.tsx                 # Main App component
│   │   │   │   ├── index.css               # Global styles
│   │   │   │   └── main.tsx                # Entry point
│   │   │   ├── .storybook/                 # Storybook configuration
│   │   │   ├── Dockerfile                  # Production Dockerfile
│   │   │   ├── Dockerfile.dev              # Development Dockerfile
│   │   │   ├── package.json                # Package configuration
│   │   │   ├── tailwind.config.js          # Tailwind CSS configuration
│   │   │   ├── tsconfig.json               # TypeScript configuration
│   │   │   └── vite.config.ts              # Vite configuration
│   │   │
│   │   ├── blog/                           # Blog feature module
│   │   │   ├── src/                        # Source code
│   │   │   │   ├── components/             # UI components
│   │   │   │   ├── hooks/                  # Custom React hooks
│   │   │   │   ├── pages/                  # Page components
│   │   │   │   ├── services/               # API services
│   │   │   │   ├── types/                  # TypeScript types
│   │   │   │   ├── BlogModule.tsx          # Main module component
│   │   │   │   └── index.ts                # Entry point
│   │   │   ├── package.json                # Package configuration
│   │   │   ├── tsconfig.json               # TypeScript configuration
│   │   │   └── vite.config.ts              # Vite configuration
│   │   │
│   │   ├── auth/                           # Auth feature module
│   │   │   └── ...                         # Similar structure to blog
│   │   │
│   │   ├── profile/                        # Profile feature module
│   │   │   └── ...                         # Similar structure to blog
│   │   │
│   │   └── admin/                          # Admin feature module
│   │       └── ...                         # Similar structure to blog
│   │
│   └── backend/                            # Backend services
│       ├── api-gateway/                    # API Gateway service
│       │   ├── config/                     # Configuration
│       │   ├── interfaces/                 # Interface adapters
│       │   │   ├── http/                   # HTTP handlers
│       │   │   │   ├── handlers/           # Request handlers
│       │   │   │   ├── middleware/         # HTTP middleware
│       │   │   │   └── routes.go           # Route definitions
│       │   │   └── grpc/                   # gRPC adapters
│       │   ├── Dockerfile                  # Production Dockerfile
│       │   ├── Dockerfile.dev              # Development Dockerfile
│       │   ├── go.mod                      # Go module file
│       │   ├── go.sum                      # Go dependencies
│       │   └── main.go                     # Entry point
│       │
│       ├── auth-service/                   # Authentication service
│       │   ├── config/                     # Configuration
│       │   ├── domain/                     # Domain layer
│       │   │   ├── entity/                 # Domain entities
│       │   │   ├── repository/             # Repository interfaces
│       │   │   ├── service/                # Domain services
│       │   │   └── valueobject/            # Value objects
│       │   ├── infrastructure/             # Infrastructure layer
│       │   │   ├── database/               # Database connection
│       │   │   ├── repository/             # Repository implementations
│       │   │   └── service/                # External service clients
│       │   ├── interfaces/                 # Interface adapters
│       │   │   ├── http/                   # HTTP handlers
│       │   │   └── grpc/                   # gRPC handlers
│       │   ├── usecases/                   # Application layer
│       │   ├── Dockerfile                  # Production Dockerfile
│       │   ├── Dockerfile.dev              # Development Dockerfile
│       │   ├── go.mod                      # Go module file
│       │   ├── go.sum                      # Go dependencies
│       │   └── main.go                     # Entry point
│       │
│       ├── blog-service/                   # Blog service
│       │   └── ...                         # Similar structure to auth-service
│       │
│       └── user-service/                   # User service
│           └── ...                         # Similar structure to auth-service
│
├── packages/                               # Shared packages
│   ├── ui/                                 # Shared UI components
│   │   ├── src/                            # Source code
│   │   │   ├── components/                 # UI components
│   │   │   │   └── ui/                     # UI components
│   │   │   ├── lib/                        # Utilities
│   │   │   └── index.ts                    # Entry point
│   │   ├── package.json                    # Package configuration
│   │   └── tsconfig.json                   # TypeScript configuration
│   │
│   ├── utils/                              # Shared utilities
│   │   ├── src/                            # Source code
│   │   │   ├── api/                        # API utilities
│   │   │   ├── formatting/                 # Formatting utilities
│   │   │   ├── validation/                 # Validation utilities
│   │   │   └── index.ts                    # Entry point
│   │   ├── package.json                    # Package configuration
│   │   └── tsconfig.json                   # TypeScript configuration
│   │
│   └── go/                                 # Shared Go packages
│       └── common/                         # Common Go utilities
│           ├── auth/                       # Authentication utilities
│           ├── database/                   # Database utilities
│           ├── logger/                     # Logging utilities
│           ├── validator/                  # Validation utilities
│           ├── go.mod                      # Go module file
│           └── go.sum                      # Go dependencies
│
├── infrastructure/                         # Infrastructure code
│   ├── docker/                             # Docker configuration
│   │   ├── postgres/                       # PostgreSQL configuration
│   │   │   └── init-multiple-dbs.sh        # Database initialization script
│   │   └── prometheus/                     # Prometheus configuration
│   │       └── prometheus.yml              # Prometheus configuration
│   │
│   └── terraform/                          # Terraform configuration
│       ├── environments/                   # Environment-specific variables
│       │   ├── dev.tfvars                  # Development variables
│       │   ├── sit.tfvars                  # SIT variables
│       │   ├── uat.tfvars                  # UAT variables
│       │   ├── nft.tfvars                  # NFT variables
│       │   ├── stg.tfvars                  # Staging variables
│       │   └── prd.tfvars                  # Production variables
│       ├── modules/                        # Terraform modules
│       │   ├── vpc/                        # VPC module
│       │   ├── rds/                        # RDS module
│       │   ├── ecr/                        # ECR module
│       │   ├── ecs/                        # ECS module
│       │   ├── cloudfront/                 # CloudFront module
│       │   └── monitoring/                 # Monitoring module
│       ├── main.tf                         # Main Terraform configuration
│       ├── variables.tf                    # Variable definitions
│       └── outputs.tf                      # Output definitions
│
├── docs/                                   # Documentation
│   ├── architecture/                       # Architecture documentation
│   │   ├── overview.md                     # Architecture overview
│   │   ├── clean-architecture.md           # Clean Architecture documentation
│   │   ├── domain-driven-design.md         # DDD documentation
│   │   ├── microservices.md                # Microservices documentation
│   │   └── micro-frontends.md              # Micro Frontend documentation
│   ├── development/                        # Development documentation
│   │   ├── getting-started.md              # Getting started guide
│   │   ├── workflow.md                     # Development workflow
│   │   ├── frontend.md                     # Frontend development
│   │   └── backend.md                      # Backend development
│   ├── deployment/                         # Deployment documentation
│   │   ├── environments.md                 # Environment configuration
│   │   ├── ci-cd.md                        # CI/CD pipeline
│   │   └── aws.md                          # AWS deployment
│   └── api/                                # API documentation
│       ├── auth-service.yaml               # Auth service OpenAPI spec
│       ├── blog-service.yaml               # Blog service OpenAPI spec
│       └── user-service.yaml               # User service OpenAPI spec
│
├── .gitignore                              # Git ignore file
├── .editorconfig                           # Editor configuration
├── .prettierrc                             # Prettier configuration
├── docker-compose.yml                      # Docker Compose configuration
├── go.work                                 # Go workspace file
├── package.json                            # Root package.json
├── pnpm-workspace.yaml                     # pnpm workspace configuration
└── README.md                               # Project README
```
