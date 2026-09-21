package repository

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	"github.com/Pengenalan-Komputasi-ITB/oculi/logs"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (r *repository) GetByID(req request.ReqContext, todoId uint64) (dao.Todo, error) {
	todo := dao.Todo{ID: todoId}
	if err := req.Transaction().
		Where("owner_id = ?", req.Identifier().ID).
		Take(&todo).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Todo.Repository.GetByID",
			logs.KeyValue("ID", todoId),
			logs.KeyValue("RequestIdentifier", req.Identifier()),
			logs.KeyValue("Error", err),
		))
		return dao.Todo{}, err
	}
	return todo, nil
}
