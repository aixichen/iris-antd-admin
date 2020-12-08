package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"iris-antd-admin/controller"
	"iris-antd-admin/models"
	"net/http"
	"strconv"
	"time"
)

func New(e *casbin.Enforcer) *Casbin {
	return &Casbin{enforcer: e}
}

func (c *Casbin) ServeHTTP(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	token := models.OauthToken{}
	token.GetOauthTokenByToken(value.Raw) //获取 access_token 信息
	if token.Revoked || token.ExpressIn < time.Now().Unix() {
		//_, _ = ctx.Writef("token 失效，请重新登录") // 输出到前端
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(controller.ApiResource(false, nil, "10001", "登录失效，请重新登录", 4, ctx.GetID().(string)))
		ctx.StopExecution()
		return
	} else if !c.Check(ctx.Request(), strconv.FormatUint(uint64(token.UserId), 10)) {
		ctx.StatusCode(http.StatusForbidden) // Status Forbidden
		ctx.JSON(controller.ApiResource(false, nil, "10001", "权限验证失败", 4, ctx.GetID().(string)))
		ctx.StopExecution()
		return
	} else {
		user, _ := models.GetUserById(token.UserId)
		if len(user.Usermobile) <= 0 {
			ctx.StatusCode(http.StatusUnauthorized)
			ctx.StopExecution()
			return
		}
		ctx.Values().Set("auth_user_id", token.UserId)
		ctx.Values().Set("auth_user_company_id", user.CompanyId)
		ctx.Values().Set("auth_user_company_name", user.CompanyName)
		ctx.Values().Set("auth_user_office_id", user.UserOfficeId)
		ctx.Values().Set("auth_user_office_name", user.UserOfficeName)
	}

	ctx.Next()
}

// Casbin is the auth services which contains the casbin enforcer.
type Casbin struct {
	enforcer *casbin.Enforcer
}

// Check checks the username, request's method and path and
// returns true if permission grandted otherwise false.
func (c *Casbin) Check(r *http.Request, userId string) bool {
	method := r.Method
	path := r.URL.Path
	fmt.Println(path)
	ok, _ := c.enforcer.Enforce(userId, path, method)
	return ok
}
