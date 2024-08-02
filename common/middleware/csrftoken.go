package middleware

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type (
	CSRFTokenMiddlewareOptions struct {
		RedisConf redis.RedisConf
	}
	CSRFTokenMiddleware struct {
		csrfRedis *redis.Redis
	}
)

func NewCSRFTokenInterceptorMiddleware(c redis.RedisConf) *CSRFTokenMiddleware {
	return &CSRFTokenMiddleware{
		csrfRedis: redis.MustNewRedis(c, redis.WithPass(c.Pass)),
	}
}

func (m *CSRFTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//csrfToken, err := r.Cookie("csrf_token")
		//if err != nil {
		//	httpx.Error(w, errors.New("Invalid CSRF token"))
		//	return
		//}
		//user, _ := r.Cookie("user_token")
		//m.csrfRedis.Get("/auth/csrf_token/" + user.Name)
		//println(csrfToken)
		next(w, r)
	}
}
