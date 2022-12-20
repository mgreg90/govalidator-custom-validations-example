package domainvalidator

import (
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()

	v.RegisterValidation("xid", validateXID)
}

func ValidateStruct(i interface{}) error {
	return v.Struct(i)
}
