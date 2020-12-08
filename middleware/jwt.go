package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"iris-antd-admin/controller"
)

/**
 * 验证 jwt
 * @method JwtHandler
 */
func JwtHandler() *jwt.Middleware {
	var mySecret = []byte("HS2JDFKhu7Y1av7b")
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx iris.Context, err error) {
			if ctx.GetCurrentRoute().Path() == "/api/auth/user/loginout" {
				ctx.JSON(controller.ApiResource(true, nil, "200", "登录过期,自动退出", 0, ctx.GetID().(string)))
			} else {
				ctx.JSON(controller.ApiResource(false, nil, "10001", "登录过期,请重新登录", 4, ctx.GetID().(string)))
			}

		},
	})
}
