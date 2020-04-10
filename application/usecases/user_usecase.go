package usecases

import (
	"github.com/codeedu/codeedu-plataforma-desafios/application/repositories"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"log"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *UserUseCase) Create(user *domain.User) (*domain.User, error) {

	user, err := u.UserRepository.Insert(user)

	if err != nil {
		log.Fatalf("Error to persist new user: %v", err)
		return user, err
	}

	return user, nil
}