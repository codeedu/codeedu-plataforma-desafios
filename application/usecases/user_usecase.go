package usecases

import (
	"fmt"
	"github.com/codeedu/codeedu-plataforma-desafios/application/repositories"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *UserUseCase) Create(user *domain.User) (*domain.User, error) {

	user, err := u.UserRepository.Insert(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserUseCase) Auth(email string, password string) (*domain.User, error) {

	user, err := u.UserRepository.Find(email)

	if err != nil {
		return nil, fmt.Errorf("The password is invalid for the user: %v", email)
	}

	if user.IsCorrectPassword(password) {
		return user, nil
	}

	return nil, fmt.Errorf("The password is invalid for the user: %v", email)

}
