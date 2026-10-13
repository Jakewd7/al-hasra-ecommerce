package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              string  `gorm:"type:varchar(100);primaryKey"`
	RoleID          *string `gorm:"type:varchar(100)"`
	Username        string  `gorm:"type:varchar(100);unique;not null"`
	Email           *string `gorm:"type:varchar(150);unique"`
	Password        string  `gorm:"type:varchar(255);not null"`
	FullName        string  `gorm:"type:varchar(150)"`
	Phone           string  `gorm:"type:varchar(20)"`
	DeviceID        string  `gorm:"type:varchar(255)"`
	IsActive        bool    `gorm:"default:true"`
	LastLoginAt     *time.Time
	EmailVerifiedAt *time.Time
	RefreshToken    string `gorm:"type:text;index"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password123"`
}

type LoginResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	UserID       string  `json:"user_id"`
	RoleID       *string `json:"role_id"`
}

type RegisterRequest struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"secret123"`
	FullName string `json:"full_name" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
	RoleID   string `json:"role_id" example:"role-uuid-string"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" example:"John Doe Updated"`
	RoleID   string `json:"role_id" example:"role-uuid-string"`
	IsActive *bool  `json:"is_active" example:"true"`
}

type UserResponse struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	FullName string  `json:"full_name"`
	Email    *string `json:"email"`
	RoleID   *string `json:"role_id"`
}
