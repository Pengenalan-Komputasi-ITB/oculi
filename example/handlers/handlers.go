package handlers

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/handlers/health"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/handlers/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/handlers/user"
	"go.uber.org/dig"
)

type (
	Handler struct {
		dig.In

		Health health.Handler
		User   user.Handler
		Todo   todo.Handler
	}
)
