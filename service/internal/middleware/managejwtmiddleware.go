package middleware

import (
	"ai-gallery/service/internal/errors"
	"context"
	"net/http"

	"ai-gallery/pkg/jwt"
	"ai-gallery/service/internal/config"
	"ai-gallery/service/internal/logic/basic"
	"ai-gallery/service/internal/model"
)

type ManageJWTMiddleware struct {
	config.JWTConf
}

func NewManageJWTMiddleware(conf config.JWTConf) *ManageJWTMiddleware {
	return &ManageJWTMiddleware{conf}
}

func (m *ManageJWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, errors.ErrPermissionProvider(r.Context()).Error(), http.StatusUnauthorized)
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
				if err := basic.CheckPermission(r.Context(), claims); err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				r = r.WithContext(context.WithValue(reqCtx, model.UserToken, claims))
				next(w, r)
				return
			}
		case jwt.ExpiredToken:
			resultMessage = errors.ErrTokenExpired(reqCtx).Error()
		case jwt.BadToken:
			resultMessage = errors.ErrTokenValid(reqCtx).Error()
		}
		http.Error(w, resultMessage, http.StatusUnauthorized)
	}
}
