package external

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/config"
	"github.com/Pengenalan-Komputasi-ITB/oculi/token"
	oculiJWT "github.com/Pengenalan-Komputasi-ITB/oculi/token/jwt"
)

func NewTokenizer(config *config.Env) token.Tokenizer {
	identifier := oculiJWT.GenerateIdentifier(config.ServiceState, 5, config.ServiceName)
	return oculiJWT.New(config.JWTKey, jwt.SigningMethodHS256.Name, identifier)
}
