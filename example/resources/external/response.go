package external

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/config"
	"github.com/Pengenalan-Komputasi-ITB/oculi/response"
	"github.com/Pengenalan-Komputasi-ITB/oculi/validator"
)

func NewResponder(v validator.Validator, config *config.Env) response.Responder {
	return response.New(v, config.IsDevelopment())
}
