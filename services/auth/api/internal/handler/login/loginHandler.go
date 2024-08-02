package login

import (
	"net/http"

	"github.com/admgo/admgo/services/auth/api/internal/logic/login"
	"github.com/admgo/admgo/services/auth/api/internal/svc"
	"github.com/admgo/admgo/services/auth/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			l.WriteUserTokenToHeader(w)
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
