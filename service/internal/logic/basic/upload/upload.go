package basic

import (
	"context"
	"io"

	"ai-gallery/service/internal/dao/setting"
	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/model"
)

var uploadServiceMap = make(map[string]UploadService, 0)

type UploadService interface {
	SetConfig(url, secretID, secretKey, bucket string)
	Put(ctx context.Context, name string, r io.Reader, options ...any) (path string, err error)
}

func UploadStrategy(ctx context.Context) (string, UploadService, error) {
	_initService()
	conf, err := setting.GetPO(ctx)
	if err != nil {
		return "", nil, err
	}
	service := uploadServiceMap[conf.StoreName]
	if service == nil {
		return "", nil, errors.ErrUploadServiceNotFound(ctx)
	}
	service.SetConfig(conf.StoreAddress, conf.SecureID, conf.SecureKey, conf.StoreBucket)
	return conf.StoreName, service, nil
}

func _initService() {
	if len(uploadServiceMap) > 0 {
		return
	}
	uploadServiceMap[model.UploadEnums.Tx.Code] = NewCosService()
	uploadServiceMap[model.UploadEnums.Aliyun.Code] = NewOssService()
	uploadServiceMap[model.UploadEnums.Local.Code] = NewLocalService()
}
