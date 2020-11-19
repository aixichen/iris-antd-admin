package transformer

type Role struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	CompanyId      int    `json:"company_id"`
	DisplayName    string `json:"display_name"`
	Description    string `json:"description"`
	RolePermission []int  `json:"role_permission"`
	CreatedAt      string `json:"created_at"`
}

type Permission struct {
	Id          int    `json:"value"`
	Name        string `json:"name"`
	DisplayName string `json:"label"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type RoleCheckbox struct {
	Id          int    `json:"value"`
	Name        string `json:"name"`
	DisplayName string `json:"label"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
