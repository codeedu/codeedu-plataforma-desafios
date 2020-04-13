package usecases_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/codeedu/codeedu-plataforma-desafios/application/repositories"
	"github.com/codeedu/codeedu-plataforma-desafios/application/usecases"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/codeedu/codeedu-plataforma-desafios/framework/utils"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"testing"
)

var db *gorm.DB

func TestLogin_Auth(t *testing.T) {
	db = utils.ConnectDB("test")

	repo := repositories.UserRepositoryDb{Db: db}
	email := faker.Email()
	password := faker.Password()

	newUser, err := domain.NewUser(faker.Name(), email, password)
	repo.Insert(newUser)

	userUserCase := usecases.UserUseCase{UserRepository: repo}
	_, err = userUserCase.Auth(email, password)
	require.Nil(t, err)
}
