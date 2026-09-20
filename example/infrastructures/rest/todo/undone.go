package todo

import (
	"github.com/labstack/echo/v4"
	oculiContext "github.com/Pengenalan-Komputasi-ITB/oculi/context"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/constants"
	request "github.com/Pengenalan-Komputasi-ITB/oculi/request/echo"
)

func (c *Controller) Undone(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database).Param("id")

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.Todo.Undone, req),
		nil,
		constants.TodoMappers,
	)

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
