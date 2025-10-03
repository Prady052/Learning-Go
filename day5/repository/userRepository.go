package repository

import (
	"errors"
	"fmt"

	"cdac.com/day5/models"
	//go get -u gorm.io/gorm
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
	//create return db pointer
	return r.db.Create(user).Error
}

func (r *userRepo) GetByID(id uint) (*models.User, error) {
	var u models.User
	err := r.db.First(&u, id).Error
	if err != nil {
		// checking the record not found error in error tree
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Update(user *models.User) (*models.User, error) {
	// Save updates value in database.
	// If value doesn't contain a matching primary key, value is inserted.
	fmt.Print("<-------------------->", user)
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
