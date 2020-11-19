package transformer

type Office struct {
	Id                 int    `json:"id"`
	OfficeName         string `json:"office_name"`
	OfficeUsername     string `json:"office_username"`
	OfficeUserMobile   string `json:"office_user_mobile"`
	OfficeCompanyId    uint   `json:"office_company_id"`
	OfficeProvinceName string `json:"office_province_name"`
	OfficeProvinceCode string `json:"office_province_code"`
	OfficeCityName     string `json:"office_city_name"`
	OfficeCityCode     string `json:"office_city_code"`
	OfficeAreaName     string `json:"office_area_name"`
	OfficeAreaCode     string `json:"office_area_code"`
	OfficeAddress      string `json:"office_address"`

	OfficeSignProvinceName string `json:"office_sign_province_name"`
	OfficeSignProvinceCode string `json:"office_sign_province_code"`
	OfficeSignCityName     string `json:"office_sign_city_name"`
	OfficeSignCityCode     string `json:"office_sign_city_code"`
	OfficeSignAreaName     string `json:"office_sign_area_name"`
	OfficeSignAreaCode     string `json:"office_sign_area_code"`
	OfficeSignAddress      string `json:"office_sign_address"`
	OfficeRemark           string `json:"office_remark"`

	CreatedAt string `json:"created_at"`
}

type OfficeSelect struct {
	Id         int    `json:"value"`
	OfficeName string `json:"label"`
}
