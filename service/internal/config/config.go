package config

import (
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB     EntConf      `json:"db"`
	JWT    JWTConf      `json:"jwt"`
	Logger logc.LogConf `json:"logger"`
}

type EntConf struct {
	Driver       string `json:"driver"`
	URL          string `json:"url"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type JWTConf struct {
	JwtKey    string `json:"jwt_key"`
	JwtExpire int64  `json:"jwt_expire"`
	JwtIssuer string `json:"jwt_issuer"`
}

const (
	ENV  = "env"
	Dev  = "dev"
	Prod = "prod"
)

var devConfigFile = flag.String(Dev, "service/etc/application-dev.yaml", "loading dev config file")
var prodConfigFile = flag.String(Prod, "service/etc/application-prod.yaml", "loading prod config file")

func LoadConfig() Config {
	var (
		env    string
		c      Config
		config *string
	)
	flag.StringVar(&env, ENV, "", "测试环境")
	flag.Parse()
	if env == Prod {
		config = prodConfigFile
	} else {
		config = devConfigFile
	}

	conf.MustLoad(*config, &c)
	logc.MustSetup(c.Logger)
	return c
}
