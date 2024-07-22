package jobs

import (
	"context"
	"time"

	enttask "ai-gallery/ent/task"
	"ai-gallery/pkg/utils"
	"ai-gallery/service/internal/config"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/model"

	"github.com/zeromicro/go-zero/core/logc"
)

func CheckTaskStatus(ctx context.Context, config config.Config) {
	logc.Info(ctx, "check task status cron.")
	// 检查 今天0晨 至 30 分钟前未完成的任务
	tasks, err := dao.EntClient.Task.Query().Where(
		enttask.CreateTimeGTE(utils.FirstTime(time.Now())),
		enttask.CreateTimeLTE(time.Now().Add(time.Minute*30)),
		enttask.Status(model.TaskStatusEnums.Progress.Code),
		enttask.DeleteTimeIsNil(),
	).All(ctx)
	if err != nil {
		logc.Errorf(ctx, "select task error: %+v", err)
		return
	}
	if len(tasks) == 0 {
		return
	}
	taskIDs := make([]int, 0)
	for _, t := range tasks {
		taskIDs = append(taskIDs, t.ID)
	}
	if err = dao.EntClient.Task.Update().
		SetStatus(model.TaskStatusEnums.Exception.Code).
		Where(enttask.IDIn(taskIDs...)).
		Exec(ctx); err != nil {
		logc.Errorf(ctx, "update task status error: %+v", err)
		return
	}
}
