package shared

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(r any) error {
	err := validate.Struct(r)

	if err == nil {
		return nil
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string

		for _, fe := range validationErrors {
			messages = append(messages, formatError(fe))
		}

		return errors.New(strings.Join(messages, "; "))
	}

	return err
}

func formatError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", fe.Field(), fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fe.Field())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters long", fe.Field(), fe.Param())
	case "uuid4":
		return fmt.Sprintf("%s must be a valid UUIDv4", fe.Field())
	// add more cases as needed
	default:
		return fmt.Sprintf("%s is not valid (%s)", fe.Field(), fe.Tag())
	}
}
