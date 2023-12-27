package common

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm(DSN string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(DSN),
		&gorm.Config{
			Logger: &dbLog{},
		})
	if err != nil {
		logx.Errorf("连接mysql数据库失败, error=" + err.Error())
		return nil, err
	}
	return db, nil
}
