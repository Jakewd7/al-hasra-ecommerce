package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	BranchID        uuid.UUID `gorm:"type:uuid;not null"`
	Username        string    `gorm:"unique;not null"`
	Email           *string   `gorm:"unique"`
	Password        string    `gorm:"not null"`
	FullName        string
	Phone           string
	RoleID          uuid.UUID `gorm:"type:uuid"`
	DeviceID        string
	IsActive        bool `gorm:"default:true"`
	LastLoginAt     *time.Time
	EmailVerifiedAt *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	RefreshToken    string `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password123"`
}

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	UserID       uuid.UUID `json:"user_id"`
	RoleID       uuid.UUID `json:"role_id"`
}

type RegisterRequest struct {
	BranchID string `json:"branch_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"secret123"`
	FullName string `json:"full_name" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
	RoleID   string `json:"role_id" example:"550e8400-e29b-41d4-a716-446655440001"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" example:"John Doe Updated"`
	RoleID   string `json:"role_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	IsActive *bool  `json:"is_active" example:"true"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Email    *string   `json:"email"`
	RoleID   uuid.UUID `json:"role_id"`
	BranchID uuid.UUID `json:"branch_id"`
}
