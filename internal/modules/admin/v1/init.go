package v1

import (
	"gin/internal/common_config"
	"gin/internal/common_middleware"
	"gin/internal/modules/admin/v1/ctrl_admin"
	"gin/internal/modules/admin/v1/ctrl_page"
	_ "gin/internal/modules/admin/v1/logic"
	"gin/internal/modules/admin/v1/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(router *gin.Engine) {

	router.Any("/v1/admin/version", func(ctx *gin.Context) {
		ctx.JSON(200, common_config.Version)
	})
	loginCtrl := new(ctrl_admin.LoginCtrl)
	router.Any("/v1/admin/login", common_middleware.RequestLog("gin_admin"), loginCtrl.Login)
	selectCtrl := new(ctrl_admin.SelectCtrl)
	router.Any("/v1/admin/select/identify_list", selectCtrl.GetIdentifySelectList)

	// todu init
	initIdentifyCtrl := new(ctrl_admin.IdentifyCtrl)
	router.POST("/v1/admin/identify/init_create", initIdentifyCtrl.InitCreate)
	userCtrl := new(ctrl_admin.UserCtrl)
	//router.Any("user/create", userCtrl.Create)
	router.Any("/v1/user/secret", userCtrl.GetSecret)

	routerGroup := router.Group("/v1/admin")
	{
		routerGroup.Use(common_middleware.RequestLog("gin_admin"))
		routerGroup.Use(middleware.CheckToken())
		routerGroup.Use(middleware.CheckIdentify())
		routerGroup.Use(middleware.CheckPermission())

		routerGroup.Any("user/items", middleware.CheckPermission("web_user_list"), userCtrl.Items)
		routerGroup.Any("user/info", middleware.CheckPermission("web_user_info"), userCtrl.Info)
		routerGroup.Any("user/create", middleware.CheckPermission("web_user_add"), userCtrl.Create)
		routerGroup.Any("user/update", middleware.CheckPermission("web_user_update"), userCtrl.Update)
		routerGroup.Any("user/delete", middleware.CheckPermission("web_user_delete"), userCtrl.Delete)

		roleCtrl := new(ctrl_admin.RoleCtrl)
		routerGroup.Any("role/items", middleware.CheckPermission("web_role_list"), roleCtrl.Items)
		routerGroup.Any("role/info", middleware.CheckPermission("web_role_info"), roleCtrl.Info)
		routerGroup.Any("role/create", middleware.CheckPermission("web_role_add"), roleCtrl.Create)
		routerGroup.Any("role/update", middleware.CheckPermission("web_role_update"), roleCtrl.Update)
		routerGroup.Any("role/delete", middleware.CheckPermission("web_role_delete"), roleCtrl.Delete)

		permissionCtrl := new(ctrl_admin.PermissionCtrl)
		routerGroup.Any("permission/items", permissionCtrl.Items)

		//routerGroup.Any("permission/user_add_permission", permissionCtrl.UserAddPermission)
		routerGroup.Any("permission/role_add_permission", middleware.CheckPermission("web_role_change_permission"), permissionCtrl.RoleAddPermission)
		//routerGroup.Any("permission/get_permission_by_user", permissionCtrl.GetPermissionByUser)
		routerGroup.Any("permission/get_all_permission_by_user", permissionCtrl.GetAllPermissionByUser)
		routerGroup.Any("permission/get_all_permission_by_role", permissionCtrl.GetAllPermissionByRole)
		//routerGroup.Any("permission/get_menu_by_user", permissionCtrl.GetMenuByUser)

		routerGroup.Any("select/role_select", selectCtrl.GetRoleSelectList)
		routerGroup.Any("select/user_select_by_identify", selectCtrl.GetUserSelectByIdentify)

		identifyCtrl := new(ctrl_admin.IdentifyCtrl)
		routerGroup.Any("identify/items", middleware.CheckPermission("web_identity_list"), identifyCtrl.Items)
		routerGroup.Any("identify/info", middleware.CheckPermission("web_identity_info"), identifyCtrl.Info)
		routerGroup.Any("identify/create", middleware.CheckPermission("web_identity_add"), identifyCtrl.Create)
		routerGroup.Any("identify/update", middleware.CheckPermission("web_identity_update"), identifyCtrl.Update)
		routerGroup.Any("identify/delete", middleware.CheckPermission("web_identity_delete"), identifyCtrl.Delete)

		// 业务
		siteCtrl := new(ctrl_page.SiteCtrl)
		routerGroup.Any("select/site", siteCtrl.SelectList)


	}
}
