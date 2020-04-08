package usecases

import (
	"io/ioutil"
	"os"
)

type Login struct{}

func (login Login) Auth(email string, password string) string {
	token := "123"

	token, err := login.persistToken(token)
	if err != nil {
		panic(err)
	}
	return token
}

func (login Login) persistToken(token string) (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	path := home + "/.codeedu"
	err = os.Mkdir(path, os.ModePerm)

	if err != nil {
		panic(err)
	}

	content := []byte(token)
	err = ioutil.WriteFile(path+"/token", content, 0644)
	if err != nil {
		panic(err)
	}

	return token, nil
}
