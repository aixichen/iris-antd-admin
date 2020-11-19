package order

import (
	"car-tms/controller"
	"car-tms/validates"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

func OrderCreate(ctx iris.Context) {
	aul := new(validates.OrderRequest)
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

	authOfficeId := ctx.Values().GetUintDefault("auth_user_office_id", 0)
	if authOfficeId <= 0 {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "1001", "当前账号未绑定网点请绑定网点后操作", 2, ctx.GetID().(string)))
		return
	}

	ctx.Application().Logger().Infof("%s 创建订单", ctx.Values().GetDefault("auth_user_id", ""))
	ctx.StatusCode(iris.StatusOK)

}
