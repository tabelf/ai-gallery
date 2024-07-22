package work

import (
	"net/http"

	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/logic/work"
	"ai-gallery/service/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTaskAuthorHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := work.NewGetTaskAuthorLogic(r.Context(), ctx)
		resp, err := l.GetTaskAuthor()
		if err != nil {
			errors.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
