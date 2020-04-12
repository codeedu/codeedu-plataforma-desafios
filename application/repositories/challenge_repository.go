package repositories

import (
	"fmt"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/jinzhu/gorm"
)

type ChallengeRepository interface {
	Insert(challenge *domain.Challenge) (*domain.Challenge, error)
	Find(id string) (*domain.Challenge, error)
	FindAll(order string) ([]domain.Challenge, error)
}

type ChallengeRepositoryDb struct {
	Db *gorm.DB
}

func (repo ChallengeRepositoryDb) Insert(challenge *domain.Challenge) (*domain.Challenge, error) {

	err := repo.Db.Create(challenge).Error

	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func (repo ChallengeRepositoryDb) Find(id string) (*domain.Challenge, error) {

	var challenge domain.Challenge
	repo.Db.Preload("ChallengeFiles").First(&challenge, "id = ?", id)

	if challenge.ID == "" {
		return nil, fmt.Errorf("Challenge does not exist")
	}

	return &challenge, nil
}

func (repo ChallengeRepositoryDb) FindAll(order string) ([]domain.Challenge, error) {

	var challenges []domain.Challenge
	repo.Db.Preload("ChallengeFiles").Find(&challenges).Order(order)

	return challenges, nil
}
