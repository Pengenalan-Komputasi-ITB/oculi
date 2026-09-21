package todo

import (
	todoDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (h *handler) Edit(req request.ReqContext, item todoDto.UpdateTodoRequest) error {
	return h.domain.Todo.Edit(req, item)
}
