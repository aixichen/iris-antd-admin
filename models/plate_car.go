package models

import (
	"gorm.io/gorm"
)

type PlateCar struct {
	gorm.Model
	PlateCarType             int    //0自营运输 2外协运输
	PlateCarIdCode           string //车牌
	PlateCarCrop             string //外协公司
	PlateCarCropPersonName   string //外协公司联系人
	PlateCarCropPersonMobile string //外协公司联系人电话
	PlateCarPersonName       string //司机姓名
	PlateCarPersonMobile     string //司机电话
	PlateCarPersonOther      string //司机其他信息
	PlateCarRemark           string
	PlateCarIsDisable        int //0未禁用 1已禁用

	CreatedUserId      uint
	CreatedUserName    string
	CreatedUserMobile  string
	CreatedCompanyId   uint
	CreatedCompanyName string
}
