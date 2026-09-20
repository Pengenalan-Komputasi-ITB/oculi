package todo

import (
	todoDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (h *handler) Create(req request.ReqContext, item todoDto.CreateTodoRequest) (todoDto.TodoResponse, error) {
	todo, err := h.domain.Todo.Create(req, item)
	if err != nil {
		return todoDto.TodoResponse{}, err
	}
	return todoDto.NewTodoResponse(todo), nil
}
