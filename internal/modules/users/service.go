package users

import (
	"errors"
	"time"

	"github.com/jakewd7/internal/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterRequest) (UserResponse, error)
	LoginUser(input LoginRequest) (LoginResponse, error)
	RefreshToken(refreshToken string) (LoginResponse, error)
	GetAllUsers() ([]UserResponse, error)
	UpdateUser(id string, input UpdateUserRequest) (UserResponse, error)
	DeleteUser(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterRequest) (UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, err
	}

	var roleID *string
	if input.RoleID != "" {
		roleID = &input.RoleID
	}

	user := User{
		Username: input.Username,
		Password: string(hashedPassword),
		FullName: input.FullName,
		Email:    &input.Email,
		RoleID:   roleID,
		IsActive: true,
	}

	if err := s.repository.Create(&user); err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		RoleID:   user.RoleID,
		Email:    user.Email,
	}, nil
}

func (s *service) LoginUser(input LoginRequest) (LoginResponse, error) {
	user, err := s.repository.FindByUsername(input.Username)
	if err != nil {
		return LoginResponse{}, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return LoginResponse{}, errors.New("invalid credentials")
	}

	var roleIDStr string
	if user.RoleID != nil {
		roleIDStr = *user.RoleID
	}

	accessToken, err := utils.GenerateJWT(user.ID, roleIDStr)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken := utils.GenerateSHA256Token()

	user.RefreshToken = refreshToken
	if err := s.repository.Update(user); err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID,
		RoleID:       user.RoleID,
	}, nil
}

func (s *service) GetAllUsers() ([]UserResponse, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var userResponses []UserResponse
	for _, u := range users {
		userResponses = append(userResponses, UserResponse{
			ID:       u.ID,
			Username: u.Username,
			FullName: u.FullName,
			Email:    u.Email,
			RoleID:   u.RoleID,
		})
	}
	return userResponses, nil
}

func (s *service) UpdateUser(id string, input UpdateUserRequest) (UserResponse, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return UserResponse{}, err
	}

	if input.FullName != "" {
		user.FullName = input.FullName
	}
	if input.RoleID != "" {
		user.RoleID = &input.RoleID
	}
	if input.IsActive != nil {
		user.IsActive = *input.IsActive
	}

	user.UpdatedAt = time.Now()

	if err := s.repository.Update(user); err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		RoleID:   user.RoleID,
		Email:    user.Email,
	}, nil
}

func (s *service) DeleteUser(id string) error {
	return s.repository.Delete(id)
}

func (s *service) RefreshToken(refreshToken string) (LoginResponse, error) {
	user, err := s.repository.FindByRefreshToken(refreshToken)
	if err != nil {
		return LoginResponse{}, errors.New("invalid refresh token")
	}

	var roleIDStr string
	if user.RoleID != nil {
		roleIDStr = *user.RoleID
	}

	newAccessToken, err := utils.GenerateJWT(user.ID, roleIDStr)
	if err != nil {
		return LoginResponse{}, err
	}

	newRefreshToken := utils.GenerateSHA256Token()
	user.RefreshToken = newRefreshToken
	if err := s.repository.Update(user); err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		UserID:       user.ID,
		RoleID:       user.RoleID,
	}, nil
}
