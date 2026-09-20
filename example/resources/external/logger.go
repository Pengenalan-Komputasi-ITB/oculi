package external

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/config"
	"github.com/Pengenalan-Komputasi-ITB/oculi/logs"
	"github.com/Pengenalan-Komputasi-ITB/oculi/logs/zap"
	z "go.uber.org/zap"
)

func NewLogger(config *config.Env) (logs.Logger, error) {
	return zap.New(config.IsDevelopment(), zap.Option{
		Level:  logs.GetLoggerLevel(config.LogLevel),
		Prefix: "",
	}, z.AddStacktrace(z.ErrorLevel), z.AddCallerSkip(1))
}
