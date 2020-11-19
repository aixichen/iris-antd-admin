package access

import (
	"car-tms/controller"
	"car-tms/models"
	"car-tms/validates"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

func UserLogin(ctx iris.Context) {
	aul := new(validates.LoginRequest)
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

	ctx.Application().Logger().Infof("%s 登录系统", aul.Usermobile)
	ctx.StatusCode(iris.StatusOK)

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
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}

	if user.IsDisable == 1 {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "账号已被禁止登录,请联系管理员", 2, ctx.GetID().(string)))
		return
	}

	response, code, msg := user.CheckLogin(aul.Password)

	if code == 200 {
		_, _ = ctx.JSON(controller.ApiResource(true, response, "200", msg, 0, ctx.GetID().(string)))
	} else {
		_, _ = ctx.JSON(controller.ApiResource(false, response, "1001", msg, 2, ctx.GetID().(string)))
	}
	return
}
