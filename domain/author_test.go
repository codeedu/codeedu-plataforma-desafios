package domain_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewAuthor(t *testing.T) {

	name := faker.Name()
	email := faker.Email()

	author, err := domain.NewAuthor(name, email)

	require.Nil(t, err)
	require.Equal(t, author.Name, name)
	require.Equal(t, author.Email, email)

	email = "invalid-email"

	author, err = domain.NewAuthor(name, email)
	require.NotNil(t, err)
	require.Nil(t, author)

	author, err = domain.NewAuthor("", faker.Email())
	require.NotNil(t, err)
	require.Nil(t, author)
}
