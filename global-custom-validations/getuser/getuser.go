package getuser

import (
	"github.com/mgreg90/govalidator-custom-validations-example/global-custom-validations/domainvalidator"
)

type GetUserInput struct {
	ID string `validate:"xid"`
}

type User struct {
	ID                  string
	Email               string
	DateOfBirth         string
	ConsentsToReminders bool
	PhoneNumber         *string
}

func GetUser(input GetUserInput) (error, *User) {
	if err := domainvalidator.ValidateStruct(input); err != nil {
		return err, nil
	}

	// Call to repository to get user
	return nil, &User{}
}
