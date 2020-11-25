package models

import (
	"fmt"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"iris-antd-admin/libs"
	"strconv"
	"time"
)

type Permission struct {
	gorm.Model
	Name        string `gorm:"not null VARCHAR(191)" json:"name" validate:"required,gte=4,lte=50" comment:"名称"`
	DisplayName string `gorm:"VARCHAR(191)" json:"display_name" comment:"显示名称"`
	Description string `gorm:"VARCHAR(191)" json:"description" comment:"描述"`
	Act         string `gorm:"VARCHAR(191)" json:"act" comment:"Act"`
}

func NewPermission() *Permission {
	return &Permission{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

/**
 * 通过 id 获取 permission 记录
 * @method GetPermissionById
 * @param  {[type]}       permission  *Permission [description]
 */
func GetPermission(search *Search) (*Permission, error) {
	p := NewPermission()
	err := Found(search).First(p).Error

	if !IsNotFound(err) {
		return p, err
	}
	return p, nil
}

/**
 * 通过 id 删除权限
 * @method DeletePermissionById
 */
func DeletePermissionById(id uint) error {
	p := NewPermission()
	p.ID = id
	if err := libs.Db.Delete(p).Error; err != nil {
		color.Red(fmt.Sprintf("DeletePermissionByIdError:%s \n", err))
		return err
	}
	return nil
}

/**
 * 获取所有的权限
 * @method GetAllPermissions
 * @param  {[type]} name string [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func GetAllPermissions(search *Search) ([]*Permission, error) {
	var permissions []*Permission
	all := GetAll(&Permission{}, search)
	if err := all.Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

func GetPermissionForRole(uid uint) []string {
	uids, err := libs.Enforcer.GetRolesForUser(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		color.Red(fmt.Sprintf("GetRolesForUser 错误: %v", err))
		return []string{}
	}

	return uids
}

/**
 * 创建
 * @method CreatePermission
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (p *Permission) CreatePermission() error {
	if err := libs.Db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

/**
 * 更新
 * @method UpdatePermission
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func UpdatePermissionById(id uint, p *Permission) error {
	if err := Update(&Permission{}, p, id); err != nil {
		return err
	}
	return nil
}
