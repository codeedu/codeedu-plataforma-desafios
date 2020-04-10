package repositories

import (
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {

	err := repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
