package customvalidations

import (
	"testing"
	"time"
)

var (
	validPhone           = "+15555555555"
	validCreateUserInput = CreateUserInput{
		Email:                        "mike@gmail.com",
		PhoneNumber:                  &validPhone,
		DateOfBirth:                  "10/25/2000",
		ConsentsToReminders:          true,
		RegistrationConfirmationTime: time.Now().Add(-24 * time.Hour), // One day ago,
	}
)

func TestCreateUser_ValidFields(t *testing.T) {
	if err, _ := CreateUser(validCreateUserInput); err != nil {
		t.Fatal("Returned error for valid input", validCreateUserInput, err)
	}
}

func TestCreateUser_InvalidEmail(t *testing.T) {
	input := validCreateUserInput
	input.Email = "fake.com"

	if err, _ := CreateUser(input); err == nil {
		t.Fatal("Failed to return error on invalid email", input)
	}
}

func TestCreateUser_InvalidDateOfBirth(t *testing.T) {
	input := validCreateUserInput
	input.DateOfBirth = "2000-10-25"

	if err, _ := CreateUser(input); err == nil {
		t.Fatal("Failed to return error on invalid date of birth", input)
	}
}

func TestCreateUser_InvalidPhoneNumber(t *testing.T) {
	input := validCreateUserInput
	ph := "(305)555-5555"
	input.PhoneNumber = &ph

	if err, _ := CreateUser(input); err == nil {
		t.Fatal("Failed to return error on invalid phone number", input)
	}
}

func TestCreateUser_InvalidConsentWithoutPhoneNumber(t *testing.T) {
	input := validCreateUserInput
	input.PhoneNumber = nil

	if err, _ := CreateUser(input); err == nil {
		t.Fatal("Failed to return error on invalid phone number", input)
	}
}

func TestCreateUser_ValidNilPhoneNumberWithoutConsent(t *testing.T) {
	input := validCreateUserInput
	input.ConsentsToReminders = false
	input.PhoneNumber = nil

	if err, _ := CreateUser(input); err != nil {
		t.Fatal("Returned an error when consent is false and phone is nil", input, err)
	}
}

func TestCreateUser_InvalidRegistrationConfirmationTime(t *testing.T) {
	input := validCreateUserInput
	input.RegistrationConfirmationTime = time.Now().Add(24 * time.Hour) // One day from now

	if err, _ := CreateUser(input); err == nil {
		t.Fatal("Failed to return error with registration time in the future", input, err)
	}
}
