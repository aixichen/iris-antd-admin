package controller

import (
	"github.com/kataras/iris/v12"
	"iris-antd-admin/models"
)

func GetCity(ctx iris.Context) {
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "parent_adcode",
				Condition: "=",
				Value:     "100000",
			},
		},
		OrderBy: "",
		Sort:    "",
		Limit:   0,
		Offset:  0,
	}

	cityData, err := models.GetAllCity(s)
	if err != nil {
		_, _ = ctx.JSON(ApiResource(false, nil, "10010", err.Error(), 4, ctx.GetID().(string)))
	} else {
		_, _ = ctx.JSON(ApiResource(true, setCity(cityData), "200", "", 0, ctx.GetID().(string)))
	}
	return
}
func setCity(cityData []*models.City) []map[string]interface{} {
	var resultData []map[string]interface{}
	var resultData1 []map[string]interface{}
	var resultData2 []map[string]interface{}
	var mapItem map[string]interface{}
	var mapItem1 map[string]interface{}
	var mapItem2 map[string]interface{}

	resultData = make([]map[string]interface{}, 0)
	for _, value := range cityData {
		mapItem = make(map[string]interface{})
		mapItem["label"] = value.Name
		mapItem["value"] = value.Adcode
		item1, _ := models.GetAllCity(&models.Search{Fields: []*models.Filed{{Key: "parent_adcode", Condition: "=", Value: value.Adcode}}})
		resultData1 = make([]map[string]interface{}, 0)
		for _, i := range item1 {
			mapItem1 = make(map[string]interface{})
			mapItem1["label"] = i.Name
			mapItem1["value"] = i.Adcode

			item2, _ := models.GetAllCity(&models.Search{Fields: []*models.Filed{{Key: "parent_adcode", Condition: "=", Value: i.Adcode}}})
			resultData2 = make([]map[string]interface{}, 0)
			for _, i2 := range item2 {
				mapItem2 = make(map[string]interface{})
				mapItem2["label"] = i2.Name
				mapItem2["value"] = i2.Adcode
				resultData2 = append(resultData2, mapItem2)
			}
			mapItem1["children"] = resultData2
			resultData1 = append(resultData1, mapItem1)
		}
		mapItem["children"] = resultData1
		resultData = append(resultData, mapItem)
	}
	return resultData
}
