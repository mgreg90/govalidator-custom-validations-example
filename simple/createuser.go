package simple

import (
	"github.com/go-playground/validator/v10"
)

type CreateUserInput struct {
	Email               string `validate:"email"`
	DateOfBirth         string `validate:"datetime=01/02/2006"`
	ConsentsToReminders bool
	PhoneNumber         *string `validate:"required_if=ConsentsToReminders true,omitempty,e164"`
}

type User struct {
	ID string
	CreateUserInput
}

func CreateUser(input CreateUserInput) (error, *User) {
	if err := validator.New().Struct(input); err != nil {
		return err, nil
	}

	// Call to repository to create user
	return nil, &User{}
}
