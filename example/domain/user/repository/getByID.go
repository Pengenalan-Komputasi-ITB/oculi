package repository

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	"github.com/Pengenalan-Komputasi-ITB/oculi/logs"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (r *repository) GetByID(req request.ReqContext, userId uint64) (dao.User, error) {
	user := dao.User{ID: userId}
	if err := req.Transaction().
		Take(&user).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"User.Repository.GetByID",
			logs.KeyValue("ID", userId),
			logs.KeyValue("Error", err),
		))
		return dao.User{}, err
	}
	return user, nil
}
