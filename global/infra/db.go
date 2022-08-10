package infra

import (
	"go_project_template/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(conf *config.AppConf) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(conf.MysqlDsn), &gorm.Config{})
	return db, err
}
