package user

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/di"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/domain/user/repository"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/domain/user/service"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).
		Provide(repository.New).
		Provide(service.New).
		Proceed()
}
