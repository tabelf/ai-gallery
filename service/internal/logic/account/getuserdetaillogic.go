package account

import (
	"ai-gallery/service/internal/logic/basic"
	"context"

	"ai-gallery/service/internal/dao/account"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailLogic) GetUserDetail() (resp *types.GetUserDetailResponse, err error) {
	jwtInfo, err := basic.ContextElement(l.ctx)
	if err != nil {
		return nil, err
	}
	acc, err := account.GetByID(l.ctx, jwtInfo.ID)
	if err != nil {
		return nil, err
	}
	return &types.GetUserDetailResponse{
		UserID:   acc.ID,
		Username: acc.Username,
		Nickname: acc.Nickname,
		Role:     acc.Role,
	}, nil
}
