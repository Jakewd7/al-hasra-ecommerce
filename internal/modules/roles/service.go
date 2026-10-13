package roles

import (
	"errors"
)

type Service interface {
	GetAllRoles() ([]RoleResponse, error)
	GetRoleByID(id string) (RoleResponse, error)
	CreateRole(input CreateRoleRequest) (RoleResponse, error)
	UpdateRole(id string, input UpdateRoleRequest) (RoleResponse, error)
	DeleteRole(id string) error

	GetAllPermissions() ([]PermissionResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func mapRoleToResponse(role Role) RoleResponse {
	permResponses := make([]PermissionResponse, 0)
	for _, p := range role.Permissions {
		permResponses = append(permResponses, PermissionResponse{
			ID:     p.ID,
			Name:   p.Name,
			Module: p.Module,
			Action: p.Action,
		})
	}
	return RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		Permissions: permResponses,
	}
}

func (s *service) GetAllRoles() ([]RoleResponse, error) {
	roles, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]RoleResponse, 0)
	for _, r := range roles {
		responses = append(responses, mapRoleToResponse(r))
	}
	return responses, nil
}

func (s *service) GetRoleByID(id string) (RoleResponse, error) {
	role, err := s.repository.FindByID(id)
	if err != nil {
		return RoleResponse{}, err
	}
	return mapRoleToResponse(*role), nil
}

func (s *service) CreateRole(input CreateRoleRequest) (RoleResponse, error) {
	perms, err := s.repository.GetPermissionsByIDs(input.PermissionIDs)
	if err != nil {
		return RoleResponse{}, err
	}

	role := Role{
		Name:        input.Name,
		Description: input.Description,
		Permissions: perms,
	}

	if err := s.repository.Create(&role); err != nil {
		return RoleResponse{}, err
	}

	return mapRoleToResponse(role), nil
}

func (s *service) UpdateRole(id string, input UpdateRoleRequest) (RoleResponse, error) {
	role, err := s.repository.FindByID(id)
	if err != nil {
		return RoleResponse{}, errors.New("role not found")
	}

	if input.Name != "" {
		role.Name = input.Name
	}
	if input.Description != "" {
		role.Description = input.Description
	}

	if err := s.repository.Update(role, input.PermissionIDs); err != nil {
		return RoleResponse{}, err
	}

	updatedRole, _ := s.repository.FindByID(id)
	return mapRoleToResponse(*updatedRole), nil
}

func (s *service) DeleteRole(id string) error {
	return s.repository.Delete(id)
}

func (s *service) GetAllPermissions() ([]PermissionResponse, error) {
	perms, err := s.repository.GetAllPermissions()
	if err != nil {
		return nil, err
	}

	responses := make([]PermissionResponse, 0)
	for _, p := range perms {
		responses = append(responses, PermissionResponse{
			ID:     p.ID,
			Name:   p.Name,
			Module: p.Module,
			Action: p.Action,
		})
	}
	return responses, nil
}
