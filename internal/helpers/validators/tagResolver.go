package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "gt":
		return fmt.Sprintf("%s must be greater than 0", fe.Field())
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	}

	return fe.Error() // default error
}
