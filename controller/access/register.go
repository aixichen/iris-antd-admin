package access

import (
	"car-tms/controller"
	"car-tms/libs"
	"car-tms/libs/sms"
	"car-tms/models"
	"car-tms/validates"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func RegisterGetCode(ctx iris.Context) {
	aul := new(validates.RegisterCompanyCodeRequest)
	if err := ctx.ReadJSON(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}

	err := validates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				ctx.StatusCode(iris.StatusOK)
				_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", e, 2, ctx.GetID().(string)))
				return
			}
		}
	}

	code := libs.GetRandomInt(6)
	ctx.Application().Logger().Infof("%s-%s 注册系统-获取CODE", aul.Usermobile, code)
	ctx.StatusCode(iris.StatusOK)

	resultBool, err := sms.RegisterSmsSend(aul.Usermobile, code)

	if resultBool {
		_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "发送成功", 0, ctx.GetID().(string)))
	}
	_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
	return
}
func RegisterCompany(ctx iris.Context) {
	aul := new(validates.RegisterCompanyRequest)
	if err := ctx.ReadJSON(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}

	err := validates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				ctx.StatusCode(iris.StatusOK)
				_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", e, 2, ctx.GetID().(string)))
				return
			}
		}
	}

	resultBool, err := sms.CheckRegisterSmsCode(aul.Usermobile, aul.Captcha)
	if !resultBool {
		_, _ = ctx.JSON(controller.ApiResource(resultBool, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}

	ctx.Application().Logger().Infof("%s-%s 注册系统", aul.Companyname, aul.Username)
	ctx.StatusCode(iris.StatusOK)
	company := &models.Company{
		Companyname:    aul.Companyname,
		CompanyEndTime: time.Now().AddDate(0, 0, 3),
	}

	if err := company.CreateCompany(); err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}

	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "usermobile",
				Condition: "=",
				Value:     aul.Usermobile,
			},
		},
	}
	user, err := models.GetUser(s)
	if err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "查询失败", 4, ctx.GetID().(string)))
		models.DeleteCompanyById(company.ID)
		return
	}
	if user.ID > 0 {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "该手机号已注册", 4, ctx.GetID().(string)))
		models.DeleteCompanyById(company.ID)
		return
	}
	roleId := createAdminRole(company.ID)
	user.Username = aul.Username
	user.Usermobile = aul.Usermobile
	user.CompanyId = company.ID
	user.CompanyName = company.Companyname
	user.Password = aul.Password
	user.RoleIds = []uint{roleId}
	if err := user.CreateUser(); err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}

	_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "注册成功", 0, ctx.GetID().(string)))
	return
}

// CreateAdminRole 新建管理角色
func createAdminRole(companyId uint) uint {
	role := &models.Role{
		CompanyId:   companyId,
		Name:        "admin",
		DisplayName: "超级管理员",
		Description: "超级管理员",
		Model:       gorm.Model{CreatedAt: time.Now()},
	}
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "name",
				Condition: "=",
				Value:     role.Name,
			},
			{
				Key:       "company_id",
				Condition: "=",
				Value:     role.CompanyId,
			},
		},
	}
	searchRole, _ := models.GetRole(s)
	if searchRole.ID <= 0 {
		var permIds []uint
		perms, _ := models.GetAllPermissions(&models.Search{
			Fields:  []*models.Filed{{}},
			OrderBy: "",
			Sort:    "",
			Limit:   0,
			Offset:  0,
		})
		for _, perm := range perms {
			permIds = append(permIds, perm.ID)
		}

		role.PermIds = permIds
		if libs.Config.Debug {
			fmt.Println(fmt.Sprintf("填充角色数据：%v", role))
		}
		if err := role.CreateRole(); err != nil {
			logger.Println(fmt.Sprintf("管理角色填充错误：%v", err))
		}
	}
	return role.ID
}
