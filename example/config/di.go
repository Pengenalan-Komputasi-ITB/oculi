package config

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/di"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.SimpleRegistrant(c, NewConfig)
}
