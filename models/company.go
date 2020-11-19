package models

import (
	"car-tms/libs"
	"fmt"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"time"
)

type Company struct {
	gorm.Model
	Companyname       string    `gorm:"not null VARCHAR(50)" json:"companyname" validate:"required,gte=2,lte=50" comment:"公司名称"`
	Companymobile     string    `gorm:"not null VARCHAR(11)" json:"companymobile" validate:"required,gte=2,lte=50" comment:"公司联系电话"`
	CompanySign       string    `gorm:"not null VARCHAR(50)" json:"company_sign" validate:"required,gte=2,lte=50" comment:"公司签名"`
	CompanyEndTime    time.Time `gorm:"not null VARCHAR(50)" json:"company_endtime" validate:"required,gte=2,lte=50" comment:"公司到期时间"`
	CompanyUserCount  uint      `gorm:"not null int(11)" json:"company_user_count"  comment:"允许"`
	CompanyUseVersion string    `gorm:"not null VARCHAR(50)" json:"company_sign" validate:"required,gte=2,lte=50" comment:"版本"`
}

func NewCompany() *Company {
	return &Company{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

/**
 * 通过 id 获取 Company 记录
 * @method GetRoleById
 * @param  {[type]}       role  *Role [description]
 */
func GetCompany(search *Search) (*Company, error) {
	c := NewCompany()
	err := Found(search).First(c).Error
	if !IsNotFound(err) {
		return c, err
	}
	return c, nil
}

/**
 * 创建
 * @method CreateUser
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (c *Company) CreateCompany() error {

	if err := libs.Db.Create(c).Error; err != nil {
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
func UpdateCompanyById(id uint, c *Company) error {
	if err := Update(&Company{}, c, id); err != nil {
		return err
	}
	return nil
}

/**
 * 通过 id 删除用户
 * @method DeleteUserById
 */
func DeleteCompanyById(id uint) error {
	c := NewCompany()
	c.ID = id
	if err := libs.Db.Delete(c).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteCompanyByIdError:%s \n", err))
		return err
	}
	return nil
}
