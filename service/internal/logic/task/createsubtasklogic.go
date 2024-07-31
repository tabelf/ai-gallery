package task

import (
	"ai-gallery/service/internal/dao"
	"bytes"
	"context"
	"encoding/base64"
	"time"

	enttask "ai-gallery/ent/task"
	"ai-gallery/service/internal/errors"
	basic "ai-gallery/service/internal/logic/basic/upload"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSubTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubTaskLogic {
	return &CreateSubTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSubTaskLogic) CreateSubTask(req *types.CreateSubTaskReq) (resp *types.CreateSubTaskResp, err error) {
	task, err := dao.EntClient.Task.Query().
		Where(enttask.ID(int(req.TaskID))).
		First(l.ctx)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(req.Base64)
	if err != nil {
		return nil, errors.ErrTaskImageBase64Parse(l.ctx)
	}
	_, upload, err := basic.UploadStrategy(l.ctx)
	if err != nil {
		return nil, err
	}
	path, err := upload.Put(l.ctx, req.Filename, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	subTask, err := dao.EntClient.TaskDetail.Create().
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		SetImageURL(path).
		SetTaskID(task.ID).
		Save(l.ctx)
	if err != nil {
		return nil, err
	}
	return &types.CreateSubTaskResp{
		SubTaskID: subTask.ID,
	}, nil
}
