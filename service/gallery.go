package service

import (
	"context"
	"flag"
	"fmt"

	"ai-gallery/ent"
	"ai-gallery/service/internal/config"
	"ai-gallery/service/internal/cron"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/dao/account"
	"ai-gallery/service/internal/dao/setting"
	"ai-gallery/service/internal/handler"
	"ai-gallery/service/internal/middleware"
	"ai-gallery/service/internal/model"
	"ai-gallery/service/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"
	"golang.org/x/crypto/bcrypt"
)

func StartHTTP(env string) {
	c := LoadConfig(env)
	ctx := context.Background()

	server := rest.MustNewServer(c.RestConf,
		rest.WithNotAllowedHandler(middleware.NewCorsMiddleware().Handler()))
	defer server.Stop()

	svcCtx := svc.NewServiceContext(c)

	_, err := InitDB(c)
	if err != nil {
		logc.Errorf(ctx, "init db failed: %+v", err)
		return
	}
	// 创建系统数据.
	if err = createSystemData(ctx); err != nil {
		logc.Errorf(ctx, "init admin account failed: %+v", err)
		return
	}

	server.Use(middleware.NewCorsMiddleware().Handle)

	handler.RegisterHandlers(server, svcCtx)

	cronJob, err := cron.CreateCronJob(c)
	if err != nil {
		logc.Errorf(ctx, "init cron failed: %+v", err)
		return
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	cronJob.Start()
	server.Start()
}

func InitDB(c config.Config) (*ent.Client, error) {
	return dao.NewDB(c.DB)
}

const (
	Dev  = "dev"
	Prod = "prod"
)

var devConfigFile = flag.String(Dev, "service/etc/application-dev.yaml", "loading dev config file")
var prodConfigFile = flag.String(Prod, "service/etc/application-prod.yaml", "loading prod config file")

func LoadConfig(env string) config.Config {
	var (
		c    config.Config
		file *string
	)
	if env == Prod {
		file = prodConfigFile
	} else {
		file = devConfigFile
	}

	conf.MustLoad(*file, &c)
	logc.MustSetup(c.Logger)
	return c
}

func createSystemData(ctx context.Context) error {
	if err := setting.Create(ctx, &setting.PO{
		StoreName:    model.UploadEnums.Local.Code,
		StoreAddress: model.LocalAddress,
		SecureID:     "",
		SecureKey:    "",
		StoreBucket:  "",
		InitPwd:      model.InitPwd,
	}); err != nil {
		return err
	}
	_, err := account.GetByUsername(ctx, model.Admin)
	if err != nil {
		if ent.IsNotFound(err) {
			initPwd, err := setting.GetInitPwd(ctx)
			if err != nil {
				return err
			}
			password, err := bcrypt.GenerateFromPassword([]byte(initPwd+model.PwdSalt), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			return account.Create(ctx, &account.PO{
				Username: model.Admin,
				Nickname: model.Admin,
				Password: string(password),
				Status:   model.ACTIVE,
				Role:     model.ADMIN,
			})
		}
		return err
	}
	return nil
}
