package analysis

import (
	"context"
	"time"

	"ai-gallery/ent/account"
	enttask "ai-gallery/ent/task"
	"ai-gallery/ent/taskdetail"
	"ai-gallery/pkg/errgroup"
	"ai-gallery/pkg/utils"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/model"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnalysisBaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAnalysisBaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnalysisBaseLogic {
	return &GetAnalysisBaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAnalysisBaseLogic) GetAnalysisBase() (resp *types.GetAnalysisBaseResponse, err error) {
	var (
		taskTotal           = 0
		workTotal           = 0
		excellentTotal      = 0
		accountTotal        = 0
		todayTaskCount      = 0
		todayWorkCount      = 0
		todayExcellentCount = 0
		g                   = errgroup.WithContext(l.ctx)
		start               = utils.FirstTime(time.Now())
		end                 = utils.LastTime(time.Now())
	)
	g.Go(func(ctx context.Context) (err error) {
		taskTotal, err = dao.EntClient.Task.Query().
			Where(enttask.DeleteTimeIsNil()).
			Count(ctx)
		return err
	})
	g.Go(func(ctx context.Context) (err error) {
		workTotal, err = dao.EntClient.TaskDetail.Query().
			Where(taskdetail.DeleteTimeIsNil()).
			Count(ctx)
		return err
	})
	g.Go(func(ctx context.Context) error {
		excellentTotal, err = dao.EntClient.TaskDetail.Query().
			Where(
				taskdetail.DeleteTimeIsNil(),
				taskdetail.HasExcellent(true),
			).
			Count(ctx)
		return err
	})
	g.Go(func(ctx context.Context) error {
		accountTotal, err = dao.EntClient.Account.Query().Where(
			account.DeleteTimeIsNil(),
			account.Status(model.ACTIVE),
		).Count(ctx)
		return err
	})
	g.Go(func(ctx context.Context) (err error) {
		todayTaskCount, err = dao.EntClient.Task.Query().
			Where(
				enttask.DeleteTimeIsNil(),
				enttask.CreateTimeGTE(start),
				enttask.CreateTimeLTE(end),
			).Count(ctx)
		return err
	})
	g.Go(func(ctx context.Context) (err error) {
		todayWorkCount, err = dao.EntClient.TaskDetail.Query().
			Where(
				taskdetail.DeleteTimeIsNil(),
				taskdetail.CreateTimeGTE(start),
				taskdetail.CreateTimeLTE(end),
			).
			Count(ctx)
		return err
	})
	g.Go(func(ctx context.Context) error {
		todayExcellentCount, err = dao.EntClient.TaskDetail.Query().
			Where(
				taskdetail.DeleteTimeIsNil(),
				taskdetail.HasExcellent(true),
				taskdetail.CreateTimeGTE(start),
				taskdetail.CreateTimeLTE(end),
			).
			Count(ctx)
		return err
	})
	if err = g.Wait(); err != nil {
		return nil, err
	}
	return &types.GetAnalysisBaseResponse{
		TaskTotal:           taskTotal,
		WorkTotal:           workTotal,
		ExcellentTotal:      excellentTotal,
		AccountTotal:        accountTotal,
		TodayTaskCount:      todayTaskCount,
		TodayWorkCount:      todayWorkCount,
		TodayExcellentCount: todayExcellentCount,
	}, nil
}
