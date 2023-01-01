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

func GetUser(input GetUserInput) (*User, error) {
	if err := domainvalidator.ValidateStruct(input); err != nil {
		return nil, err
	}

	// Call to repository to get user
	return &User{}, nil
}
