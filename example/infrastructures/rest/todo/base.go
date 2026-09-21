package todo

import (
	"github.com/labstack/echo/v4"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/handlers"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/resources"
	"github.com/Pengenalan-Komputasi-ITB/oculi/middleware/token"
	"go.uber.org/dig"
)

type (
	Controller struct {
		dig.In

		Handler  handlers.Handler
		Resource resources.Resource
	}
)

func (c Controller) Register(ec *echo.Group) error {
	g := ec.Group("/todo")
	private := token.PrivateEndpoint(c.Resource.Responder)
	g.GET("", c.Get, private)
	g.POST("", c.Create, private)
	g.PATCH("", c.Edit, private)
	g.DELETE("/:id", c.Delete, private)
	g.PATCH("/:id/done", c.Done, private)
	g.PATCH("/:id/undone", c.Undone, private)
	return nil
}
