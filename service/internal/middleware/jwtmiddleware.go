package middleware

import (
	"context"
	"net/http"

	"ai-gallery/pkg/jwt"
	"ai-gallery/service/internal/config"
	"ai-gallery/service/internal/logic/basic"
	"ai-gallery/service/internal/model"
)

type JWTMiddleware struct {
	config.JWTConf
}

func NewJWTMiddleware(conf config.JWTConf) *JWTMiddleware {
	return &JWTMiddleware{conf}
}

func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "未提供有效身份许可", http.StatusUnauthorized)
			return
		}
		tokenType, claims := jwt.IsValidateOfToken(token, m.JwtKey)
		var resultMessage string
		reqCtx := r.Context()
		switch tokenType {
		case jwt.ValidateToken:
			{
				if err := basic.CheckAccountStatus(reqCtx, claims.ID); err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				r = r.WithContext(context.WithValue(reqCtx, model.UserToken, claims))
				next(w, r)
				return
			}
		case jwt.ExpiredToken:
			resultMessage = "令牌已失效!"
		case jwt.BadToken:
			resultMessage = "未提供有效令牌"
		}
		http.Error(w, resultMessage, http.StatusUnauthorized)
	}
}
