package validates

type LoginRequest struct {
	Usermobile string `json:"usermobile" validate:"required,gte=2,lte=50" comment:"用户手机号"`
	Password   string `json:"password" validate:"required"  comment:"密码"`
}

type RegisterCompanyRequest struct {
	Username    string `json:"username" validate:"required,gte=2,lte=50" comment:"用户名"`
	Usermobile  string `json:"usermobile" validate:"required,gte=2,lte=50" comment:"用户手机号"`
	Captcha     string `json:"captcha" validate:"required,gte=6,lte=6" comment:"验证码"`
	Password    string `json:"password" validate:"required"  comment:"密码"`
	Companyname string `json:"companyname" validate:"required" comment:"公司名称"`
}

type RegisterCompanyCodeRequest struct {
	Usermobile string `json:"usermobile" validate:"required,gte=11,lte=11" comment:"用户手机号"`
}
