package domainvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)

func validateXID(fl validator.FieldLevel) bool {
	xidString := fl.Field().String()
	_, err := xid.FromString(xidString)

	return err != nil
}
