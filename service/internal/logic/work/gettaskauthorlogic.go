package work

import (
	"ai-gallery/service/internal/dao"
	"context"

	enttask "ai-gallery/ent/task"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskAuthorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskAuthorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskAuthorLogic {
	return &GetTaskAuthorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskAuthorLogic) GetTaskAuthor() (resp *types.GetTaskAuthorResponse, err error) {
	authorBos := make([]*types.TaskAuthorBo, 0)
	if err = dao.EntClient.Task.Query().
		GroupBy(enttask.FieldAuthorID, enttask.FieldAuthorName).
		Scan(l.ctx, &authorBos); err != nil {
		return nil, err
	}
	return &types.GetTaskAuthorResponse{
		TaskAuthorBo: authorBos,
	}, nil
}
