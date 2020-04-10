package domain_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewChallengeFile(t *testing.T) {

	fileName := "text.txt"
	fileURL := faker.URL() + "/" + fileName

	cfile, err := domain.NewChallengeFile(fileName, fileURL)

	require.Equal(t, cfile.Name, fileName)
	require.Nil(t, err)

	fileName = "text.txt"
	fileURL = "ivalid-url"

	cfile, err = domain.NewChallengeFile(fileName, fileURL)

	require.Nil(t, cfile)
	require.NotNil(t, err)

	fileName = ""
	fileURL = "ivalid-url"

	cfile, err = domain.NewChallengeFile(fileName, fileURL)

	require.Nil(t, cfile)
	require.NotNil(t, err)

}
