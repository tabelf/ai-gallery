package account

import (
	"context"

	"ai-gallery/ent"
	"ai-gallery/service/internal/dao/account"
	"ai-gallery/service/internal/dao/task"
	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) error {
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
	if err = account.Update(l.ctx, &account.PO{
		ID:       req.UserID,
		Username: req.Username,
		Nickname: req.Nickname,
		Status:   req.Status,
		Role:     req.Role,
	}); err != nil {
		return err
	}
	if acc.Nickname != req.Nickname {
		// 修改了昵称
		return task.UpdateAuthorInfo(l.ctx, acc.ID, req.Nickname)
	}
	return nil
}
