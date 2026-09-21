package external

import (
	"github.com/Pengenalan-Komputasi-ITB/oculi/encoding"
	"github.com/Pengenalan-Komputasi-ITB/oculi/encoding/jsoniter"
)

func NewJsonEncoding() encoding.Encoding {
	return jsoniter.New()
}
