package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"iris-antd-admin/controller"
	"iris-antd-admin/controller/access"
	"iris-antd-admin/controller/setting/office"
	"iris-antd-admin/controller/setting/role"
	"iris-antd-admin/controller/setting/user"
	"iris-antd-admin/libs"
	"iris-antd-admin/middleware"
)

func App(api *iris.Application) {
	api.UseRouter(middleware.CrsAuth())
	api.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hi IRIS ANTD ADMIN")
	})
	api.PartyFunc("/api", func(v1 router.Party) {
		v1.PartyFunc("/access", func(acs router.Party) {
			acs.Post("/login", access.UserLogin).Name = "用户登录"
			acs.Post("/register", access.RegisterCompany).Name = "公司注册"
			acs.Post("/register/code", access.RegisterGetCode).Name = "公司注册获取CODE"
		})
		v1.Get("/city", controller.GetCity).Name = "获取地区信息"
		v1.PartyFunc("/auth", func(auth router.Party) {
			casbinMiddleware := middleware.New(libs.Enforcer)
			auth.Use(middleware.JwtHandler().Serve, casbinMiddleware.ServeHTTP)
			auth.PartyFunc("/user", func(us router.Party) {
				us.Get("/profile", access.GetProfile).Name = "基础能力-获取用户信息"
				us.Get("/loginout", access.UserLogout).Name = "基础能力-退出登录"
			})

			auth.PartyFunc("/setting", func(s router.Party) {

				s.PartyFunc("/office", func(o router.Party) {
					o.Get("/list", office.OfficeList).Name = "网点列表"
					o.Post("/create", office.OfficeCreate).Name = "创建网点"
					o.Post("/update", office.OfficeUpdate).Name = "更新网点"
					o.Post("/delete", office.OfficeDelete).Name = "删除网点"
				})

				s.PartyFunc("/role", func(r router.Party) {
					r.Get("/list", role.RoleList).Name = "角色列表"
					r.Get("/permission", role.GetAllPermissions).Name = "权限列表"
					r.Post("/create", role.CreateRole).Name = "添加角色"
					r.Post("/update", role.UpdateRole).Name = "修改角色"
					r.Post("/delete", role.DeleteRole).Name = "删除角色"
				})
				s.PartyFunc("/user", func(u router.Party) {
					u.Get("/list", user.UserList).Name = "用户列表"
					u.Get("/role", user.GetAllCompanyRole).Name = "用户列表获取角色"
					u.Get("/office", user.GetAllCompanyOffice).Name = "用户列表获取网点"

					u.Post("/create", user.CreateUser).Name = "创建用户"
					u.Post("/update", user.UpdateUser).Name = "修改用户"
					u.Post("/delete", user.DeleteUser).Name = "删除用户"

				})

			})
		})
	})
}
