package user

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/constants"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/user"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (h *handler) Check(req request.ReqContext) (user.UserResponse, error) {
	credentialsData := req.Identifier()
	if credentialsData.ID == 0 {
		return user.UserResponse{}, constants.ErrNotLoggedIn
	}
	userDataBuff, _ := h.resource.Json.Marshal(credentialsData.Metadata)
	var userData dao.User
	h.resource.Json.Unmarshal(userDataBuff, &userData)
	return user.NewUserResponse(userData), nil
}
