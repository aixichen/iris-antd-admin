package access

import (
	"github.com/kataras/iris/v12"
	"iris-antd-admin/controller"
	"iris-antd-admin/libs"
	"iris-antd-admin/models"
	"net/http"
)

/**
* @api {get} /logout 用户退出登陆
* @apiName 用户退出登陆
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 用户退出登陆
* @apiSampleRequest /logout
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UserLogout(ctx iris.Context) {
	aui := ctx.Values().GetString("auth_user_id")
	uid := uint(libs.ParseInt(aui, 0))
	models.UserAdminLogout(uid)

	ctx.Application().Logger().Infof("%d 退出系统", uid)
	ctx.StatusCode(http.StatusOK)
	_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "退出成功", 0, ctx.GetID().(string)))
}
