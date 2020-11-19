package transformer

type User struct {
	Id             int      `json:"id"`
	Username       string   `json:"username"`
	Usermobile     string   `json:"usermobile"`
	CompanyId      uint     `json:"company_id"`
	CompanyName    string   `json:"company_name"`
	UserOfficeId   uint     `json:"user_office_id"`
	UserOfficeName string   `json:"user_office_name"`
	Intro          string   `json:"introduction"`
	Avatar         string   `json:"avatar"`
	RoleName       []string `json:"roles"`
	RoleIds        []int    `json:"role_ids"`
	IsDisable      uint     `json:"is_disable"`
	RolePermission []struct {
		Name string `json:"name"`
		Act  string `json:"act"`
	} `json:"role_permission"`
	CreatedAt string `json:"created_at"`
}
