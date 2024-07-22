package setting

import (
	"net/http"

	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/logic/setting"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSettingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSettingRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := setting.NewUpdateSettingLogic(r.Context(), ctx)
		err := l.UpdateSetting(&req)
		if err != nil {
			errors.HandleResponseError(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
