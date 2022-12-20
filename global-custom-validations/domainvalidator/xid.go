package domainvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)

func validateXID(fl validator.FieldLevel) bool {
	xidString := fl.Field().String()

	if _, err := xid.FromString(xidString); err != nil {
		return false
	}

	return true
}
