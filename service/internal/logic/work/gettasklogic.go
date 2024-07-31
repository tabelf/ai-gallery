package work

import (
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/dao/task"
	"context"
	"time"

	enttaskdetail "ai-gallery/ent/taskdetail"
	"ai-gallery/pkg/utils"
	"ai-gallery/service/internal/logic/basic"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/duke-git/lancet/v2/strutil"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogic {
	return &GetTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskLogic) GetTask(req *types.GetTaskRequest) (resp *types.GetTaskResponse, err error) {
	data := make([]*types.TaskBo, 0)

	authorIDs, err := basic.GetAuthorIDsByRole(l.ctx, req.AuthorIDs)
	if err != nil {
		return nil, err
	}
	start, end := time.Time{}, time.Time{}
	if strutil.IsNotBlank(req.StartTime) {
		start, err = utils.ParseYMD(req.StartTime)
		if err != nil {
			return nil, err
		}
		start = utils.FirstTime(start)
	}
	if strutil.IsNotBlank(req.EndTime) {
		end, err = utils.ParseYMD(req.EndTime)
		if err != nil {
			return nil, err
		}
		end = utils.LastTime(end)
	}
	query := task.GetQueryByPO(&task.PO{
		Category:  req.Category,
		AuthorIDs: authorIDs,
		Status:    req.Status,
		StartTime: start,
		EndTime:   end,
	})
	count, err := query.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &types.GetTaskResponse{
			Data:  data,
			Total: count,
		}, nil
	}
	tasks, err := query.Offset(req.Offset).Limit(req.Limit).All(l.ctx)
	if err != nil {
		return nil, err
	}
	taskIDs := make([]int, 0)
	for _, t := range tasks {
		taskIDs = append(taskIDs, t.ID)
	}
	taskImageMap := make(map[int][]string, 0)
	details, err := dao.EntClient.TaskDetail.Query().Where(
		enttaskdetail.DeleteTimeIsNil(),
		enttaskdetail.TaskIDIn(taskIDs...),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	for _, detail := range details {
		taskImageMap[detail.TaskID] = append(taskImageMap[detail.TaskID], detail.ImageURL)
	}
	for _, t := range tasks {
		data = append(data, &types.TaskBo{
			TaskID:       t.ID,
			Category:     t.Category,
			Prompt:       t.Prompt,
			Size:         t.ImgSize,
			AuthorID:     t.AuthorID,
			AuthorName:   t.AuthorName,
			Total:        t.Total,
			JobTimestamp: t.JobTimestamp.Format(time.DateTime),
			Status:       t.Status,
			Count:        len(taskImageMap[t.ID]),
			ImageUrls:    taskImageMap[t.ID],
		})
	}
	return &types.GetTaskResponse{
		Data:  data,
		Total: count,
	}, nil
}
