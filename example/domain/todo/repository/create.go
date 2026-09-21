package repository

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	"github.com/Pengenalan-Komputasi-ITB/oculi/logs"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (r *repository) Create(req request.ReqContext, item dao.Todo) (dao.Todo, error) {
	if err := req.Transaction().
		Create(&item).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Todo.Repository.Create",
			logs.KeyValue("Todo", item),
			logs.KeyValue("Error", err),
		))
		return dao.Todo{}, err
	}
	return item, nil
}
