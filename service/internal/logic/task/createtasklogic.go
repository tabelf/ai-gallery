package task

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"math"
	"time"

	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/logic/basic"
	u "ai-gallery/service/internal/logic/basic/upload"
	"ai-gallery/service/internal/model"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskLogic) CreateTask(req *types.CreateTaskReq) (resp *types.CreateTaskResp, err error) {
	info, err := basic.ContextElement(l.ctx)
	if err != nil {
		return nil, err
	}
	jobTime, err := time.ParseInLocation(model.DateStrLayout, req.JobTimestamp, time.Local)
	if err != nil {
		return nil, err
	}
	refImages := make([]string, 0)
	store, upload, err := u.UploadStrategy(l.ctx)
	if err != nil {
		return nil, err
	}
	for _, baseData := range req.RefImages {
		data, err := base64.StdEncoding.DecodeString(baseData.Base64)
		if err != nil {
			return nil, errors.ErrTaskImageBase64Parse(l.ctx)
		}
		path, err := upload.Put(l.ctx, baseData.Filename, bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		refImages = append(refImages, path)
	}
	ratio := calculateRatio(req.Width, req.Height)
	createSQL := dao.EntClient.Task.Create().
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		SetCategory(req.Category).
		SetPrompt(req.Prompt).
		SetImgSize(ratio).
		SetNegativePrompt(req.NegativePrompt).
		SetWeight(req.Width).
		SetHeight(req.Height).
		SetSeed(req.Seed).
		SetSamplerName(req.SamplerName).
		SetSteps(req.Steps).
		SetCfgScale(req.CfgScale).
		SetBatchSize(req.BatchSize).
		SetTotal(req.Total).
		SetSdModelName(req.SdModelName).
		SetSdModelHash(req.SdModelHash).
		SetNillableSdVaeName(req.SdVaeName).
		SetNillableSdVaeHash(req.SdVaeHash).
		SetJobTimestamp(jobTime).
		SetVersion(req.Version).
		SetStatus(model.TaskStatusEnums.Progress.Code).
		SetAuthorID(info.ID).
		SetAuthorName(info.Nickname).
		SetStore(store)
	if len(refImages) > 0 {
		createSQL.SetRefImages(refImages)
	}
	task, err := createSQL.Save(l.ctx)
	if err != nil {
		return nil, err
	}
	return &types.CreateTaskResp{
		TaskID: task.ID,
	}, nil
}

func calculateRatio(length, width float32) string {
	if width == 0 {
		return "undefined" // 防止除以零错误
	}

	// 求最大公约数
	gcdValue := gcd(float64(length), float64(width))

	// 计算比例
	ratio1 := length / float32(gcdValue)
	ratio2 := width / float32(gcdValue)

	// 保留最多一位小数
	ratio1Str := formatRatio(ratio1)
	ratio2Str := formatRatio(ratio2)

	return ratio1Str + ":" + ratio2Str
}

// gcd 计算两个数的最大公约数
func gcd(a, b float64) float64 {
	for b != 0 {
		t := b
		b = math.Mod(a, b)
		a = t
	}
	return a
}

func formatRatio(value float32) string {
	if value == float32(int(value)) {
		return fmt.Sprintf("%d", int(value))
	}
	return fmt.Sprintf("%.2f", value)
}
