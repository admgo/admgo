package middleware

import "net/http"

type CSRFTokenMiddleware struct {
}

func NewCSRFTokenMiddleware() *CSRFTokenMiddleware {
	return &CSRFTokenMiddleware{}
}

func (m *CSRFTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
