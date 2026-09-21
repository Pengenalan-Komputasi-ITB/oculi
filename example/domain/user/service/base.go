package service

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/domain/user/repository"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	userDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/user"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/resources"
	"github.com/Pengenalan-Komputasi-ITB/oculi/hash"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

type (
	Service interface {
		Login(req request.ReqContext, item userDto.LoginRequest) (user dao.User, token string, err error)
		Register(req request.ReqContext, user userDto.RegisterRequest) error
	}

	service struct {
		resource   resources.Resource
		repository repository.Repository
		hash       hash.Hash
	}
)

func New(r resources.Resource, repo repository.Repository, h hash.Hash) Service {
	return &service{
		resource:   r,
		repository: repo,
		hash:       h,
	}
}
