package handlers

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/di"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/handlers/health"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/handlers/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/handlers/user"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	return di.NewRegistrant(container).
		Provide(health.NewHandler).
		Provide(user.NewHandler).
		Provide(todo.NewHandler).
		Proceed()
}
