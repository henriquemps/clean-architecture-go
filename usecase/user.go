package usecase

import (
	"log"
	"project-name-here/api/presenter"
	"project-name-here/entity"
	"project-name-here/infrastucture/repository"
	"project-name-here/models"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repository: repository.NewUserRepository(),
	}
}

func (s *UserService) CreateUser(model models.User) (presenter.User, error) {

	user := &entity.User{
		Email:     model.Email,
		Password:  model.Password,
		FirstName: model.FirstName,
		LastName:  model.LastName,
	}

	created, err := s.repository.Create(user)

	if err != nil {
		log.Printf("error when trying create user \n details: %s", err)

		return presenter.User{}, err
	}

	return presenter.User{
		ID:        created.ID,
		Email:     created.Email,
		FirstName: created.FirstName,
		LastName:  created.LastName,
	}, nil
}

func (s *UserService) ShowUser(id int) (presenter.User, error) {

	user, err := s.repository.Show(id)

	if err != nil {
		log.Printf("error when trying request to list detail user \n details: %s", err)

		return presenter.User{}, err
	}

	return presenter.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (s *UserService) ListUsers() ([]presenter.User, error) {

	var users []presenter.User

	list, err := s.repository.List()

	if err != nil {
		log.Printf("error when trying request to list users \n details: %s", err)

		return []presenter.User{}, err
	}

	for _, user := range list {
		users = append(users, presenter.User{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return users, nil
}
