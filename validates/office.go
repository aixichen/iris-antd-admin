package validates

type (
	OfficeRequest struct {
		Id                 uint   `json:"id" comment:"网点ID"`
		OfficeName         string `json:"office_name" comment:"网点名称"`
		OfficeUsername     string `json:"office_username" comment:"网点负责人"`
		OfficeUserMobile   string `json:"office_user_mobile" comment:"网点联系电话"`
		OfficeProvinceName string `json:"office_province_name" comment:"网点:省"`
		OfficeProvinceCode string `json:"office_province_code" comment:"网点:省CODE"`
		OfficeCityName     string `json:"office_city_name" comment:"网点:市"`
		OfficeCityCode     string `json:"office_city_code" comment:"网点:市CODE"`
		OfficeAreaName     string `json:"office_area_name" comment:"网点:区"`
		OfficeAreaCode     string `json:"office_area_code" comment:"网点:区CODE"`
		OfficeAddress      string `json:"office_address" comment:"网点地址"`

		OfficeSignProvinceName string `json:"office_sign_province_name" comment:"网点签收地址:省"`
		OfficeSignProvinceCode string `json:"office_sign_province_code" comment:"网点签收地址:省CODE"`
		OfficeSignCityName     string `json:"office_sign_city_name" comment:"网点签收地址:市"`
		OfficeSignCityCode     string `json:"office_sign_city_code" comment:"网点签收地址:市CODE"`
		OfficeSignAreaName     string `json:"office_sign_area_name" comment:"网点签收地址:区"`
		OfficeSignAreaCode     string `json:"office_sign_area_code" comment:"网点签收地址:区CODE"`
		OfficeSignAddress      string `json:"office_sign_address" comment:"网点签收地址"`
		OfficeRemark           string `json:"office_remark"` //备注
	}
	OfficeDeleteRequest struct {
		Ids []uint `json:"ids" validate:"required" comment:"ID"`
	}
)
