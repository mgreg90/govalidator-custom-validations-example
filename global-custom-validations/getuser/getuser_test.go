package getuser

import (
	"github.com/rs/xid"
	"testing"
)

var (
	validGetUserInput = GetUserInput{
		ID: xid.New().String(),
	}
)

func TestGetUser_ValidFields(t *testing.T) {
	if _, err := GetUser(validGetUserInput); err != nil {
		t.Fatal("Returned error for valid input", validGetUserInput, err)
	}
}

func TestGetUser_InvalidXID(t *testing.T) {
	input := validGetUserInput
	input.ID = "something-invalid"

	if _, err := GetUser(input); err == nil {
		t.Fatal("Failed to return error on invalid xid", input)
	}
}
