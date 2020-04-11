package usecases

import (
	"github.com/codeedu/codeedu-plataforma-desafios/application/repositories"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
)

type ChallengeUseCase struct {
	ChallengeRepository repositories.ChallengeRepository
}

func (c *ChallengeUseCase) Create(challenge *domain.Challenge) (*domain.Challenge, error) {

	res, err := c.ChallengeRepository.Insert(challenge)

	if err != nil {
		return res, err
	}

	return res, nil
}
