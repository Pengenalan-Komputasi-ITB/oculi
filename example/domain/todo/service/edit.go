package service

import (
	todoDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (s *service) Edit(req request.ReqContext, item todoDto.UpdateTodoRequest) error {
	_, err := s.repository.GetByID(req, item.ID)
	if err != nil {
		return err
	}

	return s.repository.Update(req, item.ID, item.ToUpdateMapRequest())
}
