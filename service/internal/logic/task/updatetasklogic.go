package task

import (
	"bytes"
	"context"
	"encoding/base64"
	"time"

	enttask "ai-gallery/ent/task"
	"ai-gallery/ent/taskdetail"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/errors"
	basic "ai-gallery/service/internal/logic/basic/upload"
	"ai-gallery/service/internal/model"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/duke-git/lancet/v2/strutil"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogic {
	return &UpdateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTaskLogic) UpdateTask(req *types.UpdateTaskReq) error {
	task, err := dao.EntClient.Task.Query().
		Where(enttask.ID(req.TaskID)).
		First(l.ctx)
	if err != nil {
		return err
	}
	count, err := dao.EntClient.TaskDetail.Query().Where(
		taskdetail.TaskID(req.TaskID),
		taskdetail.DeleteTimeIsNil(),
	).Count(l.ctx)
	if err != nil {
		return err
	}
	status := model.TaskStatusEnums.Complete.Code
	if task.Total != count {
		status = model.TaskStatusEnums.Exception.Code
	}
	path := ""
	if strutil.IsNotBlank(req.Base64) {
		data, err := base64.StdEncoding.DecodeString(req.Base64)
		if err != nil {
			return errors.ErrTaskImageBase64Parse(l.ctx)
		}
		_, upload, err := basic.UploadStrategy(l.ctx)
		if err != nil {
			return err
		}
		path, err = upload.Put(l.ctx, req.Filename, bytes.NewReader(data))
		if err != nil {
			return err
		}
	}
	if err = dao.EntClient.Task.Update().
		SetStatus(status).
		SetUpdateTime(time.Now()).
		SetGridImageURL(path).
		SetCount(count).
		SetExtra(req.Extra).
		Where(enttask.ID(req.TaskID)).
		Exec(l.ctx); err != nil {
		return err
	}
	return nil
}
