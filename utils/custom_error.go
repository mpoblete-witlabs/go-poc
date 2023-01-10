package utils

import (
	"github.com/go-playground/validator/v10"
)

func MsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Este campo es obligatorio"
	case "email":
		return "Debe ser formato Mail"
	case "gte":
		return "Valor menor al permitido"
	case "lte":
		return "Valor mayor al permitido"
	}

	return fe.Error()
}
