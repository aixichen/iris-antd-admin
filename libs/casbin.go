package libs

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	logger "github.com/sirupsen/logrus"
	"path/filepath"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	var err error
	var conn string
	if Config.DB.Adapter == "mysql" {
		conn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", Config.DB.User, Config.DB.Password, Config.DB.Host, Config.DB.Port, Config.DB.Name)
	} else {
		logger.Println(errors.New("not supported database adapter"))
	}

	if len(conn) == 0 {
		logger.Println(fmt.Sprintf("数据链接不可用: %s", conn))
	}

	c, err := gormadapter.NewAdapter(Config.DB.Adapter, conn, true) // Your driver and data source.
	if err != nil {
		logger.Println(fmt.Sprintf("NewAdapter 错误: %v,Path: %s", err, conn))
	}

	casbinModelPath := filepath.Join(CWD(), "rbac_model.conf")
	Enforcer, err = casbin.NewEnforcer(casbinModelPath, c)

	if err != nil {
		logger.Println(fmt.Sprintf("NewEnforcer 错误: %v", err))
	}

	_ = Enforcer.LoadPolicy()

}
