package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	FindByUsername(username string) (*User, error)
	FindAll() ([]User, error)
	FindByID(id uuid.UUID) (*User, error)
	Update(user *User) error
	Delete(id uuid.UUID) error

	FindByRefreshToken(token string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *repository) FindByRefreshToken(token string) (*User, error) {
	var user User
	err := r.db.Where("refresh_token = ?", token).First(&user).Error
	return &user, err
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	// Bisa ditambahkan Preload("Role") jika ada relasi struct Role
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindByID(id uuid.UUID) (*User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
