package config

import (
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
