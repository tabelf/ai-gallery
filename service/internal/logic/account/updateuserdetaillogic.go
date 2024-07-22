package account

import (
	"ai-gallery/ent"
	"ai-gallery/service/internal/dao/account"
	"ai-gallery/service/internal/dao/task"
	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/model"
	"context"
	"github.com/duke-git/lancet/v2/strutil"
	"golang.org/x/crypto/bcrypt"

	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserDetailLogic {
	return &UpdateUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserDetailLogic) UpdateUserDetail(req *types.UpdateUserDetailRequest) error {
	acc, err := account.GetByID(l.ctx, req.UserID)
	if err != nil {
		return err
	}
	if acc.Username != req.Username {
		other, err := account.GetByUsername(l.ctx, req.Username)
		if err != nil && !ent.IsNotFound(err) {
			return err
		}
		err = nil
		if other != nil {
			return errors.ErrAccountUsernameExist(l.ctx)
		}
	}
	newPwd := make([]byte, 0)
	if strutil.IsNotBlank(req.NewPassword) {
		if err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(req.OldPassword+model.PwdSalt)); err != nil {
			return errors.ErrAccountPwd(l.ctx)
		}
		newPwd, err = bcrypt.GenerateFromPassword([]byte(req.NewPassword+model.PwdSalt), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
	}
	if err = account.Update(l.ctx, &account.PO{
		ID:       req.UserID,
		Username: req.Username,
		Nickname: req.Nickname,
		Password: string(newPwd),
		Status:   -1,
	}); err != nil {
		return err
	}
	if acc.Nickname != req.Nickname {
		// 修改了昵称
		return task.UpdateAuthorInfo(l.ctx, acc.ID, req.Nickname)
	}
	return nil
}
