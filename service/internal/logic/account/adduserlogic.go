package account

import (
	"ai-gallery/ent"
	"ai-gallery/service/internal/dao/setting"
	"ai-gallery/service/internal/errors"
	"context"

	"ai-gallery/service/internal/dao/account"
	"ai-gallery/service/internal/model"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserRequest) error {
	acc, err := account.GetByUsername(l.ctx, req.Username)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	err = nil
	if acc != nil {
		return errors.ErrAccountUsernameExist(l.ctx)
	}
	initPwd, err := setting.GetInitPwd(l.ctx)
	if err != nil {
		return err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(initPwd+model.PwdSalt), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return account.Create(l.ctx, &account.PO{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: string(password),
		Status:   model.DISABLED,
		Role:     req.Role,
	})
}
