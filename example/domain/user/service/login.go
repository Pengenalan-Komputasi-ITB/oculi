package service

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/common/model/dto/auth"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/constants"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dao"
	userDto "github.com/Pengenalan-Komputasi-ITB/oculi/example/model/dto/user"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (s *service) Login(req request.ReqContext, item userDto.LoginRequest) (dao.User, string, error) {
	user, err := s.repository.GetByUsername(req, item.Username)
	if err != nil {
		return dao.User{}, "", err
	}

	if errPassword := s.hash.Verify(item.Password, user.Password); errPassword != nil {
		return dao.User{}, "", constants.ErrWrongPassword
	}
	token, err := s.resource.Tokenizer.
		CreateAccessAndEncode(
			auth.StandardCredentials{
				ID:       user.ID,
				Metadata: user,
			},
			s.resource.Config.JWTExp,
		)

	if err != nil {
		return dao.User{}, "", err
	}
	return user, token, nil
}
