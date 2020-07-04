package users

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

// User holds all data required to represent a user
type User struct {
	FirstName string     `json:"firstName,omitempty"`
	LastName  string     `json:"lastName,omitempty"`
	Mobile    string     `json:"mobile,omitempty"`
	Email     string     `json:"email,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func (u *User) setDefaults() {
	now := time.Now()
	if u.CreatedAt == nil {
		u.CreatedAt = &now
	}

	if u.UpdatedAt == nil {
		u.UpdatedAt = &now
	}
}

// Sanitize is used to sanitize/cleanup the fields of User
func (u *User) Sanitize() {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(u.Email)
	u.Mobile = strings.TrimSpace(u.Mobile)
}

// Validate is used to validate the fields of User
func (u *User) Validate() error {
	if u.Email != "" {
		parts := strings.Split(u.Email, "@")
		if len(parts) != 2 {
			return errors.New("invalid email address provided")
		}
	}
	return nil
}

// Users struct holds all the dependencies required for the users package. And exposes all services
// provided by this package as its methods
type Users struct {
	// logger
	// cache
	store store
}

// CreateUser creates a new user
func (us *Users) CreateUser(ctx context.Context, u *User) (*User, error) {
	u.setDefaults()
	u.Sanitize()

	err := u.Validate()
	if err != nil {
		// this wrapping helps identify where the error originated when logging at a higher level
		// e.g. if logging is done at `api` package
		return nil, fmt.Errorf("Validate: %w", err)
	}

	err = us.store.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("store.Create: %w", err)
	}

	return u, nil
}

// NewService initializes the Users struct with all its dependencies and returns a new instance
func NewService() (*Users, error) {
	// all dependencies of Users should be sent as arguments of NewService
	return &Users{}, nil
}