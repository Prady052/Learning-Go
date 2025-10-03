package service

import (
	"errors"
	"fmt"
	"strings"

	"cdac.com/day5/models"
	"cdac.com/day5/repository"
)

var (
	ErrNotFound    = errors.New("user not found")
	ErrInvalidData = errors.New("invalid data")
	departments    = map[string]int{
		"IT":        101,
		"HR":        102,
		"FINANCE":   103,
		"MARKETING": 105,
		"SUPPORT":   106,
		"R&D":       107,
	}
)

type UserService interface {
	CreateUser(u *models.User) error
	GetUser(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(id uint, input *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

// struct that implements UserService interface
type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(u *models.User) error {
	err := validateDetails(u)
	if err != nil {
		return err
	}
	return s.repo.Create(u)
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, ErrNotFound
	}
	return u, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) UpdateUser(id uint, input *models.User) (*models.User, error) {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, ErrNotFound
	}
	err = validateDetails(input)
	if err != nil {
		return nil, err
	}
	input.ID = u.ID
	input.CreatedAt = u.CreatedAt
	u, err = s.repo.Update(input)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *userService) DeleteUser(id uint) error {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if u == nil {
		return ErrNotFound
	}
	return s.repo.Delete(id)
}

func validateDetails(u *models.User) error {
	fmt.Println(u)
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	if u.Name == "" || u.Email == "" {
		return ErrInvalidData
	}
	u.Deparment = strings.ToUpper(u.Deparment)
	if departments[u.Deparment] == 0 {
		return ErrInvalidData
	}
	return nil
}
