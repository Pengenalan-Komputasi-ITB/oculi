package repository

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	"github.com/Pengenalan-Komputasi-ITB/oculi/logs"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (r *repository) Create(req request.ReqContext, user dao.User) (dao.User, error) {
	if err := req.Transaction().Create(&user).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"User.Repository.Create",
			logs.KeyValue("User", user),
			logs.KeyValue("Error", err),
		))
		return dao.User{}, err
	}
	return user, nil
}
