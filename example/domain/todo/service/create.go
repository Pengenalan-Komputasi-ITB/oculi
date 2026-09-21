package service

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	todoDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (s *service) Create(req request.ReqContext, item todoDto.CreateTodoRequest) (dao.Todo, error) {
	result, err := s.repository.Create(req, item.ToDAO(req))
	if err != nil {
		return dao.Todo{}, err
	}
	return result, nil
}
