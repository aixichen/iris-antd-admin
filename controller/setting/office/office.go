package office

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

func OfficeCreate(ctx iris.Context) {
	aul := new(validates.OfficeRequest)
	authUserCompanyId := ctx.Values().GetStringDefault("auth_user_company_id", "0")
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

	ctx.Application().Logger().Infof("%s 创建网点", ctx.Values().GetDefault("auth_user_id", ""))
	ctx.StatusCode(iris.StatusOK)

	office := models.NewOffice()
	office.OfficeCompanyId = uint(libs.ParseInt(authUserCompanyId, 0))
	office.OfficeName = aul.OfficeName
	office.OfficeUsername = aul.OfficeUsername
	office.OfficeUserMobile = aul.OfficeUserMobile
	office.OfficeProvinceName = aul.OfficeProvinceName
	office.OfficeProvinceCode = aul.OfficeProvinceCode
	office.OfficeCityName = aul.OfficeCityName
	office.OfficeCityCode = aul.OfficeCityCode
	office.OfficeAreaName = aul.OfficeAreaName
	office.OfficeAreaCode = aul.OfficeAreaCode
	office.OfficeAddress = aul.OfficeAddress

	office.OfficeSignProvinceName = aul.OfficeSignProvinceName
	office.OfficeSignProvinceCode = aul.OfficeSignProvinceCode
	office.OfficeSignCityName = aul.OfficeSignCityName
	office.OfficeSignCityCode = aul.OfficeSignCityCode
	office.OfficeSignAreaName = aul.OfficeSignAreaName
	office.OfficeSignAreaCode = aul.OfficeSignAreaCode
	office.OfficeSignAddress = aul.OfficeSignAddress
	office.OfficeRemark = aul.OfficeRemark
	resultErr := office.CreateOffice()
	if resultErr == nil {
		_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "处理成功", 0, ctx.GetID().(string)))
		return
	} else {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", resultErr.Error(), 4, ctx.GetID().(string)))
		return
	}
}

func OfficeUpdate(ctx iris.Context) {
	aul := new(validates.OfficeRequest)
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
	if aul.Id <= 0 {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", "网点ID必填", 4, ctx.GetID().(string)))
		return
	}

	ctx.Application().Logger().Infof("%s 修改网点", ctx.Values().GetDefault("auth_user_id", ""))
	ctx.StatusCode(iris.StatusOK)
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     aul.Id,
			},
		},
	}

	office, _ := models.GetOffice(s)
	if office == nil {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", "未查到相关数据", 4, ctx.GetID().(string)))
		return
	}
	office.OfficeName = aul.OfficeName
	office.OfficeUsername = aul.OfficeUsername
	office.OfficeUserMobile = aul.OfficeUserMobile
	office.OfficeProvinceName = aul.OfficeProvinceName
	office.OfficeProvinceCode = aul.OfficeProvinceCode
	office.OfficeCityName = aul.OfficeCityName
	office.OfficeCityCode = aul.OfficeCityCode
	office.OfficeAreaName = aul.OfficeAreaName
	office.OfficeAreaCode = aul.OfficeAreaCode
	office.OfficeAddress = aul.OfficeAddress

	office.OfficeSignProvinceName = aul.OfficeSignProvinceName
	office.OfficeSignProvinceCode = aul.OfficeSignProvinceCode
	office.OfficeSignCityName = aul.OfficeSignCityName
	office.OfficeSignCityCode = aul.OfficeSignCityCode
	office.OfficeSignAreaName = aul.OfficeSignAreaName
	office.OfficeSignAreaCode = aul.OfficeSignAreaCode
	office.OfficeSignAddress = aul.OfficeSignAddress
	office.OfficeRemark = aul.OfficeRemark
	resultErr := models.UpdateOfficeById(aul.Id, office)
	if resultErr == nil {
		_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "处理成功", 0, ctx.GetID().(string)))
		return
	} else {
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", resultErr.Error(), 4, ctx.GetID().(string)))
		return
	}
}

func OfficeList(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	office_name := ctx.FormValueDefault("office_name", "")
	office_username := ctx.FormValueDefault("office_username", "")
	office_user_mobile := ctx.FormValueDefault("office_user_mobile", "")

	offset := libs.ParseInt(ctx.FormValue("current"), 1)
	limit := libs.ParseInt(ctx.FormValue("pageSize"), 20)
	authUserCompanyId := ctx.Values().GetStringDefault("auth_user_company_id", "0")
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "office_company_id",
				Condition: "=",
				Value:     authUserCompanyId,
			},
		},
		OrderBy: "",
		Sort:    "",
		Limit:   limit,
		Offset:  offset,
	}
	s.Fields = append(s.Fields, models.GetSearche("office_name", office_name, "like"))
	s.Fields = append(s.Fields, models.GetSearche("office_name", office_username, "like"))
	s.Fields = append(s.Fields, models.GetSearche("office_name", office_user_mobile, "like"))

	offices, count, err := models.QueryPageOffices(s)
	if err != nil {
		ctx.JSON(controller.ApiResource(false, nil, "10001", "查询失败", 4, ctx.GetID().(string)))
	} else {
		ctx.JSON(controller.List{
			Data:     officesTransform(offices),
			Success:  true,
			Total:    count,
			Current:  limit,
			PageSize: offset,
		})
	}
	return
}

func OfficeDelete(ctx iris.Context) {
	aul := new(validates.OfficeDeleteRequest)
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
		_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", "网点ID必填", 4, ctx.GetID().(string)))
		return
	}

	ctx.Application().Logger().Infof("%s 批量删除网点", ctx.Values().GetDefault("auth_user_id", ""))
	ctx.StatusCode(iris.StatusOK)
	for _, value := range aul.Ids {
		err := models.DeleteOfficeById(value)
		if err != nil {
			_, _ = ctx.JSON(controller.ApiResource(false, nil, "10001", err.Error(), 4, ctx.GetID().(string)))
			return
		}

	}

	_, _ = ctx.JSON(controller.ApiResource(true, nil, "200", "处理成功", 0, ctx.GetID().(string)))
	return
}

func officesTransform(offices []*models.Office) []*transformer.Office {
	var rs []*transformer.Office
	for _, office := range offices {
		r := officeTransform(office)
		rs = append(rs, r)
	}
	return rs
}

func officeTransform(office *models.Office) *transformer.Office {
	u := &transformer.Office{}
	g := gf.NewTransform(u, office, time.RFC3339)
	_ = g.Transformer()
	return u
}
