package external

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/validator"
	v10 "github.com/Pengenalan-Komputasi-ITB/oculi/validator/v10"
)

func NewValidator() (validator.Validator, error) {
	return v10.New()
}
