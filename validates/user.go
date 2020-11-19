package validates

type UserRequest struct {
	Username   string `json:"username" validate:"required,gte=2,lte=50" comment:"用户名"`
	Usermobile string `json:"usermobile" validate:"required,gte=2,lte=50" comment:"用户手机号"`

	UserOfficeId   uint   `json:"user_office_id" validate:"required" comment:"网点ID"`
	UserOfficeName string `json:"user_office_name" validate:"required" comment:"网点名称"`
	Intro          string `json:"intro" comment:"简介"`
	Avatar         string `json:"avatar" comment:"头像"`
	IsDisable      uint   `json:"is_disable" comment:"是否被禁用"`
	RoleIds        []uint `json:"role_ids"  validate:"required" comment:"角色"`
}

type UserCreateRequest struct {
	UserRequest
	Password string `json:"password" validate:"required"  comment:"密码"`
}

type UserUpdateRequest struct {
	Id       uint   `json:"id" validate:"required" comment:"员工ID"`
	Password string `json:"password"  comment:"密码"`
	UserRequest
}

type UserDeleteRequest struct {
	Ids []uint `json:"ids" validate:"required" comment:"ID"`
}
