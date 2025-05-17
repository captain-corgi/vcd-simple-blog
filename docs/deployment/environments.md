# Environments

The VCD Simple Blog application supports multiple deployment environments to facilitate the development, testing, and release process.

## Environment Types

We maintain the following environments:

### 1. Development (DEV)

- **Purpose**: Daily development and integration
- **Audience**: Developers
- **Data**: Non-sensitive test data
- **Deployment**: Automatic on merge to develop branch
- **Configuration**: Minimal resources, debug enabled

### 2. System Integration Testing (SIT)

- **Purpose**: Integration testing of features
- **Audience**: QA team, developers
- **Data**: Structured test data
- **Deployment**: Manual trigger after DEV validation
- **Configuration**: Similar to production but with debug enabled

### 3. User Acceptance Testing (UAT)

- **Purpose**: User acceptance testing
- **Audience**: Business stakeholders, QA team
- **Data**: Production-like data
- **Deployment**: Manual trigger after SIT validation
- **Configuration**: Production-like

### 4. Non-Functional Testing (NFT)

- **Purpose**: Performance, security, and load testing
- **Audience**: Performance engineers, security team
- **Data**: Production-like volume data
- **Deployment**: Manual trigger as needed
- **Configuration**: Production-like or higher capacity

### 5. Staging (STG)

- **Purpose**: Final verification before production
- **Audience**: Release team, QA team
- **Data**: Sanitized production data
- **Deployment**: Manual trigger after UAT validation
- **Configuration**: Identical to production

### 6. Production (PRD)

- **Purpose**: Live application
- **Audience**: End users
- **Data**: Real production data
- **Deployment**: Manual trigger after STG validation
- **Configuration**: Full production resources

## Environment Configuration

Environment-specific configuration is managed through:

1. **Environment Variables**: Set in the deployment pipeline
2. **Terraform Variables**: Different values per environment
3. **Feature Flags**: Control feature availability per environment

### Example Environment Variables

```
# Development
ENV=development
LOG_LEVEL=debug
ENABLE_SWAGGER=true

# Production
ENV=production
LOG_LEVEL=info
ENABLE_SWAGGER=false
```

### Example Terraform Variables

```hcl
# Development
environment     = "dev"
instance_class  = "db.t3.small"
desired_count   = 1

# Production
environment     = "prd"
instance_class  = "db.m5.large"
desired_count   = 3
```

## Environment Isolation

Each environment is completely isolated from others:

1. **Separate AWS Accounts**: Each environment has its own AWS account
2. **Separate Databases**: No shared databases between environments
3. **Separate Domains**: Each environment has its own domain or subdomain

Example domains:
- dev.vcd-simple-blog.example.com
- sit.vcd-simple-blog.example.com
- uat.vcd-simple-blog.example.com
- nft.vcd-simple-blog.example.com
- stg.vcd-simple-blog.example.com
- vcd-simple-blog.example.com (Production)

## Promotion Process

Code changes follow a promotion process through environments:

1. Developers work on feature branches
2. Feature branches are merged to develop branch (deploys to DEV)
3. Develop branch is merged to sit branch (deploys to SIT)
4. SIT branch is merged to uat branch (deploys to UAT)
5. UAT branch is merged to staging branch (deploys to STG)
6. Staging branch is merged to main branch (deploys to PRD)

Each promotion requires:
- Successful tests in the previous environment
- Code review and approval
- Manual approval for production deployment

## Environment Management

### Creation

New environments are created using Terraform:

```bash
cd infrastructure/terraform
terraform workspace select dev
terraform apply -var-file=environments/dev.tfvars
```

### Updates

Environment updates follow the same process:

```bash
cd infrastructure/terraform
terraform workspace select dev
terraform apply -var-file=environments/dev.tfvars
```

### Teardown

Temporary environments can be torn down when no longer needed:

```bash
cd infrastructure/terraform
terraform workspace select feature-x
terraform destroy -var-file=environments/feature-x.tfvars
```

## Access Control

Access to environments is restricted based on role:

- **DEV**: All developers have full access
- **SIT/UAT**: Developers have read access, QA has full access
- **NFT**: Performance engineers have full access
- **STG**: Release team has full access
- **PRD**: Limited access through break-glass procedures

## Monitoring

All environments are monitored, with different alert thresholds:

- **DEV/SIT/UAT**: Alerts go to development team
- **NFT**: Alerts go to performance team
- **STG/PRD**: Alerts go to operations team and on-call rotation
