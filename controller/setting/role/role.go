package role

import (
	"car-tms/controller"
	"car-tms/libs"
	"car-tms/models"
	"car-tms/transformer"
	"car-tms/validates"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	gf "github.com/snowlyg/gotransformer"
	"time"
)

func RoleList(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	name := ctx.FormValueDefault("name", "")

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
	s.Fields = append(s.Fields, models.GetSearche("name", name, "like"))

	roles, count, err := models.QueryPageRoles(s)
	if err != nil {
		ctx.JSON(controller.ApiResource(false, nil, "10001", "查询失败", 4, ctx.GetID().(string)))
	} else {
		ctx.JSON(controller.List{
			Data:     rolesTransform(roles),
			Success:  true,
			Total:    count,
			Current:  limit,
			PageSize: offset,
		})
	}
	return
}

func GetAllPermissions(ctx iris.Context) {

	perms, _ := models.GetAllPermissions(&models.Search{
		Fields:  []*models.Filed{{}},
		OrderBy: "",
		Sort:    "",
		Limit:   0,
		Offset:  0,
	})

	ctx.JSON(controller.ApiResource(true, permissionsTransform(perms), "200", "查询成功", 0, ctx.GetID().(string)))
	return
}

/**
* @api {post} /admin/roles/ 新建角色
* @apiName 新建角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 新建角色
* @apiSampleRequest /admin/roles/
* @apiParam {string} name 角色名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CreateRole(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	authUserCompanyId := ctx.Values().GetStringDefault("auth_user_company_id", "0")
	role := new(models.Role)
	role.CompanyId = uint(libs.ParseInt(authUserCompanyId, 0))

	if err := ctx.ReadJSON(role); err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}

	err := validates.Validate.Struct(*role)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", e, 2, ctx.GetID().(string)))
				return
			}
		}
	}

	err = role.CreateRole()
	if err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}
	if role.ID == 0 {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "创建失败", 2, ctx.GetID().(string)))
		return
	}
	ctx.JSON(controller.ApiResource(true, roleTransform(role), "200", "查询成功", 0, ctx.GetID().(string)))
	return

}

/**
* @api {post} /admin/roles/:id/update 更新角色
* @apiName 更新角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 更新角色
* @apiSampleRequest /admin/roles/:id/update
* @apiParam {string} name 角色名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UpdateRole(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	aul := new(validates.RoleUpdateRequest)

	if err := ctx.ReadJSON(aul); err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}

	err := validates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", e, 2, ctx.GetID().(string)))
				return
			}
		}
	}
	role, roleSearchErr := models.GetRoleById(aul.Id)
	if roleSearchErr != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}
	role.Name = aul.Name
	role.DisplayName = aul.DisplayName
	role.Description = aul.Description
	role.PermIds = aul.PermIds
	err = models.UpdateRole(role.ID, role)
	if err != nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}
	ctx.JSON(controller.ApiResource(true, roleTransform(role), "200", "查询成功", 0, ctx.GetID().(string)))
	return

}

/**
* @api {delete} /admin/roles/:id/delete 删除角色
* @apiName 删除角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 删除角色
* @apiSampleRequest /admin/roles/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func DeleteRole(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	aul := new(validates.RoleDeleteRequest)
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
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", "角色ID必填", 4, ctx.GetID().(string)))
		return
	}

	ctx.Application().Logger().Infof("%s 批量删除角色", ctx.Values().GetDefault("auth_user_id", ""))
	ctx.StatusCode(iris.StatusOK)
	for _, value := range aul.Ids {
		err := models.DeleteRoleById(value)
		if err != nil {
			_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", err.Error(), 4, ctx.GetID().(string)))
			return
		}

	}

	_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "处理成功", 0, ctx.GetID().(string)))
	return
}

func rolesTransform(roles []*models.Role) []*transformer.Role {
	var rs []*transformer.Role
	for _, role := range roles {
		r := roleTransform(role)
		rs = append(rs, r)
	}
	return rs
}

func roleTransform(role *models.Role) *transformer.Role {
	u := &transformer.Role{}
	g := gf.NewTransform(u, role, time.RFC3339)
	_ = g.Transformer()
	permissionArr := role.RolePermissions()
	var permissionId []int
	for _, value := range permissionArr {
		permissionId = append(permissionId, int(value.ID))
	}
	u.RolePermission = permissionId
	return u
}

func permissionsTransform(permissions []*models.Permission) []*transformer.Permission {
	var rs []*transformer.Permission
	for _, permission := range permissions {
		r := permissionTransform(permission)
		rs = append(rs, r)
	}
	return rs
}

func permissionTransform(permission *models.Permission) *transformer.Permission {
	u := &transformer.Permission{}
	g := gf.NewTransform(u, permission, time.RFC3339)
	_ = g.Transformer()
	return u
}
