package seeders

import (
	"fmt"
	"log"

	"github.com/jakewd7/internal/modules/roles"
	"gorm.io/gorm"
)

type ModuleMap struct {
	Module    string
	Resources []string
}

func SeedDatabase(db *gorm.DB) {
	modules := []ModuleMap{
		// define your module/routes to the middleware db seed
		{
			Module:    "dashboard",
			Resources: []string{"dashboard"},
		},
	}

	// Actions standar CRUD
	actions := []string{"read", "create", "update", "delete"}

	var allPermissions []roles.Permission

	log.Println("Menanam pundi-pundi permission...")

	for _, mod := range modules {
		for _, res := range mod.Resources {
			for _, act := range actions {
				permName := fmt.Sprintf("%s.%s.%s", mod.Module, res, act)

				permission := roles.Permission{
					Name:        permName,
					Module:      mod.Module,
					Action:      act,
					Description: fmt.Sprintf("Allow user to %s %s in %s", act, res, mod.Module),
				}

				if err := db.Where("name = ?", permName).FirstOrCreate(&permission).Error; err != nil {
					log.Printf("Error seeding permission %s: %v", permName, err)
				} else {
					allPermissions = append(allPermissions, permission)
				}
			}
		}
	}

	log.Println("Permissions seeded successfully.")

	seedRoles(db, allPermissions)
}

func seedRoles(db *gorm.DB, allPermissions []roles.Permission) {
	log.Println("Menanam Pundi-pundi rupiah...")

	superAdmin := roles.Role{
		Name:        "superadmin",
		Description: "Has access to all features",
	}

	if err := db.Where("name = ?", superAdmin.Name).FirstOrCreate(&superAdmin).Error; err != nil {
		log.Printf("Error creating Super Admin: %v", err)
	}

	if err := db.Model(&superAdmin).Association("Permissions").Replace(allPermissions); err != nil {
		log.Printf("Error assigning permissions to Super Admin: %v", err)
	}

	log.Println("Super Admin created/updated with all permissions.")
}
