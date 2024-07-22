package work

import (
	"context"

	"ai-gallery/ent"
	enttask "ai-gallery/ent/task"
	"ai-gallery/ent/taskdetail"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkWorkExcellentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMarkWorkExcellentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkWorkExcellentLogic {
	return &MarkWorkExcellentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MarkWorkExcellentLogic) MarkWorkExcellent(req *types.MarkWorkExcellentRequest) error {
	if err := dao.WithTx(l.ctx, dao.EntClient, func(tx *ent.Tx) (err error) {
		if err = tx.TaskDetail.Update().
			SetHasExcellent(req.HasExcellent).
			Where(
				taskdetail.ID(req.SubTaskID),
				taskdetail.DeleteTimeIsNil(),
			).Exec(l.ctx); err != nil {
			return err
		}
		exist, err := tx.TaskDetail.Query().Where(
			taskdetail.TaskID(req.TaskID),
			taskdetail.HasExcellent(true),
			taskdetail.DeleteTimeIsNil(),
		).Exist(l.ctx)
		if err != nil {
			return err
		}
		return tx.Task.Update().
			SetHasExcellent(exist).
			Where(enttask.ID(req.TaskID)).
			Exec(l.ctx)
	}); err != nil {
		return err
	}
	return nil
}
