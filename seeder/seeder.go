package seeder

import (
	"car-tms/libs"
	"car-tms/models"
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"time"
)

var Seeds = struct {
	Perms []*models.Permission
}{}

func Run() {

	AutoMigrates()
	CreatePerms()
	fmt.Println(fmt.Sprintf("权限填充完成！！"))
}

// CreatePerms 新建权限
func CreatePerms() {
	if libs.Config.Debug {
		fmt.Println(fmt.Sprintf("填充权限：%v", Seeds))
	}

	for _, m := range Seeds.Perms {
		if !strings.HasPrefix(m.Name, "/api/auth/") {
			continue
		}
		perm := &models.Permission{
			Model:       gorm.Model{CreatedAt: time.Now()},
			Name:        m.Name,
			DisplayName: m.DisplayName,
			Description: m.Description,
			Act:         m.Act,
		}
		s := &models.Search{Fields: []*models.Filed{
			{
				Key:       "name",
				Condition: "=",
				Value:     m.Name,
			},
			{
				Key:       "act",
				Condition: "=",
				Value:     m.Act,
			},
		}}
		serachPerm, _ := models.GetPermission(s)
		if serachPerm.ID <= 0 {
			if err := perm.CreatePermission(); err != nil {
				logger.Println(fmt.Sprintf("权限填充错误：%v", err))
			}
		}
	}
}

/*
	AutoMigrates 重置数据表
	sysinit.Db.DropTableIfExists 删除存在数据表
	sysinit.Db.AutoMigrate 重建数据表
*/
func AutoMigrates() {
	libs.Db.AutoMigrate(
		&models.User{},
		&models.Sms{},
		&models.Role{},
		&models.Company{},
		&models.Permission{},
		&models.OauthToken{},
		&gormadapter.CasbinRule{},
		&models.Order{},
		&models.Office{},
		&models.Waybill{},
		&models.WaybillDetail{},
		&models.City{},
	)
}
