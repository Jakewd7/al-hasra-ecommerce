package roles

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	GetAllRoles() ([]RoleResponse, error)
	GetRoleByID(id uuid.UUID) (RoleResponse, error)
	CreateRole(input CreateRoleRequest) (RoleResponse, error)
	UpdateRole(id uuid.UUID, input UpdateRoleRequest) (RoleResponse, error)
	DeleteRole(id uuid.UUID) error

	GetAllPermissions() ([]PermissionResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func mapRoleToResponse(role Role) RoleResponse {
	var permResponses []PermissionResponse
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

	var responses []RoleResponse
	for _, r := range roles {
		responses = append(responses, mapRoleToResponse(r))
	}
	return responses, nil
}

func (s *service) GetRoleByID(id uuid.UUID) (RoleResponse, error) {
	role, err := s.repository.FindByID(id)
	if err != nil {
		return RoleResponse{}, err
	}
	return mapRoleToResponse(*role), nil
}

func (s *service) CreateRole(input CreateRoleRequest) (RoleResponse, error) {
	var permUUIDs []uuid.UUID
	for _, idStr := range input.PermissionIDs {
		id, err := uuid.Parse(idStr)
		if err == nil {
			permUUIDs = append(permUUIDs, id)
		}
	}

	perms, err := s.repository.GetPermissionsByIDs(permUUIDs)
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

func (s *service) UpdateRole(id uuid.UUID, input UpdateRoleRequest) (RoleResponse, error) {
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

	var permUUIDs []uuid.UUID
	for _, idStr := range input.PermissionIDs {
		pid, err := uuid.Parse(idStr)
		if err == nil {
			permUUIDs = append(permUUIDs, pid)
		}
	}

	if err := s.repository.Update(role, permUUIDs); err != nil {
		return RoleResponse{}, err
	}

	updatedRole, _ := s.repository.FindByID(id)
	return mapRoleToResponse(*updatedRole), nil
}

func (s *service) DeleteRole(id uuid.UUID) error {
	return s.repository.Delete(id)
}

func (s *service) GetAllPermissions() ([]PermissionResponse, error) {
	perms, err := s.repository.GetAllPermissions()
	if err != nil {
		return nil, err
	}

	var responses []PermissionResponse
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
