package account

import (
	"net/http"

	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/logic/account"
	"ai-gallery/service/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserDetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewGetUserDetailLogic(r.Context(), ctx)
		resp, err := l.GetUserDetail()
		if err != nil {
			errors.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
