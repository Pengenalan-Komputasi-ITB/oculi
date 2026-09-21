package todo

import (
	todoDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (h *handler) GetAllByOwner(req request.ReqContext) (todoDto.TodosResponse, error) {
	todos, err := h.domain.Todo.GetAllByOwner(req)
	if err != nil {
		return todoDto.TodosResponse{}, err
	}
	return todoDto.NewTodosResponse(todos), nil
}
