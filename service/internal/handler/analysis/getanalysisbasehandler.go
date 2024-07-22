package analysis

import (
	"net/http"

	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/logic/analysis"
	"ai-gallery/service/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAnalysisBaseHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := analysis.NewGetAnalysisBaseLogic(r.Context(), ctx)
		resp, err := l.GetAnalysisBase()
		if err != nil {
			errors.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
