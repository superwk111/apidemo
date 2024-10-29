package svc

import (
	"apidemo/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Mysql  sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config: c,
		Mysql:  sqlConn,
	}
}
