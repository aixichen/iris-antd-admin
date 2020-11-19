package models

import (
	"fmt"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name         string
	Adcode       string
	Citycode     string
	ParentAdcode string `gorm:"not null VARCHAR(50)" json:"parent_adcode" comment:"父级ID"`
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
func GetAllCity(search *Search) ([]*City, error) {
	var citys []*City
	q := GetAll(&City{}, nil)
	if err := q.Find(&citys).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllUserErr:%s \n ", err))
		return nil, err
	}
	return citys, nil
}
