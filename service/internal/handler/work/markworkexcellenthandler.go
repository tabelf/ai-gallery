package work

import (
	"net/http"

	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/logic/work"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MarkWorkExcellentHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MarkWorkExcellentRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := work.NewMarkWorkExcellentLogic(r.Context(), ctx)
		err := l.MarkWorkExcellent(&req)
		if err != nil {
			errors.HandleResponseError(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
