package access

import (
	"car-tms/controller"
	"car-tms/libs"
	"car-tms/models"
	"car-tms/transformer"
	"github.com/kataras/iris/v12"
	gf "github.com/snowlyg/gotransformer"
	"strconv"
	"time"
)

/**
* @api {get} /admin/users/profile 获取登陆用户信息
* @apiName 获取登陆用户信息
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 获取登陆用户信息
* @apiSampleRequest /admin/users/profile
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetProfile(ctx iris.Context) {
	userId := ctx.Values().Get("auth_user_id").(uint)
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     userId,
			},
		},
	}
	user, _ := models.GetUser(s)
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(controller.ApiResource(true, userTransform(user), "200", "", 0, ctx.GetID().(string)))
}

/**
* @api {get} /admin/users/:id 根据id获取用户信息
* @apiName 根据id获取用户信息
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 根据id获取用户信息
* @apiSampleRequest /admin/users/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}
	user, _ := models.GetUser(s)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(controller.ApiResource(true, userTransform(user), "200", "操作成功", 0, ctx.GetID().(string)))
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
