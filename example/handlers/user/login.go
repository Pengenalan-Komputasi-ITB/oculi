package user

import (
	userDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/user"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (h *handler) Login(req request.ReqContext, item userDto.LoginRequest) (userDto.CredentialResponse, error) {
	user, token, err := h.domain.User.Login(req, item)
	if err != nil {
		return userDto.CredentialResponse{}, err
	}

	result := userDto.NewCredentialResponse(user, token)
	return result, nil
}
