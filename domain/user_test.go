package domain_test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
)

func TestNewUser(t *testing.T) {

	t.Parallel()

	var scenarios = []struct {
		Descr               string
		GivenUsername       string
		GivenEmail          string
		GivenPassword       string
		ShouldReturnAnError bool
	}{
		{
			Descr:               "Should return an error when an invalid username (empty string) is provided",
			GivenUsername:       "",
			GivenEmail:          "dearrudam@gmail.com",
			GivenPassword:       "aa11@@",
			ShouldReturnAnError: true,
		},
		{
			Descr:               "Should return an error when an invalid username (blank string) is provided",
			GivenUsername:       " ",
			GivenEmail:          "dearrudam@gmail.com",
			GivenPassword:       "aa11@@",
			ShouldReturnAnError: true,
		},
		{
			Descr:               "Should return an error when an invalid e-mail (invalid format) is provided",
			GivenUsername:       "Maximillian",
			GivenEmail:          "dearrudamxxxgmail.com",
			GivenPassword:       "aa11@@",
			ShouldReturnAnError: true,
		},
		{
			Descr:               "Should return an error when an invalid e-mail (empty string) is provided",
			GivenUsername:       "Maximillian",
			GivenEmail:          "",
			GivenPassword:       "aa11@@",
			ShouldReturnAnError: true,
		},
		{
			Descr:               "Should return an error when an invalid password (empty password) is provided",
			GivenUsername:       "Maximillian",
			GivenEmail:          "dearrudamxxxgmail.com",
			GivenPassword:       "",
			ShouldReturnAnError: true,
		},
		{
			Descr:               "Should return an error when an invalid password (length is less than 6 characters) is provided",
			GivenUsername:       "Maximillian",
			GivenEmail:          "dearrudamxxxgmail.com",
			GivenPassword:       "dscss",
			ShouldReturnAnError: true,
		},
		{
			Descr:               "Should return an error when an invalid password (it's missing 2 alphanumeric characters at least) is provided",
			GivenUsername:       "Maximillian",
			GivenEmail:          "dearrudamxxxgmail.com",
			GivenPassword:       "11@@",
			ShouldReturnAnError: true,
		},
		{
			Descr:               "Should return an error when an invalid password (it's missing 2 numeric characters at least) is provided",
			GivenUsername:       "Maximillian",
			GivenEmail:          "dearrudamxxxgmail.com",
			GivenPassword:       "aa@@",
			ShouldReturnAnError: true,
		},
		{
			Descr:               "Should return an error when an invalid password (it's missing 2 special characters (_@./#&+-) at least) is provided",
			GivenUsername:       "Maximillian",
			GivenEmail:          "dearrudamxxxgmail.com",
			GivenPassword:       "aa11",
			ShouldReturnAnError: true,
		},
		{
			Descr:               "A valid User reference must be returned when a valid username, valid email and valid password ( it must contains 2 alphanumeric characters, 2 numeric characters and 2 special characters (_@./#&+-) at least) is provided",
			GivenUsername:       "Maximillian",
			GivenEmail:          "dearrudam@gmail.com",
			GivenPassword:       "aa11!@",
			ShouldReturnAnError: false,
		},
	}

	for _, scenario := range scenarios {
		user, err := domain.NewUser(scenario.GivenUsername, scenario.GivenEmail, scenario.GivenPassword)
		if scenario.ShouldReturnAnError {
			if err == nil {
				t.Fatalf("%v: Error", scenario.Descr)
				continue
			}
		} else {
			if user == nil {
				t.Fatalf("%v: Error - Returned user cannot be a nil reference", scenario.Descr)
				continue
			}
			if !govalidator.IsUUIDv4(user.Token) {
				t.Fatalf("%v: Error - user.Token needs to be a valid UUID v4", scenario.Descr)
				continue
			}
		}
		t.Logf("%v: OK", scenario.Descr)
	}

	// _, err := domain.NewUser("Wesley", "wesleyxwesleyxxxgmail.com", "ab12!@")
	// require.Error(t, err)

	// _, err = domain.NewUser("Wesley", "wesleyxwesleyxxxgmail.com", "")
	// require.Error(t, err)

	// _, err = domain.NewUser("", "wesleyxwesleyxxxgmail.com", "")
	// require.Error(t, err)

	// user, err := domain.NewUser("Wesley", "wesleyxwesleyxxx@gmail.com", "ab12!@")
	// require.Nil(t, err)

	// user, err = domain.NewUser("Wesley", "wesleyxwesley@gmail.com", "1")
	// require.Nil(t, err)

	// govalidator.IsUUIDv4(user.Token)

}
