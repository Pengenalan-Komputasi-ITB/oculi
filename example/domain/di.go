package domain

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/di"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/domain/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/domain/user"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).
		Register(user.Register).
		Register(todo.Register).
		Proceed()
}
