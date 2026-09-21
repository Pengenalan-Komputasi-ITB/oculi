package resources

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/Pengenalan-Komputasi-ITB/oculi/docs"
	"github.com/Pengenalan-Komputasi-ITB/oculi/encoding"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/config"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/resources/external"
	"github.com/Pengenalan-Komputasi-ITB/oculi/logs"
	"github.com/Pengenalan-Komputasi-ITB/oculi/persistent/sql"
	"github.com/Pengenalan-Komputasi-ITB/oculi/response"
	"github.com/Pengenalan-Komputasi-ITB/oculi/token"
	"github.com/Pengenalan-Komputasi-ITB/oculi/validator"
	"go.uber.org/dig"
)

var (
	once       sync.Once
	identifier string
	uptime     time.Time
)

type (
	Resource struct {
		dig.In

		EchoData      *echo.Echo
		Log           logs.Logger
		Responder     response.Responder
		Config        *config.Env
		ValidatorData validator.Validator
		Database      sql.API
		Tokenizer     token.Tokenizer
		Documentation docs.Documentation
		DBManager     *external.DBManager
		Json          encoding.Encoding
	}
)

func (r Resource) Echo() *echo.Echo {
	return r.EchoData
}
func (r Resource) ServiceName() string {
	return r.Config.ServiceName
}
func (r Resource) ServerPort() int {
	return r.Config.ServerPort
}

func (r Resource) Identifier() string {
	once.Do(func() {
		uptime = time.Now()
		if identifier == "" {
			identifier = r.Config.ServiceName + " " + uptime.String()
		}
	})
	return identifier
}

func (r Resource) Uptime() time.Time {
	once.Do(func() {
		uptime = time.Now()
		if identifier == "" {
			identifier = r.Config.ServiceName + " " + uptime.String()
		}
	})
	return uptime
}

func (r Resource) ServerGracefullyDuration() time.Duration {
	return r.Config.GracefullyDuration
}

func (r Resource) Logger() logs.Logger {
	return r.Log
}

func (r Resource) Validator() validator.Validator {
	return r.ValidatorData
}

func (r Resource) Close() error {
	var errMessage = make([]string, 0)

	if err := r.EchoData.Close(); err != nil {
		errMessage = append(errMessage, err.Error())
	}

	if len(errMessage) > 0 {
		return errors.New(strings.Join(errMessage, "\n"))
	}
	return nil
}
