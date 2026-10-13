package roles

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Role, error)
	FindByID(id string) (*Role, error)
	Create(role *Role) error
	Update(role *Role, permissionIDs []string) error
	Delete(id string) error

	GetAllPermissions() ([]Permission, error)
	GetPermissionsByIDs(ids []string) ([]Permission, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Role, error) {
	var roles []Role
	err := r.db.Preload("Permissions").Find(&roles).Error
	return roles, err
}

func (r *repository) FindByID(id string) (*Role, error) {
	var role Role
	err := r.db.Preload("Permissions").First(&role, "id = ?", id).Error
	return &role, err
}

func (r *repository) Create(role *Role) error {
	return r.db.Create(role).Error
}

func (r *repository) Update(role *Role, permissionIDs []string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(role).Updates(role).Error; err != nil {
			return err
		}

		if len(permissionIDs) > 0 {
			var perms []Permission
			if err := tx.Where("id IN ?", permissionIDs).Find(&perms).Error; err != nil {
				return err
			}
			if err := tx.Model(role).Association("Permissions").Replace(perms); err != nil {
				return err
			}
		} else {
			if err := tx.Model(role).Association("Permissions").Clear(); err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *repository) Delete(id string) error {
	return r.db.Select("Permissions").Delete(&Role{ID: id}).Error
}

func (r *repository) GetAllPermissions() ([]Permission, error) {
	var perms []Permission
	err := r.db.Find(&perms).Error
	return perms, err
}

func (r *repository) GetPermissionsByIDs(ids []string) ([]Permission, error) {
	var perms []Permission
	if len(ids) == 0 {
		return perms, nil
	}
	err := r.db.Where("id IN ?", ids).Find(&perms).Error
	return perms, err
}
