package rest

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/infrastructures/rest/health"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/infrastructures/rest/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/infrastructures/rest/user"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/resources"
	"go.uber.org/dig"
)

type (
	Rest struct {
		dig.In

		Controller Controller
		Resource   resources.Resource
	}

	Controller struct {
		dig.In

		Health health.Controller
		User   user.Controller
		Todo   todo.Controller
	}
)
