package anonymous

import (
	"net/http"

	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/logic/anonymous"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := anonymous.NewAuthLoginLogic(r.Context(), ctx)
		resp, err := l.AuthLogin(&req)
		if err != nil {
			errors.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
