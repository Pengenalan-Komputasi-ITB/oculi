package service

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/domain/todo/repository"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	todoDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/todo"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/resources"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

type (
	Service interface {
		Create(req request.ReqContext, todo todoDto.CreateTodoRequest) (dao.Todo, error)
		Done(req request.ReqContext, todoId uint64) error
		Undone(req request.ReqContext, todoId uint64) error
		Edit(req request.ReqContext, todo todoDto.UpdateTodoRequest) error
		Delete(req request.ReqContext, todoId uint64) error
		GetAllByOwner(req request.ReqContext) ([]dao.Todo, error)
	}

	service struct {
		resource   resources.Resource
		repository repository.Repository
	}
)

func New(r resources.Resource, repo repository.Repository) Service {
	return &service{
		resource:   r,
		repository: repo,
	}
}
