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
	err := user.Prepare()

	if err != nil {
		log.Fatal(err)
	}

	user, err = u.UserRepository.Insert(user)
	if err != nil {
		log.Fatalf("Error to create new user: %v", err)
	}

	return user, nil
}
