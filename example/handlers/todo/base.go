package todo

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/domain"
	todoDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/resources"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

type (
	handler struct {
		domain   domain.Domain
		resource resources.Resource
	}

	Handler interface {
		Create(req request.ReqContext, item todoDto.CreateTodoRequest) (todoDto.TodoResponse, error)
		Done(req request.ReqContext) error
		Undone(req request.ReqContext) error
		Edit(req request.ReqContext, item todoDto.UpdateTodoRequest) error
		Delete(req request.ReqContext) error
		GetAllByOwner(req request.ReqContext) (todoDto.TodosResponse, error)
	}
)

func NewHandler(domain domain.Domain, resource resources.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
