package token

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/Pengenalan-Komputasi-ITB/oculi/common/model/dto/auth"
	errConsts "github.com/Pengenalan-Komputasi-ITB/oculi/constant/errors"
	consts "github.com/Pengenalan-Komputasi-ITB/oculi/constant/key"
	"github.com/Pengenalan-Komputasi-ITB/oculi/context"
	errorUtil "github.com/Pengenalan-Komputasi-ITB/oculi/errors"
	"github.com/Pengenalan-Komputasi-ITB/oculi/response"
	"github.com/Pengenalan-Komputasi-ITB/oculi/token"
)

type (
	ClaimFunction func(req *http.Request) (token.Claims, error)
)

var (
	tokenMiddlewareActivated = false
)

func fetchAuthentication(cf ClaimFunction, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, ok := c.(*context.Context)
		if !ok {
			ctx = context.New(c)
		}
		claims, err := cf(ctx.Request())
		if err != nil {
			ctx.Set(consts.KeyCredentials, auth.StandardCredentials{
				ID:       0,
				Metadata: errorUtil.InjectDetails(errConsts.ErrUnauthorized, err.Error()),
			})
		} else {
			ctx.Set(consts.KeyCredentials, claims.Credentials())
		}
		return next(c)
	}
}

func EchoMiddleware(cf ClaimFunction) echo.MiddlewareFunc {
	tokenMiddlewareActivated = true
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return fetchAuthentication(cf, next)
	}
}

func skip(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func sendErrorIfNeeded(next echo.HandlerFunc, r response.Responder) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*context.Context)
		credentials := ctx.Get(consts.KeyCredentials)
		if credentials == nil {
			ctx.AddError(
				http.StatusUnauthorized,
				errorUtil.InjectDetails(
					errConsts.ErrUnauthorized,
					"unauthorized: credentials not found",
				),
			)
			return r.NewJSONResponse(ctx, nil, nil)
		} else {
			cdto := credentials.(auth.StandardCredentials)
			if cdto.ID == 0 {
				ctx.AddError(http.StatusUnauthorized, cdto.Metadata.(error))
				return r.NewJSONResponse(ctx, nil, nil)
			}
		}
		return next(c)
	}
}

func PrivateEndpoint(r response.Responder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		if !tokenMiddlewareActivated {
			return skip(next)
		}

		return sendErrorIfNeeded(next, r)
	}
}
