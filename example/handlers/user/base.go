package user

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/domain"
	userDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/user"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/resources"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

type (
	handler struct {
		domain   domain.Domain
		resource resources.Resource
	}

	Handler interface {
		Login(req request.ReqContext, item userDto.LoginRequest) (userDto.CredentialResponse, error)
		Register(req request.ReqContext, item userDto.RegisterRequest) error
		Check(req request.ReqContext) (userDto.UserResponse, error)
	}
)

func NewHandler(domain domain.Domain, resource resources.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
