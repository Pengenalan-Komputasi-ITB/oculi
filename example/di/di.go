package di

import (
	"sync"

	"github.com/Pengenalan-Komputasi-ITB/oculi/di"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/config"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/domain"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/handlers"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/resources"
	"go.uber.org/dig"
)

var (
	container *dig.Container
	once      sync.Once
)

func Container() (*dig.Container, error) {
	items := []di.Registerable{
		config.Register,
		resources.Register,
		domain.Register,
		handlers.Register,
	}
	return di.Container(items)(&once, container)
}
