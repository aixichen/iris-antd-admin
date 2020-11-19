package models

import (
	"car-tms/libs"
	"fmt"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"time"
)

type Office struct {
	gorm.Model
	OfficeName             string
	OfficeUsername         string
	OfficeUserMobile       string
	OfficeCompanyId        uint
	OfficeProvinceName     string
	OfficeProvinceCode     string
	OfficeCityName         string
	OfficeCityCode         string
	OfficeAreaName         string
	OfficeAreaCode         string
	OfficeAddress          string
	OfficeSignProvinceName string
	OfficeSignProvinceCode string
	OfficeSignCityName     string
	OfficeSignCityCode     string
	OfficeSignAreaName     string
	OfficeSignAreaCode     string
	OfficeSignAddress      string
	OfficeRemark           string
}

func NewOffice() *Office {
	return &Office{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func GetOffice(search *Search) (*Office, error) {
	o := NewOffice()
	err := Found(search).First(o).Error

	if !IsNotFound(err) {
		return o, err
	}
	return o, nil
}

/**
 * 通过 id 删除用户
 * @method DeleteUserById
 */
func DeleteOfficeById(id uint) error {
	o := NewOffice()
	o.ID = id
	if err := libs.Db.Delete(o).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteOfficeByIdError:%s \n", err))
		return err
	}
	return nil
}

/**
 * 获取所有的账号
 * @method QueryPageOffices
 * @param  {[type]} name string [description]
 * @param  {[type]} username string [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func QueryPageOffices(search *Search) ([]*Office, int64, error) {
	var offices []*Office
	var count int64
	all := GetAll(&Office{}, search)
	if err := all.Count(&count).Error; err != nil {
		return nil, count, err
	}

	all = all.Scopes(Paginate(search.Offset, search.Limit))

	if err := all.Find(&offices).Error; err != nil {
		return nil, count, err
	}

	return offices, count, nil
}

/**
 * 获取所有的账号
 * @method GetAllUser
 * @param  {[type]} name string [description]
 * @param  {[type]} username string [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func GetAllOffices(search *Search) ([]*Office, error) {
	var offices []*Office
	all := GetAll(&Office{}, search)

	if err := all.Find(&offices).Error; err != nil {
		return nil, err
	}

	return offices, nil
}

/**
 * 创建
 * @method CreateUser
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (o *Office) CreateOffice() error {

	if err := libs.Db.Create(o).Error; err != nil {
		color.Red(fmt.Sprintf("CreateOfficeError:%s \n", err))
		return err
	}

	return nil
}

/**
 * 更新
 * @method UpdateUser
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func UpdateOfficeById(id uint, o *Office) error {

	if err := Update(&Office{}, o, id); err != nil {
		return err
	}
	return nil
}
