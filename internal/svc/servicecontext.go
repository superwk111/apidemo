package svc

import (
	"apidemo/internal/config"
	"apidemo/internal/db"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Mysql  sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mysql:  db.NewMysql(c.MysqlConfig),
	}
}
