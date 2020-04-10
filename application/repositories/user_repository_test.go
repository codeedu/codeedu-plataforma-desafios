package repositories_test

import (
	"github.com/asaskevich/govalidator"
	"github.com/bxcodec/faker/v3"
	"github.com/codeedu/codeedu-plataforma-desafios/application/repositories"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/codeedu/codeedu-plataforma-desafios/framework/utils"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var db *gorm.DB

func TestUserRepositoryDb_Find(t *testing.T) {
	db = utils.ConnectDB()
	db.LogMode(true)

	repo := repositories.UserRepositoryDb{Db: db}
	newUser, err := domain.NewUser(faker.Name(), faker.Email(), "123")

	if err != nil {
		log.Fatalf("%v", err)
	}
	repo.Insert(newUser)

	res, _ := repo.Find(newUser.Email)

	require.NotEmpty(t, res.Token)
	require.True(t, govalidator.IsUUIDv4(res.Token))

}
