package analysis

import (
	"context"
	"time"

	enttask "ai-gallery/ent/task"
	"ai-gallery/ent/taskdetail"
	"ai-gallery/pkg/errgroup"
	"ai-gallery/pkg/utils"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnalysisTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAnalysisTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnalysisTaskLogic {
	return &GetAnalysisTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAnalysisTaskLogic) GetAnalysisTask(req *types.GetAnalysisTaskRequest) (resp *types.GetAnalysisTaskResponse, err error) {
	var (
		// 任务数
		taskData      = make([]int, req.Day)
		workData      = make([]int, req.Day)
		excellentData = make([]int, req.Day)

		g     = errgroup.WithContext(l.ctx)
		times = make([]string, 0)
	)

	for i := req.Day - 1; i >= 0; i-- {
		idx := req.Day - (i + 1)
		curr := time.Now().AddDate(0, 0, -1*i)
		times = append(times, curr.Format(time.DateOnly))
		start := utils.FirstTime(curr)
		end := utils.LastTime(curr)

		g.Go(func(ctx context.Context) (err error) {
			count, err := dao.EntClient.Task.Query().Where(
				enttask.DeleteTimeIsNil(),
				enttask.CreateTimeGTE(start),
				enttask.CreateTimeLTE(end),
			).Count(ctx)
			if err != nil {
				return err
			}
			taskData[idx] = count
			return nil
		})
		g.Go(func(ctx context.Context) (err error) {
			count, err := dao.EntClient.TaskDetail.Query().Where(
				taskdetail.DeleteTimeIsNil(),
				taskdetail.CreateTimeGTE(start),
				taskdetail.CreateTimeLTE(end),
			).Count(ctx)
			if err != nil {
				return err
			}
			workData[idx] = count
			return nil
		})
		g.Go(func(ctx context.Context) (err error) {
			count, err := dao.EntClient.TaskDetail.Query().
				Where(
					taskdetail.DeleteTimeIsNil(),
					taskdetail.HasExcellent(true),
					taskdetail.CreateTimeGTE(start),
					taskdetail.CreateTimeLTE(end),
				).
				Count(ctx)
			if err != nil {
				return err
			}
			excellentData[idx] = count
			return nil
		})
	}
	if err = g.Wait(); err != nil {
		return nil, err
	}
	return &types.GetAnalysisTaskResponse{
		Times: times,
		AnalysisTaskBo: []*types.AnalysisTaskBo{
			{
				Name:   "任务数",
				EnName: "tasks",
				Data:   taskData,
			},
			{
				Name:   "作品数",
				EnName: "works",
				Data:   workData,
			},
			{
				Name:   "优秀作品数",
				EnName: "excellent works",
				Data:   excellentData,
			},
		},
	}, nil
}
