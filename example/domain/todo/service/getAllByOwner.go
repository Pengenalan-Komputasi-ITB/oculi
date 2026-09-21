package service

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (s *service) GetAllByOwner(req request.ReqContext) ([]dao.Todo, error) {
	return s.repository.GetAllByOwner(req)
}
