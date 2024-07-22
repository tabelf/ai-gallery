package account

import (
	"ai-gallery/service/internal/dao/setting"
	"context"

	"ai-gallery/service/internal/dao/account"
	"ai-gallery/service/internal/model"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type ResetPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPwdLogic {
	return &ResetPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPwdLogic) ResetPwd(req *types.ResetPwdRequest) error {
	initPwd, err := setting.GetInitPwd(l.ctx)
	if err != nil {
		return err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(initPwd+model.PwdSalt), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return account.Update(l.ctx, &account.PO{
		ID:       req.UserID,
		Password: string(password),
		Status:   -1,
	})
}
