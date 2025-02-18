package user

import (
	"errors"
	"square-pos/pkg/config"
	"square-pos/pkg/types"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

// NewStore initializes the user store with a database connection
func NewStore() *Store {
	return &Store{db: config.DB}
}

// CreateUser inserts a new user into the database
func (s *Store) CreateUser(user *types.User) error {
	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
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
func (s *Store) GetUserByID(id uint) (*types.User, error) {
	var user types.User
	result := s.db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}
