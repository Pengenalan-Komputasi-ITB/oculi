package user

import (
	userDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/user"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (h *handler) Register(req request.ReqContext, item userDto.RegisterRequest) error {
	return h.domain.User.Register(req, item)
}
