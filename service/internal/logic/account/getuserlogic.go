package account

import (
	"context"

	entaccount "ai-gallery/ent/account"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"entgo.io/ent/dialect/sql"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	query := dao.EntClient.Account.Query().Where(
		entaccount.DeleteTimeIsNil(),
	)
	if strutil.IsNotBlank(req.Username) {
		query.Where(func(s *sql.Selector) {
			s.Where(sql.Like(entaccount.FieldUsername, "%"+req.Username+"%"))
		})
	}
	if strutil.IsNotBlank(req.Nickname) {
		query.Where(func(s *sql.Selector) {
			s.Where(sql.Like(entaccount.FieldNickname, "%"+req.Nickname+"%"))
		})
	}
	if strutil.IsNotBlank(req.Role) {
		query.Where(entaccount.Role(req.Role))
	}
	if req.Status != -1 {
		query.Where(entaccount.Status(req.Status))
	}
	count, err := query.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	userBo := make([]*types.UserBo, 0)
	if count == 0 {
		return &types.GetUserResponse{
			Total: count,
			Data:  userBo,
		}, nil
	}
	accounts, err := query.Offset(req.Offset).Limit(req.Limit).All(l.ctx)
	if err != nil {
		return nil, err
	}
	for _, acc := range accounts {
		userBo = append(userBo, &types.UserBo{
			UserID:   acc.ID,
			Username: acc.Username,
			Nickname: acc.Nickname,
			Role:     acc.Role,
			Status:   acc.Status,
		})
	}
	return &types.GetUserResponse{
		Total: count,
		Data:  userBo,
	}, nil
}
