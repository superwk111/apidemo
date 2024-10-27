package db

import (
	"apidemo/internal/config"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

func NewMysql(mysqlConfig config.MysqlConfig) sqlx.SqlConn {
	mysql := sqlx.NewMysql(mysqlConfig.DataSource)

	db, err := mysql.RawDB()
	if err != nil {
		logx.Error(err)
	}

	//Set connetion settings
	db.SetConnMaxLifetime(time.Duration(mysqlConfig.MaxLifetime) * time.Second)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mysqlConfig.ConnectTimeout)*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	return mysql
}
