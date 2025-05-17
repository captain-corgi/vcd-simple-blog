# Clean Architecture

The VCD Simple Blog application follows Clean Architecture principles to ensure separation of concerns, testability, and maintainability.

## Overview

Clean Architecture organizes code into concentric layers, with each layer having a specific responsibility. The fundamental rule is that dependencies can only point inward - outer layers can depend on inner layers, but inner layers cannot depend on outer layers.

## Layers

Our implementation of Clean Architecture consists of the following layers:

### 1. Domain Layer (Innermost)

The domain layer contains the core business logic and entities of the application. It has no dependencies on other layers or external frameworks.

**Components:**
- **Entities**: Core business objects (e.g., Blog, User, Comment)
- **Value Objects**: Immutable objects that represent concepts with no identity (e.g., BlogStatus)
- **Domain Services**: Services that operate on multiple entities
- **Repository Interfaces**: Interfaces defining data access methods

**Example:**
```go
// Domain Entity
type Blog struct {
    ID          string
    Title       string
    Content     string
    AuthorID    string
    Status      valueobject.BlogStatus
    Tags        []string
    PublishedAt *time.Time
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// Domain Repository Interface
type BlogRepository interface {
    FindAll(ctx context.Context, limit, offset int) ([]*entity.Blog, error)
    FindByID(ctx context.Context, id string) (*entity.Blog, error)
    Create(ctx context.Context, blog *entity.Blog) error
    Update(ctx context.Context, blog *entity.Blog) error
    Delete(ctx context.Context, id string) error
}
```

### 2. Application Layer

The application layer contains the use cases of the application. It orchestrates the flow of data to and from entities and directs them to perform their business logic.

**Components:**
- **Use Cases**: Application-specific business rules
- **DTOs**: Data Transfer Objects for input/output
- **Interfaces**: Interfaces for external services

**Example:**
```go
// Use Case
type BlogUseCase struct {
    blogRepo repository.BlogRepository
}

func (uc *BlogUseCase) CreateBlog(ctx context.Context, title, content, authorID string, tags []string) (*entity.Blog, error) {
    id := uuid.New().String()
    blog, err := entity.NewBlog(id, title, content, authorID, tags)
    if err != nil {
        return nil, err
    }

    if err := uc.blogRepo.Create(ctx, blog); err != nil {
        return nil, err
    }

    return blog, nil
}
```

### 3. Infrastructure Layer

The infrastructure layer contains implementations of the interfaces defined in the inner layers. It deals with external concerns like databases, external APIs, and frameworks.

**Components:**
- **Repository Implementations**: Database access implementations
- **External Service Clients**: HTTP clients, gRPC clients
- **ORM Models**: Database models and mappings

**Example:**
```go
// Repository Implementation
type BlogRepository struct {
    db *gorm.DB
}

func (r *BlogRepository) FindByID(ctx context.Context, id string) (*entity.Blog, error) {
    var model models.Blog
    if err := r.db.First(&model, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return model.ToEntity(), nil
}
```

### 4. Interfaces Layer (Outermost)

The interfaces layer contains code that interacts with the outside world, such as API controllers, CLI commands, and event handlers.

**Components:**
- **API Controllers**: HTTP handlers
- **Middleware**: Authentication, logging, etc.
- **Request/Response Models**: API-specific data structures

**Example:**
```go
// API Controller
type BlogHandler struct {
    blogUseCase *usecases.BlogUseCase
}

func (h *BlogHandler) CreateBlog(c echo.Context) error {
    var req requests.CreateBlogRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }
    
    userID := c.Get("userID").(string)
    blog, err := h.blogUseCase.CreateBlog(c.Request().Context(), req.Title, req.Content, userID, req.Tags)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    
    return c.JSON(http.StatusCreated, responses.FromBlog(blog))
}
```

## Benefits

1. **Independence from Frameworks**: The core business logic doesn't depend on external frameworks or libraries.
2. **Testability**: Each layer can be tested in isolation with appropriate mocks.
3. **Maintainability**: Changes in one layer don't affect other layers as long as interfaces remain stable.
4. **Flexibility**: External components like databases can be replaced without affecting the core business logic.
