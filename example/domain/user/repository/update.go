package repository

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/common/model/dto"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	"github.com/Pengenalan-Komputasi-ITB/oculi/logs"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (r *repository) Update(req request.ReqContext, userId uint64, request dto.Map) error {
	if err := req.Transaction().Model(dao.User{}).
		Where("id = ?", userId).
		Updates(request.ToMap()).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"User.Repository.Update",
			logs.KeyValue("ID", userId),
			logs.KeyValue("RequestMap", request.ToMap()),
			logs.KeyValue("Error", err),
		))
		return err
	}
	return nil
}
