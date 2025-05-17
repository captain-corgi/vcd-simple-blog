package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/vcd-simple-blog/apps/backend/user-service/domain/entity"
	"github.com/vcd-simple-blog/apps/backend/user-service/domain/repository"
	"github.com/vcd-simple-blog/apps/backend/user-service/domain/valueobject"
)

// UserUseCase implements user-related use cases
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// GetUserByID retrieves a user by ID
func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return uc.userRepo.FindByID(ctx, id)
}

// GetUserByUserID retrieves a user by auth service user ID
func (uc *UserUseCase) GetUserByUserID(ctx context.Context, userID string) (*entity.User, error) {
	return uc.userRepo.FindByUserID(ctx, userID)
}

// GetUserByUsername retrieves a user by username
func (uc *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return uc.userRepo.FindByUsername(ctx, username)
}

// GetAllUsers retrieves all users with pagination
func (uc *UserUseCase) GetAllUsers(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	if limit <= 0 {
		limit = 10
	}
	return uc.userRepo.FindAll(ctx, limit, offset)
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(ctx context.Context, userID, username, email string) (*entity.User, error) {
	// Check if user already exists with this userID
	existingUser, err := uc.userRepo.FindByUserID(ctx, userID)
	if err == nil && existingUser != nil {
		return nil, errors.New("user already exists with this user ID")
	}

	// Check if username is taken
	existingUser, err = uc.userRepo.FindByUsername(ctx, username)
	if err == nil && existingUser != nil {
		return nil, errors.New("username is already taken")
	}

	// Create new user
	id := uuid.New().String()
	user, err := entity.NewUser(id, userID, username, email)
	if err != nil {
		return nil, err
	}

	// Save user to database
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserProfile updates a user's profile
func (uc *UserUseCase) UpdateUserProfile(ctx context.Context, id, displayName, bio, avatarURL string) (*entity.User, error) {
	// Find user
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update profile
	user.UpdateProfile(displayName, bio, avatarURL)

	// Save changes
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateProfileStatus updates a user's profile status
func (uc *UserUseCase) UpdateProfileStatus(ctx context.Context, id string, status valueobject.ProfileStatus) (*entity.User, error) {
	// Find user
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update profile status
	user.SetProfileStatus(status)

	// Save changes
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user
func (uc *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	// Check if user exists
	_, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete user
	return uc.userRepo.Delete(ctx, id)
}
