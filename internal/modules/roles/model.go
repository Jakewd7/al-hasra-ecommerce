package roles

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Name        string       `gorm:"type:varchar(50);unique;not null"`
	Description string       `gorm:"type:text"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Permission struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"type:varchar(50);unique;not null"`
	Description string    `gorm:"type:text"`
	Module      string    `gorm:"type:varchar(50)"`
	Action      string    `gorm:"type:varchar(50)"`
	Roles       []Role    `gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RolePermission struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	PermissionID uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

func (rp *RolePermission) BeforeCreate(tx *gorm.DB) error {
	if rp.ID == uuid.Nil {
		rp.ID = uuid.New()
	}
	return nil
}

type CreateRoleRequest struct {
	Name          string   `json:"name" example:"Store Manager" binding:"required"`
	Description   string   `json:"description" example:"Manager of a specific store branch"`
	PermissionIDs []string `json:"permission_ids" example:"['uuid-1', 'uuid-2']"`
}

type UpdateRoleRequest struct {
	Name          string   `json:"name" example:"Store Manager Updated"`
	Description   string   `json:"description"`
	PermissionIDs []string `json:"permission_ids"`
}

type RoleResponse struct {
	ID          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Permissions []PermissionResponse `json:"permissions"`
}

type PermissionResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Module string    `json:"module"`
	Action string    `json:"action"`
}
