package analysis

import (
	"context"
	"time"

	"ai-gallery/ent/account"
	enttask "ai-gallery/ent/task"
	"ai-gallery/pkg/utils"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/model"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnalysisUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAnalysisUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnalysisUserLogic {
	return &GetAnalysisUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAnalysisUserLogic) GetAnalysisUser(req *types.GetAnalysisUserRequest) (resp *types.GetAnalysisUserResponse, err error) {
	start := utils.FirstTime(time.Now()).AddDate(0, 0, -1*req.Day-1)
	end := utils.LastTime(time.Now())

	authorBo := make([]*types.AnalysisUserBo, 0)
	accounts, err := dao.EntClient.Account.Query().Where(
		account.DeleteTimeIsNil(),
		account.Status(model.ACTIVE),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	for _, acc := range accounts {
		count, err := dao.EntClient.Task.Query().Where(
			enttask.DeleteTimeIsNil(),
			enttask.CreateTimeGTE(start),
			enttask.CreateTimeLTE(end),
			enttask.AuthorID(acc.ID),
		).Count(l.ctx)
		if err != nil {
			return nil, err
		}
		authorBo = append(authorBo, &types.AnalysisUserBo{
			AuthorID:   acc.ID,
			AuthorName: acc.Nickname,
			Count:      count,
		})
	}
	return &types.GetAnalysisUserResponse{
		AnalysisUserBo: authorBo,
	}, nil
}
