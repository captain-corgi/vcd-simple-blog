# Git Flow: A Comprehensive Guide

Git Flow is a branching model that provides a structured approach to managing your Git repository. This guide will walk you through how to use Git Flow effectively, using our VCD Simple Blog project as a practical example.

## Table of Contents

1. [Introduction to Git Flow](#introduction-to-git-flow)
2. [Setting Up Git Flow](#setting-up-git-flow)
3. [Working with Features](#working-with-features)
4. [Managing Releases](#managing-releases)
5. [Handling Hotfixes](#handling-hotfixes)
6. [Commit Message Conventions](#commit-message-conventions)
7. [Best Practices](#best-practices)
8. [Real-World Example: VCD Simple Blog](#real-world-example-vcd-simple-blog)

## Introduction to Git Flow

Git Flow defines a strict branching model designed around project releases. It assigns specific roles to different branches and defines how they should interact.

### Core Branches

- **master/main**: Contains production-ready code
- **develop**: Main development branch where features are integrated

### Supporting Branches

- **feature/\***: For developing new features
- **release/\***: For preparing new production releases
- **hotfix/\***: For critical fixes to production code

## Setting Up Git Flow

### Installation

First, install the Git Flow extension:

```bash
# For Windows (with Git for Windows)
# It's included by default

# For macOS (using Homebrew)
brew install git-flow-avh

# For Linux
apt-get install git-flow  # Debian/Ubuntu
yum install gitflow       # CentOS/RHEL
```

### Initializing Git Flow in a Repository

```bash
# Navigate to your repository
cd your-repository

# Initialize Git Flow with default settings
git flow init -d

# Or for custom settings
git flow init
```

**Example from VCD Simple Blog:**

```bash
git flow init -d
```

This created our `master` and `develop` branches with `develop` as our default working branch.

## Working with Features

Features are developed in dedicated branches and later merged back into the `develop` branch.

### Starting a Feature

```bash
git flow feature start feature-name
```

This creates a new branch `feature/feature-name` based on `develop`.

**Example from VCD Simple Blog:**

```bash
git flow feature start frontend-auth-module
```

This created a new branch `feature/frontend-auth-module` where we implemented the authentication UI components.

### Working on a Feature

Once the feature branch is created, you can work on it as you would with any Git branch:

```bash
# Make changes
# Add files
git add .

# Commit changes
git commit -m "feat: add login component"
```

**Example from VCD Simple Blog:**

```bash
# Created authentication components
git add apps/frontend/auth/
git commit -m "feat(frontend): add authentication module with login, register, and password reset pages"
```

### Finishing a Feature

When the feature is complete, finish it with:

```bash
git flow feature finish feature-name
```

This will:
1. Merge the feature branch into `develop`
2. Delete the feature branch
3. Switch back to the `develop` branch

**Example from VCD Simple Blog:**

```bash
git flow feature finish frontend-auth-module
```

This merged our authentication module into the `develop` branch and deleted the feature branch.

## Managing Releases

Releases prepare the code for production deployment.

### Starting a Release

```bash
git flow release start version-number
```

This creates a new branch `release/version-number` based on `develop`.

**Example from VCD Simple Blog:**

```bash
git flow release start 1.1.0
```

This created a `release/1.1.0` branch where we prepared for the 1.1.0 release.

### Working on a Release

During the release phase, you typically:
- Bump version numbers
- Update CHANGELOG
- Make minor bug fixes
- Prepare documentation

```bash
# Update version
# Edit files as needed
git add .
git commit -m "chore: bump version to 1.1.0"
```

**Example from VCD Simple Blog:**

```bash
# Updated version in package.json
git add package.json
# Added CHANGELOG.md
git add CHANGELOG.md
git commit -m "chore: bump version to 1.1.0 and add CHANGELOG"
```

### Finishing a Release

```bash
git flow release finish version-number
```

This will:
1. Merge the release branch into `master`
2. Tag the release with its version number
3. Merge the release back into `develop`
4. Delete the release branch

**Example from VCD Simple Blog:**

```bash
git flow release finish -m "Release version 1.1.0" 1.1.0
```

This merged our release into both `master` and `develop`, created a tag `1.1.0`, and deleted the release branch.

## Handling Hotfixes

Hotfixes are used to quickly patch production releases.

### Starting a Hotfix

```bash
git flow hotfix start version-number
```

This creates a new branch `hotfix/version-number` based on `master`.

**Example from VCD Simple Blog:**

```bash
git flow hotfix start 1.0.1
```

This created a `hotfix/1.0.1` branch to fix a critical issue in the production code.

### Working on a Hotfix

```bash
# Fix the issue
git add .
git commit -m "fix: critical issue"
```

**Example from VCD Simple Blog:**

```bash
# Added missing format script and updated version
git add package.json
git commit -m "fix: add format script and bump version to 1.0.1"
```

### Finishing a Hotfix

```bash
git flow hotfix finish version-number
```

This will:
1. Merge the hotfix branch into `master`
2. Tag the release with the new version number
3. Merge the hotfix back into `develop`
4. Delete the hotfix branch

**Example from VCD Simple Blog:**

```bash
git flow hotfix finish -m "Hotfix 1.0.1" 1.0.1
```

This merged our hotfix into both `master` and `develop`, created a tag `1.0.1`, and deleted the hotfix branch.

## Commit Message Conventions

Using conventional commit messages makes your repository history more readable and enables automated tools like semantic versioning.

### Format

```
<type>(<scope>): <subject>
```

### Types

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation changes
- **style**: Code style changes (formatting, etc.)
- **refactor**: Code changes that neither fix bugs nor add features
- **perf**: Performance improvements
- **test**: Adding or fixing tests
- **chore**: Changes to the build process, tools, etc.

### Examples from VCD Simple Blog

```bash
git commit -m "feat(frontend): add authentication module with login, register, and password reset pages"
git commit -m "feat(backend): add auth service with user and token domain entities"
git commit -m "fix: add format script and bump version to 1.0.1"
git commit -m "chore: bump version to 1.1.0 and add CHANGELOG"
```

## Best Practices

1. **Keep feature branches focused**: Each feature branch should implement a single feature or fix.
2. **Regularly pull from develop**: Keep your feature branches up-to-date with the latest changes in develop.
3. **Use descriptive branch names**: Name your branches clearly to indicate what they contain.
4. **Write meaningful commit messages**: Follow the conventional commit format.
5. **Create a CHANGELOG**: Document changes for each release.
6. **Tag releases properly**: Use semantic versioning for your tags.
7. **Clean up old branches**: Delete branches that are no longer needed.

## Real-World Example: VCD Simple Blog

Let's walk through the complete Git Flow process we used for the VCD Simple Blog project:

### 1. Initial Setup

```bash
# Initialize Git Flow
git init
git flow init -d

# Initial commit on develop branch
git add .
git commit -m "Initial project structure for VCD Simple Blog"
```

### 2. First Release (1.0.0)

```bash
# Create release branch
git flow release start 1.0.0

# Update version
git add package.json
git commit -m "Bump version to 1.0.0"

# Finish release
git flow release finish -m "Release version 1.0.0" 1.0.0
```

### 3. Feature Development: Frontend Auth Module

```bash
# Create feature branch
git flow feature start frontend-auth-module

# Implement feature
git add apps/frontend/auth/
git commit -m "feat(frontend): add authentication module with login, register, and password reset pages"

# Finish feature
git flow feature finish frontend-auth-module
```

### 4. Feature Development: Backend Auth Service

```bash
# Create feature branch
git flow feature start backend-auth-service

# Implement feature
git add apps/backend/auth-service/
git commit -m "feat(backend): add auth service with user and token domain entities"

# Finish feature
git flow feature finish backend-auth-service
```

### 5. Hotfix for Production (1.0.1)

```bash
# Create hotfix branch
git flow hotfix start 1.0.1

# Fix issue and bump version
git add package.json
git commit -m "fix: add format script and bump version to 1.0.1"

# Finish hotfix
git flow hotfix finish -m "Hotfix 1.0.1" 1.0.1
```

### 6. Feature Development: Infrastructure CI/CD

```bash
# Create feature branch
git flow feature start infrastructure-ci-cd

# Implement feature
git add .github/workflows/cd.yml apps/frontend/shell/Dockerfile apps/frontend/shell/nginx.conf apps/backend/auth-service/Dockerfile
git commit -m "feat(infrastructure): add CI/CD workflow and Dockerfiles"

# Finish feature
git flow feature finish infrastructure-ci-cd
```

### 7. Second Release (1.1.0)

```bash
# Create release branch
git flow release start 1.1.0

# Update version and add CHANGELOG
git add package.json CHANGELOG.md
git commit -m "chore: bump version to 1.1.0 and add CHANGELOG"

# Finish release
git flow release finish -m "Release version 1.1.0" 1.1.0
```

### 8. Final Repository State

```bash
# Check branches
git branch
# * develop
#   master

# Check commit history
git log --oneline -n 10
# a72146a (HEAD -> develop) Merge tag '1.1.0' into develop
# 184ae5f (tag: 1.1.0, master) Merge branch 'release/1.1.0'
# dfbdad1 chore: bump version to 1.1.0 and add CHANGELOG
# b0141c7 feat(infrastructure): add CI/CD workflow and Dockerfiles
# b923936 Merge tag '1.0.1' into develop
# 98424c8 (tag: 1.0.1, tag: 1.0.0) Merge branch 'hotfix/1.0.1'
# b8d9dea fix: add format script and bump version to 1.0.1
# 4fa9ad8 feat(backend): add auth service with user and token domain entities
# cd6e37e feat(frontend): add authentication module with login, register, and password reset pages
# 992d8a1 Merge release 1.0.0 into master
```

## Conclusion

Git Flow provides a structured approach to managing your Git repository, making it easier to develop features, prepare releases, and fix critical issues. By following this branching model, you can maintain a clean and organized repository history, facilitate collaboration among team members, and ensure a smooth release process.

The VCD Simple Blog project demonstrates how Git Flow can be applied to a real-world project, from initial setup to feature development, hotfixes, and releases. By adopting these practices, you can improve your development workflow and maintain a high-quality codebase.
