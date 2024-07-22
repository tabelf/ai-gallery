// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	account "ai-gallery/service/internal/handler/account"
	analysis "ai-gallery/service/internal/handler/analysis"
	anonymous "ai-gallery/service/internal/handler/anonymous"
	base "ai-gallery/service/internal/handler/base"
	setting "ai-gallery/service/internal/handler/setting"
	task "ai-gallery/service/internal/handler/task"
	work "ai-gallery/service/internal/handler/work"
	"ai-gallery/service/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JWT},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/user/detail",
					Handler: account.GetUserDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/manage/user/detail",
					Handler: account.UpdateUserDetailHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ManageJWT},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/manage/user",
					Handler: account.AddUserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/user",
					Handler: account.GetUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/manage/user/:user_id",
					Handler: account.UpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/v1/manage/user/:user_id",
					Handler: account.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/manage/user/:user_id/reset",
					Handler: account.ResetPwdHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ManageJWT},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/analysis/base",
					Handler: analysis.GetAnalysisBaseHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/analysis/task",
					Handler: analysis.GetAnalysisTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/analysis/user",
					Handler: analysis.GetAnalysisUserHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/anonymous/login",
				Handler: anonymous.AuthLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/base/ping",
				Handler: base.PingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ManageJWT},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/setting",
					Handler: setting.GetSettingHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/manage/setting",
					Handler: setting.UpdateSettingHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Headers},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/sub_task",
					Handler: task.CreateSubTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/task",
					Handler: task.CreateTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/task/:task_id",
					Handler: task.UpdateTaskHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JWT},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/task",
					Handler: work.GetTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/task/author",
					Handler: work.GetTaskAuthorHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/task/model",
					Handler: work.GetTaskModelHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/task/size",
					Handler: work.GetTaskSizeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/work",
					Handler: work.GetWorkHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/manage/work/:task_id",
					Handler: work.GetWorkDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/manage/work/excellent",
					Handler: work.MarkWorkExcellentHandler(serverCtx),
				},
			}...,
		),
	)
}
