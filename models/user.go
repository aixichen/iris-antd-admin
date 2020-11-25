package models

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"gorm.io/gorm"
	"iris-antd-admin/libs"
	"strconv"
	"time"
)

type User struct {
	gorm.Model
	CompanyId      uint
	CompanyName    string
	Username       string `gorm:"not null VARCHAR(50)" json:"username" validate:"required,gte=2,lte=50" comment:"用户名"`
	Usermobile     string `gorm:"not null VARCHAR(50)" json:"usermobile" comment:"用户手机号"`
	UserOfficeId   uint
	UserOfficeName string
	Password       string `gorm:"not null VARCHAR(50)" json:"password" validate:"required,gte=2,lte=50"  comment:"密码"`
	Intro          string `gorm:"not null VARCHAR(500)" json:"introduction" comment:"简介"`
	Avatar         string `gorm:"not null VARCHAR(500)" json:"avatar"  comment:"头像"`
	IsDisable      uint   `gorm:"int(1)" json:"is_disable" comment:"是否禁止登录 1是 0 否"`
	RoleIds        []uint `gorm:"-" json:"role_ids"  validate:"required" comment:"角色"`
}

func NewUser() *User {
	return &User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// GetUser get user
func GetUser(search *Search) (*User, error) {
	t := NewUser()
	err := Found(search).First(t).Error
	if !IsNotFound(err) {
		return t, err
	}
	return t, nil
}

// GetUser get user
func GetUserById(id uint) (*User, error) {
	t := NewUser()
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

// DeleteUser del user . if user's username is username ,can't del it.
func DeleteUser(id uint) error {
	s := &Search{
		Fields: []*Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}
	u, err := GetUser(s)
	if err != nil {
		return err
	}
	userId := strconv.FormatUint(uint64(u.ID), 10)
	if _, err := libs.Enforcer.DeleteRolesForUser(userId); err != nil {
		color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
	}

	if err := libs.Db.Delete(u, id).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteUserByIdErr:%s \n ", err))
		return err
	}
	return nil
}

// QueryPageUsers get all users
func QueryPageUsers(s *Search) ([]*User, int64, error) {
	var users []*User
	var count int64
	q := GetAll(&User{}, s)
	if err := q.Count(&count).Error; err != nil {
		return nil, count, err
	}
	q = q.Scopes(Paginate(s.Offset, s.Limit))
	if err := q.Find(&users).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllUserErr:%s \n ", err))
		return nil, count, err
	}
	return users, count, nil
}

// GetAllUsers get all users
func GetAllUsers(s *Search) ([]*User, error) {
	var users []*User
	q := GetAll(&User{}, s)
	if err := q.Find(&users).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllUserErr:%s \n ", err))
		return nil, err
	}
	return users, nil
}

// CreateUser create user
func (u *User) CreateUser() error {
	u.Password = libs.HashPassword(u.Password)
	if err := libs.Db.Create(u).Error; err != nil {
		return err
	}

	addRoles(u)

	return nil
}

// UpdateUserById update user by id
func UpdateUserById(id uint, nu *User, setPassWord bool) error {
	if setPassWord {
		nu.Password = libs.HashPassword(nu.Password)
	}
	if err := Update(&User{}, nu, id); err != nil {
		return err
	}

	addRoles(nu)
	return nil
}

func addRoles(user *User) {
	if len(user.RoleIds) > 0 {
		userId := strconv.FormatUint(uint64(user.ID), 10)
		if _, err := libs.Enforcer.DeleteRolesForUser(userId); err != nil {
			color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
		}

		for _, roleId := range user.RoleIds {
			roleId := strconv.FormatUint(uint64(roleId), 10)
			if _, err := libs.Enforcer.AddRoleForUser(userId, roleId); err != nil {
				color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
			}
		}
	}
}

/**
 * 判断用户是否登录
 * @method CheckLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func (u *User) CheckLogin(password string) (*Token, int64, string) {
	if u.ID == 0 {
		return nil, 400, "用户不存在"
	} else {
		if ok := bcrypt.Match(password, u.Password); ok {
			token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
				"iat": time.Now().Unix(),
			})
			tokenString, _ := token.SignedString([]byte("HS2JDFKhu7Y1av7b"))

			oauthToken := new(OauthToken)
			oauthToken.Token = tokenString
			oauthToken.UserId = u.ID
			oauthToken.Secret = "secret"
			oauthToken.Revoked = false
			oauthToken.ExpressIn = time.Now().Add(time.Hour * time.Duration(1)).Unix()
			oauthToken.CreatedAt = time.Now()

			response := oauthToken.OauthTokenCreate()

			return response, 200, "登陆成功"
		} else {
			return nil, 400, "用户名或密码错误"
		}
	}
}

/**
* 用户退出登陆
* @method UserAdminLogout
* @param  {[type]} ids string [description]
 */
func UserAdminLogout(userId uint) bool {
	ot := OauthToken{}
	ot.UpdateOauthTokenByUserId(userId)
	return ot.Revoked
}
