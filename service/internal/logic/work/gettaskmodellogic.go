package work

import (
	"context"

	enttask "ai-gallery/ent/task"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskModelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskModelLogic {
	return &GetTaskModelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskModelLogic) GetTaskModel() (resp *types.GetTaskModelResponse, err error) {
	var sd []struct {
		SdModelName string `json:"sd_model_name"`
	}
	if err = dao.EntClient.Task.Query().
		GroupBy(enttask.FieldSdModelName).
		Scan(l.ctx, &sd); err != nil {
		return nil, err
	}
	models := make([]string, 0)
	for _, s := range sd {
		models = append(models, s.SdModelName)
	}
	return &types.GetTaskModelResponse{
		Models: models,
	}, nil
}
