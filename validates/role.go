package validates

type (
	RoleRequest struct {
		Name        string `json:"name" validate:"required,gte=4,lte=50" comment:"名称"`
		DisplayName string `json:"display_name" comment:"显示名称"`
		Description string `json:"description" comment:"描述"`
		PermIds     []uint `json:"role_permission" comment:"权限id"`
	}
	RoleCreateRequest struct {
		RoleRequest
	}
	RoleUpdateRequest struct {
		RoleRequest
		Id uint `json:"id" validate:"required" comment:"ID必填"`
	}
	RoleDeleteRequest struct {
		Ids []uint `json:"ids" validate:"required" comment:"ID"`
	}
)
