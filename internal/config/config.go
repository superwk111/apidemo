package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySQL struct {
		DataSource string
	}
	JWTSecret string
}
type MysqlConfig struct {
	DataSource     string
	MaxLifetime    int
	ConnectTimeout int
}

