package handler

import (
	"net/http"

	"github.com/admgo/admgo/services/storager/api/internal/logic"
	"github.com/admgo/admgo/services/storager/api/internal/svc"
	"github.com/admgo/admgo/services/storager/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StoragerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewStoragerLogic(r.Context(), svcCtx)
		resp, err := l.Storager(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
