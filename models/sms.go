package models

import (
	"fmt"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"iris-antd-admin/libs"
	"time"
)

type Sms struct {
	gorm.Model
	SecretId         string `gorm:"not null VARCHAR(50)" json:"secret_id" validate:"required,gte=2,lte=50" comment:"SecretId"`
	Sign             string `gorm:"not null VARCHAR(50)" json:"sign" validate:"required,gte=2,lte=50" comment:"短信签名"`
	SdkAppid         string `gorm:"not null VARCHAR(50)" json:"sdk_app_id" validate:"required,gte=2,lte=50" comment:"SdkAppid"`
	Mobile           string
	TemplateID       string `gorm:"not null VARCHAR(50)" json:"template_id" validate:"required,gte=2,lte=50" comment:"TemplateID"`
	TemplateParamSet string
	SessionContext   string
	Status           uint //2过期
	Error            string
	Response         string
}

func NewSms() *Sms {
	return &Sms{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func GetSms(search *Search) (*Sms, error) {
	t := NewSms()
	err := Found(search).First(t).Error
	if !IsNotFound(err) {
		return t, err
	}
	return t, nil
}

/**
 * 通过 id 删除用户
 * @method DeleteUserById
 */
func DeleteSmsById(id uint) error {
	s := NewSms()
	s.ID = id
	if err := libs.Db.Delete(s).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteSmsByIdErr:%s \n", err))
		return err
	}
	return nil
}

/**
 * 创建
 * @method CreateUser
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (u *Sms) CreateSms() error {
	if err := libs.Db.Create(u).Error; err != nil {
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
func UpdateSmsById(id uint, s *Sms) error {
	if err := Update(&Sms{}, s, id); err != nil {
		return err
	}
	return nil
}
