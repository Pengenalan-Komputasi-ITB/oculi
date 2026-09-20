package health

import (
	consts "github.com/Pengenalan-Komputasi-ITB/oculi/constant/key"
	"github.com/Pengenalan-Komputasi-ITB/oculi/example/constants"
	"github.com/Pengenalan-Komputasi-ITB/oculi/request"
)

func (h *handler) Reset(ctx request.ReqContext) error {
	if ctx.GetOrDefault(consts.QueryPrefix("key"), "").(string) != h.resource.Config.DatabaseResetKey {
		return constants.ErrResetUnauthorized
	}
	h.resource.DBManager.Reset()
	h.resource.DBManager.Install()
	return nil
}
