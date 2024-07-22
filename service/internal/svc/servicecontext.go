package svc

import (
	"ai-gallery/service/internal/config"
	"ai-gallery/service/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Headers   rest.Middleware
	JWT       rest.Middleware
	ManageJWT rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Headers:   middleware.NewHeadersMiddleware().Handle,
		JWT:       middleware.NewJWTMiddleware(c.JWT).Handle,
		ManageJWT: middleware.NewManageJWTMiddleware(c.JWT).Handle,
	}
}
