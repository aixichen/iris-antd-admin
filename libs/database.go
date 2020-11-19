package libs

import (
	"errors"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	Db *gorm.DB
)

func InitDb() {
	var err error
	var conn string
	if Config.DB.Adapter == "mysql" {
		conn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", Config.DB.User, Config.DB.Password, Config.DB.Host, Config.DB.Port, Config.DB.Name)
	} else {
		logger.Println(errors.New("not supported database adapter"))
	}

	Db, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: gormlogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			gormlogger.Config{
				SlowThreshold: time.Second,     // 慢 SQL 阈值
				LogLevel:      gormlogger.Info, // Log level
				Colorful:      true,            // 禁用彩色打印

			},
		),
		NamingStrategy: schema.NamingStrategy{SingularTable: true, TablePrefix: Config.DB.Prefix},
	})

	if err != nil {
		logger.Println(err)
	}
}
