package dao

import (
	"github.com/isther/clinic/conf"
	"github.com/isther/clinic/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	initSQL()
}

func initSQL() {
	var err error

	// 连接数据库
	DB, err = gorm.Open(postgres.Open(conf.Server.DSN()), &gorm.Config{})
	if err != nil {
		logrus.Fatalln(err)
	}

	// 绑定模型
	err = DB.AutoMigrate(&model.FormSql{})
	if err != nil {
		logrus.Fatalln(err)
	}
}
