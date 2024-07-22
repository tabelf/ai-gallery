package account

import (
	"net/http"

	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/logic/account"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := account.NewDeleteUserLogic(r.Context(), ctx)
		err := l.DeleteUser(&req)
		if err != nil {
			errors.HandleResponseError(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
