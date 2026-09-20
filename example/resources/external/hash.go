package external

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/hash"
	"github.com/Pengenalan-Komputasi-ITB/oculi/hash/bcrypt"
)

func NewHash() (hash.Hash, error) {
	return bcrypt.NewHash()
}
