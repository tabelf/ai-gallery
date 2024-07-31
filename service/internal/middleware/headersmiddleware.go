package middleware

import (
	"ai-gallery/service/internal/errors"
	"context"
	"encoding/json"
	"net/http"

	"ai-gallery/pkg/jwt"
	"ai-gallery/service/internal/model"

	"ai-gallery/service/internal/logic/basic"
	"github.com/duke-git/lancet/v2/strutil"
)

type HeadersMiddleware struct{}

func NewHeadersMiddleware() *HeadersMiddleware {
	return &HeadersMiddleware{}
}

const NewHeaderKey = "user_header"

type HeaderUserBoInfo struct {
	UserID   int    `json:"user_id"`
	Nickname string `json:"nickname"`
}

type UserRo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m *HeadersMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header.Get("X-Ag-H-U")
		if strutil.IsBlank(headers) {
			http.Error(w, errors.ErrHeaderNotFound(r.Context()).Error(), http.StatusUnauthorized)
			return
		}
		var user UserRo
		if err := json.Unmarshal([]byte(headers), &user); err != nil {
			http.Error(w, errors.ErrHeaderParsing(r.Context()).Error(), http.StatusUnauthorized)
			return
		}
		account, err := basic.ValidateAccount(r.Context(), user.Username, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), model.UserToken, &jwt.Claims{
			ID:       account.ID,
			Role:     account.Role,
			Nickname: account.Nickname,
		}))
		next(w, r)
	}
}
