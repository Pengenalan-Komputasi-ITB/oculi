package app

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/app"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/di"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/infrastructures"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/resources"
	mw "github.com/Pengenalan-Komputasi-ITB/oculi/middleware/token"
	"github.com/Pengenalan-Komputasi-ITB/oculi/server"

	webserver "github.com/Pengenalan-Komputasi-ITB/oculi/server/echo"
	"go.uber.org/dig"
)

func Run() {
	invoker := func(container *dig.Container) error {
		return container.Invoke(func(i infrastructures.Component, r resources.Resource) error {
			s := webserver.New(i, r)
			if r.Config.IsDevelopment() {
				s.DevelopmentMode()
			}
			s.BeforeRun(func(res server.Resource) error {
				r := res.(resources.Resource)
				res.Echo().Use(mw.EchoMiddleware(r.Tokenizer.DecodeAccessHeader))
				return nil
			})
			return s.Run()
		})
	}

	if err := app.Run(di.Container, invoker); err != nil {
		panic(err)
	}
}
