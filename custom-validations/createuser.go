package customvalidations

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type CreateUserInput struct {
	Email                        string `validate:"email"`
	DateOfBirth                  string `validate:"datetime=01/02/2006"`
	ConsentsToReminders          bool
	PhoneNumber                  *string `validate:"required_if=ConsentsToReminders true,omitempty,e164"`
	RegistrationConfirmationTime time.Time
}

type User struct {
	ID string
	CreateUserInput
}

func CreateUser(input CreateUserInput) (error, *User) {
	if err := input.validate(); err != nil {
		return err, nil
	}

	// Call to repository to create user
	return nil, &User{}
}

func (i CreateUserInput) validate() error {
	v := validator.New()
	v.RegisterStructValidation(i.validateRegistrationConfirmationTime, CreateUserInput{})

	return v.Struct(i)
}

func (i CreateUserInput) validateRegistrationConfirmationTime(sl validator.StructLevel) {
	now := time.Now()

	if i.RegistrationConfirmationTime.After(now) {
		sl.ReportError(i.RegistrationConfirmationTime, "RegistrationConfirmationTime", "RegistrationConfirmationTime", "can not be in the future", "")
	}
}
