package middleware

import (
	"context"
	"github.com/admgo/admgo/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type (
	AuthenticationOptions struct {
		RedisConf redis.RedisConf
	}
	AuthenticationMiddleware struct {
		sessionRedis *redis.Redis
	}
)

func NewAuthenticationMiddleware(c redis.RedisConf) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		sessionRedis: redis.MustNewRedis(c, redis.WithPass(c.Pass)),
	}
}

func (m *AuthenticationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqCtx := r.Context()
		cookie, err := r.Cookie("access_token")
		if err != nil {
			logx.Error(err)
			httpx.Error(w, errorx.NoLogin)
			return
		}
		isExist, err := m.sessionRedis.ExistsCtx(reqCtx, "/user/rest/session/"+cookie.Value)
		if err != nil {
			logx.Error(err)
			httpx.Error(w, errorx.ServerErrForRedis)
			return
		}
		if !isExist {
			logx.Errorf("%s", err)
			httpx.Error(w, errorx.NoLogin)
			return
		}
		val, err := m.sessionRedis.HgetCtx(r.Context(), "/user/rest/session/"+cookie.Value, "user_id")
		if err != nil {
			httpx.Error(w, errorx.ServerErrForRedis)
			return
		}
		ctx := context.WithValue(reqCtx, "UserID", val)
		newReq := r.WithContext(ctx)

		next(w, newReq)
	}
}
