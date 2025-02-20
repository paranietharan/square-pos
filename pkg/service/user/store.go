package user

import (
	"errors"
	"square-pos/pkg/dto"
	"square-pos/pkg/types"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

// NewStore initializes the user store with a database connection
func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// CreateUser inserts a new user into the database
func (s *Store) CreateUser(user *types.User) (dto.UserCreateResponse, error) {
	result := s.db.Create(user)
	if result.Error != nil {
		return dto.UserCreateResponse{}, result.Error
	}
	return dto.UserCreateResponse{Message: "UserCreated successfully"}, nil
}

// GetUserByEmail retrieves a user by their email address
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	result := s.db.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}

// GetUserByID retrieves a user by their ID
func (s *Store) GetUserByID(id int) (*types.User, error) {
	var user types.User
	result := s.db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}
