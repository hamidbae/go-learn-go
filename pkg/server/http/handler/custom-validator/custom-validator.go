package customvalidator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsISO8601Date(fl validator.FieldLevel) bool {
	ISO8601DateRegexString := "^\\d{4}[\\/\\-](0?[1-9]|1[012])[\\/\\-](0?[1-9]|[12][0-9]|3[01])$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(fl.Field().String())
}
