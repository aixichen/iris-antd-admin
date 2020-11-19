package models

import (
	"gorm.io/gorm"
)

type CustomerPerson struct {
	gorm.Model
	CustomerPersonCrop      string //公司名称
	CustomerPersonName      string //姓名
	CustomerPersonMobile    string //电话
	CustomerPersonIdCode    string //身份证号
	CustomerPersonOther     string //其他信息
	CustomerPersonRemark    string
	CustomerPersonIsDisable int //0未禁用 1已禁用

	CreatedUserId      uint
	CreatedUserName    string
	CreatedUserMobile  string
	CreatedCompanyId   uint
	CreatedCompanyName string
}
