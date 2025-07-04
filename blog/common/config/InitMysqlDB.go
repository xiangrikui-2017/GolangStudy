package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitMysqlDB() *gorm.DB {
	host := Conf.DB.Host
	port := Conf.DB.Port
	username := Conf.DB.Username
	password := Conf.DB.Password
	database := Conf.DB.Database
	charset := Conf.DB.Charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "t_",
		},
	})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
