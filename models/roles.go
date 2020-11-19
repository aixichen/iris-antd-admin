package models

import (
	"car-tms/libs"
	"fmt"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Role struct {
	gorm.Model
	CompanyId   uint   `gorm:"int(11)" json:"company_id" comment:"公司ID"`
	Name        string `gorm:"not null VARCHAR(191)" json:"name" validate:"required,gte=4,lte=50" comment:"名称"`
	DisplayName string `gorm:"VARCHAR(191)" json:"display_name" comment:"显示名称"`
	Description string `gorm:"VARCHAR(191)" json:"description" comment:"描述"`
	PermIds     []uint `gorm:"-" json:"role_permission" comment:"权限id"`
}

func NewRole() *Role {
	return &Role{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// GetRole get role
func GetRole(search *Search) (*Role, error) {
	t := NewRole()
	err := Found(search).First(t).Error
	if !IsNotFound(err) {
		return t, err
	}
	return t, nil
}

// GetRole get role
func GetRoleById(id uint) (*Role, error) {
	t := NewRole()
	s := &Search{
		Fields: []*Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}

	err := Found(s).First(t).Error
	if !IsNotFound(err) {
		return t, err
	}
	return t, nil
}

// DeleteRoleById del role by id
func DeleteRoleById(id uint) error {
	r := NewRole()
	r.ID = id
	roleId := strconv.FormatUint(uint64(id), 10)
	if _, err := libs.Enforcer.DeleteRole(roleId); err != nil {
		color.Red(fmt.Sprintf("DeleteRoleErr:%s \n", err))
	}
	err := libs.Db.Delete(r).Error
	if err != nil {
		color.Red(fmt.Sprintf("DeleteRoleErr:%s \n", err))
		return err
	}

	return nil
}

// QueryPageRoles get all roles
func QueryPageRoles(s *Search) ([]*Role, int64, error) {
	var roles []*Role
	var count int64
	all := GetAll(&Role{}, s)
	if err := all.Count(&count).Error; err != nil {
		return nil, count, err
	}
	all = all.Scopes(Paginate(s.Offset, s.Limit))
	if err := all.Find(&roles).Error; err != nil {
		return nil, count, err
	}
	return roles, count, nil
}

// GetAllRoles get all roles
func GetAllRoles(s *Search) ([]*Role, error) {
	var roles []*Role
	all := GetAll(&Role{}, s)
	if err := all.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// CreateRole create role
func (r *Role) CreateRole() error {
	if err := libs.Db.Create(r).Error; err != nil {
		return err
	}

	addPerms(r.PermIds, r)

	return nil
}

// addPerms add perms
func addPerms(permIds []uint, role *Role) {
	if len(permIds) > 0 {
		roleId := strconv.FormatUint(uint64(role.ID), 10)
		if _, err := libs.Enforcer.DeletePermissionsForUser(roleId); err != nil {
			color.Red(fmt.Sprintf("AppendPermsErr:%s \n", err))
		}
		var perms []Permission
		libs.Db.Where("id in (?)", permIds).Find(&perms)
		for _, perm := range perms {
			if _, err := libs.Enforcer.AddPolicy(roleId, perm.Name, perm.Act); err != nil {
				color.Red(fmt.Sprintf("AddPolicy:%s \n", err))
			}
		}
	} else {
		color.Yellow(fmt.Sprintf("没有角色：%s 权限为空 \n", role.Name))
	}
}

// UpdateRole update role
func UpdateRole(id uint, nr *Role) error {
	if err := Update(&Role{}, nr, id); err != nil {
		return err
	}

	addPerms(nr.PermIds, nr)

	return nil
}

// RolePermissions get role's permissions
func (r *Role) RolePermissions() []*Permission {
	perms := GetPermissionsForUser(r.ID)
	var ps []*Permission
	for _, perm := range perms {
		if len(perm) >= 3 && len(perm[1]) > 0 && len(perm[2]) > 0 {
			s := &Search{
				Fields: []*Filed{
					{
						Key:       "name",
						Condition: "=",
						Value:     perm[1],
					},
					{
						Key:       "act",
						Condition: "=",
						Value:     perm[2],
					},
				},
			}
			p, err := GetPermission(s)
			if err == nil && p.ID > 0 {
				ps = append(ps, p)
			}
		}
	}
	return ps
}
