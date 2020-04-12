package repositories_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/codeedu/codeedu-plataforma-desafios/application/repositories"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/codeedu/codeedu-plataforma-desafios/framework/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChallengeRepositoryDb_Insert(t *testing.T) {
	db := utils.ConnectDB("test")
	defer db.Close()

	challenge := domain.NewChallenge()
	challenge.Title = faker.Word()
	challenge.Slug = faker.Word()
	challenge.Description = faker.Sentence()
	challenge.Tags = faker.Word()
	challenge.Requirements = faker.Sentence()
	challenge.Level = faker.Word()

	require.Nil(t, challenge.Valid())

	repo := repositories.ChallengeRepositoryDb{Db: db}
	repo.Insert(challenge)

	c, err := repo.Find(challenge.ID)
	require.NotEmpty(t, c.ID)
	require.Equal(t, c.ID, challenge.ID)
	require.Nil(t, err)
}

func TestChallengeRepositoryDb_InsertWithChallengeFiles(t *testing.T) {
	db := utils.ConnectDB("test")
	defer db.Close()

	challenge := domain.NewChallenge()
	challenge.Title = faker.Word()
	challenge.Slug = faker.Word()
	challenge.Description = faker.Sentence()
	challenge.Tags = faker.Word()
	challenge.Requirements = faker.Sentence()
	challenge.Level = faker.Word()
	require.Nil(t, challenge.Valid())

	file1, err := domain.NewChallengeFile("test.txt", "http://google.com/test.txt")
	file1.ChallengeID = challenge.ID
	require.Nil(t, file1.Valid())

	file2, err := domain.NewChallengeFile("test2.txt", "http://google.com/test2.txt")
	file2.ChallengeID = challenge.ID
	require.Nil(t, file2.Valid())

	var files []*domain.ChallengeFile
	files = append(files, file1)
	files = append(files, file2)

	challenge.ChallengeFiles = files

	repo := repositories.ChallengeRepositoryDb{Db: db}
	repo.Insert(challenge)

	c, err := repo.Find(challenge.ID)
	require.NotEmpty(t, c.ID)
	require.Equal(t, c.ID, challenge.ID)
	require.Nil(t, err)

	require.Equal(t, 2, len(c.ChallengeFiles))
}
