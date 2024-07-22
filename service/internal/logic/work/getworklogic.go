package work

import (
	"context"

	"ai-gallery/ent"
	enttask "ai-gallery/ent/task"
	enttaskdetail "ai-gallery/ent/taskdetail"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/dao/task"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWorkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkLogic {
	return &GetWorkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWorkLogic) GetWork(req *types.GetWorkRequest) (resp *types.GetWorkResponse, err error) {
	query := task.GetQueryByPO(&task.PO{
		Size:         req.Size,
		SdModelName:  req.SdModelName,
		HasRefImage:  req.HasRefImage,
		HasExcellent: req.HasExcellent,
		Sorted:       req.Sorted,
	}).Where(enttask.CountGT(0))

	count, err := query.Count(l.ctx)
	if err != nil {
		return nil, err
	}

	data := make([]*types.WorkBo, 0)
	if count == 0 {
		return &types.GetWorkResponse{
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
	headImageMap := make(map[int]*ent.TaskDetail, 0)
	details, err := dao.EntClient.TaskDetail.Query().Where(
		enttaskdetail.DeleteTimeIsNil(),
		enttaskdetail.TaskIDIn(taskIDs...),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	for _, detail := range details {
		if headImageMap[detail.TaskID] == nil ||
			(detail.HasExcellent && !headImageMap[detail.TaskID].HasExcellent) {
			headImageMap[detail.TaskID] = detail
		}
	}
	for _, t := range tasks {
		if headImageMap[t.ID] == nil {
			continue
		}
		data = append(data, &types.WorkBo{
			TaskID:     t.ID,
			Category:   t.Category,
			Prompt:     t.Prompt,
			Size:       t.ImgSize,
			AuthorID:   t.AuthorID,
			AuthorName: t.AuthorName,
			Count:      t.Count,
			HeadImage:  headImageMap[t.ID].ImageURL,
		})
	}
	return &types.GetWorkResponse{
		Data:  data,
		Total: count,
	}, nil
}
