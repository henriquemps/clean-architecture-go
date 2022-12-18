package repository

import (
	"gorm.io/gorm"
	"log"
	"project-name-here/boot/database"
	"project-name-here/entity"
)

type UserRepository struct {
	db *gorm.DB
}

var repository *UserRepository

func NewUserRepository() *UserRepository {

	if repository != nil {
		return repository
	}

	repository = &UserRepository{
		db: database.GetConnection(),
	}

	return repository
}

func (r *UserRepository) Create(user *entity.User) (*entity.User, error) {

	result := r.db.Create(user)

	if result.Error != nil {
		log.Printf("error when trying to create user \n details: %s", result.Error)

		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) Show(id int) (*entity.User, error) {

	var user *entity.User

	result := r.db.First(&user, id)

	if result.Error != nil {
		log.Printf("error when trying to show user \n details: %s", result.Error)

		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) List() ([]entity.User, error) {

	var users []entity.User

	result := r.db.Find(&users)

	if result.Error != nil {
		log.Printf("error when trying to list users \n details: %s", result.Error)

		return nil, result.Error
	}

	return users, nil
}
