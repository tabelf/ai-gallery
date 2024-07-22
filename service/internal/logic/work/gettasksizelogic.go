package work

import (
	"ai-gallery/service/internal/dao"
	"context"

	enttask "ai-gallery/ent/task"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskSizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskSizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskSizeLogic {
	return &GetTaskSizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskSizeLogic) GetTaskSize() (resp *types.GetTaskSizeResponse, err error) {
	var size []struct {
		Size string `json:"img_size"`
	}
	if err = dao.EntClient.Task.Query().
		GroupBy(enttask.FieldImgSize).
		Scan(l.ctx, &size); err != nil {
		return nil, err
	}
	sizes := make([]string, 0)
	for _, s := range size {
		sizes = append(sizes, s.Size)
	}
	return &types.GetTaskSizeResponse{
		Sizes: sizes,
	}, nil
}
