package setting

import (
	"ai-gallery/service/internal/dao/setting"
	"ai-gallery/service/internal/logic/basic"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSettingLogic {
	return &UpdateSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSettingLogic) UpdateSetting(req *types.UpdateSettingRequest) error {
	jwtInfo, err := basic.ContextElement(l.ctx)
	if err != nil {
		return err
	}
	return setting.Update(l.ctx, &setting.PO{
		StoreName:    req.StoreName,
		StoreAddress: req.StoreAddress,
		SecureID:     req.SecureID,
		SecureKey:    req.SecureKey,
		StoreBucket:  req.StoreBucket,
		InitPwd:      req.InitPwd,
		OperateID:    jwtInfo.ID,
		OperateName:  jwtInfo.Nickname,
	})
}
