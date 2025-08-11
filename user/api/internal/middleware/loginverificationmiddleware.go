package middleware

import "net/http"

type LoginVerificationMiddleware struct {
}

func NewLoginVerificationMiddleware() *LoginVerificationMiddleware {
	return &LoginVerificationMiddleware{}
}

func (m *LoginVerificationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 验证token通过
		if r.Header.Get("token") == "123456" {
			next(w, r)
			return
		}

		// Passthrough to next handler if need
		w.Write([]byte("无效token"))
		return
	}
}
