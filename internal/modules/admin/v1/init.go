package v1

import (
	"gin/internal/config"
	"gin/internal/middleware"
	"gin/internal/modules/admin/v1/ctrl"
	_ "gin/internal/modules/admin/v1/logic"
	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(router *gin.Engine) {

	router.Any("/v1/admin/version", func(ctx *gin.Context) {
		ctx.JSON(200, config.Version)
	})
	loginHandler := new(ctrl.LoginHandler)
	router.Any("/v1/admin/login", middleware.RequestLog("gin_admin"), loginHandler.Login)
	selectHandler := new(ctrl.SelectHandler)
	router.Any("/v1/admin/select/identify_list", selectHandler.GetIdentifySelectList)

	// todu init
	initIdentifyHandler := new(ctrl.IdentifyHandler)
	router.POST("/v1/admin/identify/init_create", initIdentifyHandler.InitCreate)
	userHandler := new(ctrl.UserHandler)
	//router.Any("user/create", userHandler.Create)
	router.Any("/v1/user/secret", userHandler.GetSecret)

	routerGroup := router.Group("/v1/admin")
	{
		routerGroup.Use(middleware.RequestLog("gin_admin"))
		routerGroup.Use(middleware.CheckToken())
		routerGroup.Use(middleware.CheckIdentify())
		routerGroup.Use(middleware.CheckPermission())

		routerGroup.Any("user/items", middleware.CheckPermission("web_user_list"), userHandler.Items)
		routerGroup.Any("user/info", middleware.CheckPermission("web_user_info"), userHandler.Info)
		routerGroup.Any("user/create", middleware.CheckPermission("web_user_add"), userHandler.Create)
		routerGroup.Any("user/update", middleware.CheckPermission("web_user_update"), userHandler.Update)
		routerGroup.Any("user/delete", middleware.CheckPermission("web_user_delete"), userHandler.Delete)

		roleHandler := new(ctrl.RoleHandler)
		routerGroup.Any("role/items", middleware.CheckPermission("web_role_list"), roleHandler.Items)
		routerGroup.Any("role/info", middleware.CheckPermission("web_role_info"), roleHandler.Info)
		routerGroup.Any("role/create", middleware.CheckPermission("web_role_add"), roleHandler.Create)
		routerGroup.Any("role/update", middleware.CheckPermission("web_role_update"), roleHandler.Update)
		routerGroup.Any("role/delete", middleware.CheckPermission("web_role_delete"), roleHandler.Delete)

		permissionHandler := new(ctrl.PermissionHandler)
		routerGroup.Any("permission/items", permissionHandler.Items)

		//routerGroup.Any("permission/user_add_permission", permissionHandler.UserAddPermission)
		routerGroup.Any("permission/role_add_permission", middleware.CheckPermission("web_role_change_permission"), permissionHandler.RoleAddPermission)
		//routerGroup.Any("permission/get_permission_by_user", permissionHandler.GetPermissionByUser)
		routerGroup.Any("permission/get_all_permission_by_user", permissionHandler.GetAllPermissionByUser)
		routerGroup.Any("permission/get_all_permission_by_role", permissionHandler.GetAllPermissionByRole)
		//routerGroup.Any("permission/get_menu_by_user", permissionHandler.GetMenuByUser)

		routerGroup.Any("select/role_select", selectHandler.GetRoleSelectList)
		routerGroup.Any("select/user_select_by_identify", selectHandler.GetUserSelectByIdentify)

		identifyHandler := new(ctrl.IdentifyHandler)
		routerGroup.Any("identify/items", middleware.CheckPermission("web_identity_list"), identifyHandler.Items)
		routerGroup.Any("identify/info", middleware.CheckPermission("web_identity_info"), identifyHandler.Info)
		routerGroup.Any("identify/create", middleware.CheckPermission("web_identity_add"), identifyHandler.Create)
		routerGroup.Any("identify/update", middleware.CheckPermission("web_identity_update"), identifyHandler.Update)
		routerGroup.Any("identify/delete", middleware.CheckPermission("web_identity_delete"), identifyHandler.Delete)

	}
}
