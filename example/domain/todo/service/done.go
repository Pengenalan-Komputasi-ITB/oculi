package service

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/constants"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (s *service) Done(req request.ReqContext, todoId uint64) error {
	t, err := s.repository.GetByID(req, todoId)
	if err != nil {
		return err
	}

	if t.IsDone {
		return constants.ErrTodoAlreadyDoneState
	}

	return s.repository.Update(req, todoId, map[string]interface{}{
		"is_done": true,
	})
}
