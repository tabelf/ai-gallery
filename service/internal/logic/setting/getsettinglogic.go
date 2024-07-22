package setting

import (
	"ai-gallery/service/internal/dao/setting"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSettingLogic {
	return &GetSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSettingLogic) GetSetting() (resp *types.GetSettingResponse, err error) {
	po, err := setting.GetPO(l.ctx)
	if err != nil {
		return nil, err
	}
	return &types.GetSettingResponse{
		StoreName:    po.StoreName,
		StoreAddress: po.StoreAddress,
		SecureID:     po.SecureID,
		SecureKey:    po.SecureKey,
		StoreBucket:  po.StoreBucket,
		InitPwd:      po.InitPwd,
	}, nil
}
