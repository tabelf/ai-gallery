package work

import (
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/dao/task"
	"context"
	"time"

	enttaskdetail "ai-gallery/ent/taskdetail"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWorkDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkDetailLogic {
	return &GetWorkDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWorkDetailLogic) GetWorkDetail(req *types.GetWorkDetailRequest) (resp *types.GetWorkDetailResponse, err error) {
	t, err := task.GetByID(l.ctx, req.TaskID)
	if err != nil {
		return nil, err
	}
	details, err := dao.EntClient.TaskDetail.Query().Where(
		enttaskdetail.DeleteTimeIsNil(),
		enttaskdetail.TaskID(t.ID),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	bos := make([]*types.WorkDetailBo, 0)
	for _, d := range details {
		bos = append(bos, &types.WorkDetailBo{
			SubTaskID:    d.ID,
			ImageURL:     d.ImageURL,
			HasExcellent: d.HasExcellent,
		})
	}
	return &types.GetWorkDetailResponse{
		TaskID:         t.ID,
		Category:       t.Category,
		Prompt:         t.Prompt,
		NegativePrompt: t.NegativePrompt,
		Width:          t.Weight,
		Height:         t.Height,
		Size:           t.ImgSize,
		Seed:           t.Seed,
		SamplerName:    t.SamplerName,
		Steps:          t.Steps,
		CfgScale:       t.CfgScale,
		BatchSize:      t.BatchSize,
		Total:          t.Total,
		SdModelHash:    t.SdModelHash,
		SdModelName:    t.SdModelName,
		SdVaeHash:      t.SdVaeHash,
		SdVaeName:      t.SdVaeName,
		JobTimestamp:   t.JobTimestamp.Format(time.DateTime),
		Version:        t.Version,
		AuthorID:       t.AuthorID,
		AuthorName:     t.AuthorName,
		RefImages:      t.RefImages,
		WorkDetailBos:  bos,
	}, nil
}
