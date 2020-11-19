package user

import (
	"car-tms/controller"
	"car-tms/libs"
	"car-tms/models"
	"car-tms/transformer"
	"car-tms/validates"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	gf "github.com/snowlyg/gotransformer"
	"strconv"
	"time"
)

func UserList(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	username := ctx.FormValueDefault("username", "")
	usermobile := ctx.FormValueDefault("usermobile", "")
	user_office_name := ctx.FormValueDefault("user_office_name", "")
	is_disable := ctx.FormValueDefault("is_disable", "")

	offset := libs.ParseInt(ctx.FormValue("current"), 1)
	limit := libs.ParseInt(ctx.FormValue("pageSize"), 20)
	sorter := ctx.FormValue("sorter")
	authUserCompanyId := ctx.Values().GetStringDefault("auth_user_company_id", "0")
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "company_id",
				Condition: "=",
				Value:     authUserCompanyId,
			},
		},
		OrderBy: "",
		Sort:    "",
		Sorter:  sorter,
		Limit:   limit,
		Offset:  offset,
	}
	s.Fields = append(s.Fields, models.GetSearche("username", username, "like"))
	s.Fields = append(s.Fields, models.GetSearche("usermobile", usermobile, "like"))
	s.Fields = append(s.Fields, models.GetSearche("user_office_name", user_office_name, "like"))
	s.Fields = append(s.Fields, models.GetSearche("is_disable", is_disable, "="))

	users, count, err := models.QueryPageUsers(s)
	if err != nil {
		ctx.JSON(controller.ApiResource(false, nil, "10001", "查询失败", 4, ctx.GetID().(string)))
	} else {
		ctx.JSON(controller.List{
			Data:     usersTransform(users),
			Success:  true,
			Total:    count,
			Current:  limit,
			PageSize: offset,
		})
	}
	return
}

func CreateUser(ctx iris.Context) {
	aul := new(validates.UserCreateRequest)
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
		return
	}
	if user.ID > 0 {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "该手机号已注册", 4, ctx.GetID().(string)))
		return
	}
	authUserCompanyId := ctx.Values().GetStringDefault("auth_user_company_id", "0")
	authUserCompanyName := ctx.Values().GetStringDefault("auth_user_company_name", "")

	user.CompanyId = uint(libs.ParseInt(authUserCompanyId, 0))
	user.CompanyName = authUserCompanyName
	user.Username = aul.Username
	user.Usermobile = aul.Usermobile
	user.UserOfficeId = aul.UserOfficeId
	user.UserOfficeName = aul.UserOfficeName
	user.Password = aul.Password
	user.Intro = aul.Intro
	user.Avatar = aul.Avatar
	user.IsDisable = aul.IsDisable
	user.RoleIds = aul.RoleIds
	if err := user.CreateUser(); err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}

	_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "注册成功", 0, ctx.GetID().(string)))
	return
}

func UpdateUser(ctx iris.Context) {
	aul := new(validates.UserUpdateRequest)
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

	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "usermobile",
				Condition: "=",
				Value:     aul.Usermobile,
			},
			{
				Key:       "id",
				Condition: "!=",
				Value:     aul.Id,
			},
		},
	}

	user, err := models.GetUser(s)
	if err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "查询失败", 4, ctx.GetID().(string)))
		return
	}
	if user.ID > 0 {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "该手机号已注册", 4, ctx.GetID().(string)))
		return
	}
	authUserCompanyId := ctx.Values().GetStringDefault("auth_user_company_id", "0")
	authUserCompanyName := ctx.Values().GetStringDefault("auth_user_company_name", "")
	user, err = models.GetUserById(aul.Id)
	if err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "查询失败", 4, ctx.GetID().(string)))
		return
	}

	if user.ID <= 0 {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "数据验证失败", 4, ctx.GetID().(string)))
		return
	}
	user.CompanyId = uint(libs.ParseInt(authUserCompanyId, 0))
	user.CompanyName = authUserCompanyName
	user.Username = aul.Username
	user.Usermobile = aul.Usermobile
	user.UserOfficeId = aul.UserOfficeId
	user.UserOfficeName = aul.UserOfficeName

	user.Intro = aul.Intro
	user.Avatar = aul.Avatar
	user.IsDisable = aul.IsDisable
	user.RoleIds = aul.RoleIds
	setPassWord := false
	if len(aul.Password) > 0 {
		user.Password = aul.Password
		setPassWord = true
	}
	if err := models.UpdateUserById(aul.Id, user, setPassWord); err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}
	_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "注册成功", 0, ctx.GetID().(string)))
	return
}

func DeleteUser(ctx iris.Context) {
	aul := new(validates.UserDeleteRequest)
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
	if len(aul.Ids) <= 0 {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", "用户ID必填", 4, ctx.GetID().(string)))
		return
	}

	ctx.Application().Logger().Infof("%s 批量删除用户", ctx.Values().GetDefault("auth_user_id", ""))
	ctx.StatusCode(iris.StatusOK)
	for _, value := range aul.Ids {
		err := models.DeleteUser(value)
		if err != nil {
			_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", err.Error(), 4, ctx.GetID().(string)))
			return
		}

	}

	_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "处理成功", 0, ctx.GetID().(string)))
	return
}

func GetAllCompanyRole(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	authUserCompanyId := ctx.Values().GetStringDefault("auth_user_company_id", "0")
	roles, _ := models.GetAllRoles(&models.Search{
		Fields: []*models.Filed{
			{
				Key:       "company_id",
				Condition: "=",
				Value:     authUserCompanyId,
			},
		},
		OrderBy: "",
		Sort:    "",
		Limit:   0,
		Offset:  0,
	})

	ctx.JSON(controller.ApiResource(true, roleCheckboxsTransform(roles), "200", "查询成功", 0, ctx.GetID().(string)))
	return
}

func GetAllCompanyOffice(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	authUserCompanyId := ctx.Values().GetStringDefault("auth_user_company_id", "0")
	offices, _ := models.GetAllOffices(&models.Search{
		Fields: []*models.Filed{
			{
				Key:       "office_company_id",
				Condition: "=",
				Value:     authUserCompanyId,
			},
		},
		OrderBy: "",
		Sort:    "",
		Limit:   0,
		Offset:  0,
	})

	ctx.JSON(controller.ApiResource(true, officesSelectTransform(offices), "200", "查询成功", 0, ctx.GetID().(string)))
	return
}

func usersTransform(users []*models.User) []*transformer.User {
	var rs []*transformer.User
	for _, user := range users {
		r := userTransform(user)
		rs = append(rs, r)
	}
	return rs
}

func userTransform(user *models.User) *transformer.User {
	u := &transformer.User{}
	g := gf.NewTransform(u, user, time.RFC3339)
	_ = g.Transformer()

	roleIds := models.GetRolesForUser(user.ID)
	var ris []int
	var roleName []string

	var rolePermission []struct {
		Name string `json:"name"`
		Act  string `json:"act"`
	}
	for _, roleId := range roleIds {
		ri, _ := strconv.Atoi(roleId)
		ris = append(ris, ri)
		s := &models.Search{
			Fields: []*models.Filed{
				{
					Key:       "id",
					Condition: "=",
					Value:     ri,
				},
			},
		}
		role, _ := models.GetRole(s)

		roleName = append(roleName, role.Name)
		userPermission, _ := libs.Enforcer.GetImplicitPermissionsForUser(libs.ParseString(int(u.Id)))
		for _, value := range userPermission {
			rolePermission = append(rolePermission, struct {
				Name string `json:"name"`
				Act  string `json:"act"`
			}{Name: value[1], Act: value[2]})

		}
	}
	u.RoleIds = ris
	u.RoleName = roleName
	u.RolePermission = rolePermission

	return u
}

func roleCheckboxsTransform(roleCheckboxs []*models.Role) []*transformer.RoleCheckbox {
	var rs []*transformer.RoleCheckbox
	for _, ro := range roleCheckboxs {
		r := roleCheckboxTransform(ro)
		rs = append(rs, r)
	}
	return rs
}

func roleCheckboxTransform(RoleCheckbox *models.Role) *transformer.RoleCheckbox {
	u := &transformer.RoleCheckbox{}
	g := gf.NewTransform(u, RoleCheckbox, time.RFC3339)
	_ = g.Transformer()
	return u
}

func officesSelectTransform(offices []*models.Office) []*transformer.OfficeSelect {
	var rs []*transformer.OfficeSelect
	for _, ro := range offices {
		r := officeTransform(ro)
		rs = append(rs, r)
	}
	return rs
}

func officeTransform(office *models.Office) *transformer.OfficeSelect {
	u := &transformer.OfficeSelect{}
	g := gf.NewTransform(u, office, time.RFC3339)
	_ = g.Transformer()
	return u
}
