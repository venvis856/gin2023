package router

import (
	"gin/app/http/controllers/admin"
	"gin/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(router *gin.Engine) {
	loginController := new(admin.LoginController)
	router.Any("/admin/login", loginController.Login)

	// 预览
	cmsPreviewController := new(admin.CmsPreviewController)
	router.Any("page/preview", cmsPreviewController.Preview)
	userController := new(admin.UserController)
	router.Any("admin/user/initcreate", userController.Create)

	videoController2 := new(admin.VideoController)
	router.Any("/video/items", videoController2.WSearchItems)

	routerGroup := router.Group("/admin")
	{
		routerGroup.Use(middleware.CheckToken())

		routerGroup.Any("user/items", userController.Items)
		routerGroup.Any("user/info", userController.Info)
		routerGroup.Any("user/create", userController.Create)
		routerGroup.Any("user/update", userController.Update)
		routerGroup.Any("user/delete", userController.Delete)

		nodeBookController := new(admin.NoteBookController)
		routerGroup.Any("notebook/items", nodeBookController.Items)
		routerGroup.Any("notebook/info", nodeBookController.Info)
		routerGroup.Any("notebook/create_or_update", nodeBookController.CreateOrUpdate)
		routerGroup.Any("notebook/delete", nodeBookController.Delete)

		roleController := new(admin.RoleController)
		routerGroup.Any("role/items", roleController.Items)
		routerGroup.Any("role/info", roleController.Info)
		routerGroup.Any("role/create", roleController.Create)
		routerGroup.Any("role/update", roleController.Update)
		routerGroup.Any("role/delete", roleController.Delete)

		permissionController := new(admin.PermissionController)
		routerGroup.Any("permission/items", permissionController.Items)
		routerGroup.Any("permission/create", permissionController.Create)
		routerGroup.Any("permission/update", permissionController.Update)
		routerGroup.Any("permission/delete", permissionController.Delete)
		routerGroup.Any("permission/user_add_permission", permissionController.UserAddPermission)
		routerGroup.Any("permission/user_delete_permission", permissionController.UserDeletePermission)
		routerGroup.Any("permission/role_add_permission", permissionController.RoleAddPermission)
		//routerGroup.Any("permission/role_delete_permission", permissionController.RoleDeletePermission)
		routerGroup.Any("permission/get_permission_by_user", permissionController.GetPermissionByUser)
		routerGroup.Any("permission/get_all_permission_by_user", permissionController.GetAllPermissionByUser)
		routerGroup.Any("permission/get_all_permission_by_role", permissionController.GetAllPermissionByRole)
		routerGroup.Any("permission/get_menu_by_user", permissionController.GetMenuByUser)

		// cms
		cmsAuthorController := new(admin.CmsAuthorController)
		routerGroup.Any("author/items", cmsAuthorController.Items)
		routerGroup.Any("author/info", cmsAuthorController.Info)
		routerGroup.Any("author/create", cmsAuthorController.Create)
		routerGroup.Any("author/update", cmsAuthorController.Update)
		routerGroup.Any("author/delete", cmsAuthorController.Delete)

		cmsClassifyController := new(admin.CmsClassifyController)
		routerGroup.Any("classify/items", cmsClassifyController.Items)
		routerGroup.Any("classify/info", cmsClassifyController.Info)
		routerGroup.Any("classify/create", cmsClassifyController.Create)
		routerGroup.Any("classify/update", cmsClassifyController.Update)
		routerGroup.Any("classify/delete", cmsClassifyController.Delete)

		cmsModuleController := new(admin.CmsModuleController)
		routerGroup.Any("module/items", cmsModuleController.Items)
		routerGroup.Any("module/info", cmsModuleController.Info)
		routerGroup.Any("module/create", cmsModuleController.Create)
		routerGroup.Any("module/update", cmsModuleController.Update)
		routerGroup.Any("module/delete", cmsModuleController.Delete)

		cmsPageController := new(admin.CmsPageController)
		routerGroup.Any("page/items", cmsPageController.Items)
		routerGroup.Any("page/info", cmsPageController.Info)
		routerGroup.Any("page/create", cmsPageController.Create)
		routerGroup.Any("page/update", cmsPageController.Update)
		routerGroup.Any("page/article_create", cmsPageController.ArticleCreate)
		routerGroup.Any("page/article_update", cmsPageController.ArticleUpdate)
		routerGroup.Any("page/delete", cmsPageController.Delete)

		cmsProductController := new(admin.CmsProductController)
		routerGroup.Any("product/items", cmsProductController.Items)
		routerGroup.Any("product/info", cmsProductController.Info)
		routerGroup.Any("product/create", cmsProductController.Create)
		routerGroup.Any("product/update", cmsProductController.Update)
		routerGroup.Any("product/delete", cmsProductController.Delete)

		cmsTemplateController := new(admin.CmsTemplateController)
		routerGroup.Any("template/items", cmsTemplateController.Items)
		routerGroup.Any("template/info", cmsTemplateController.Info)
		routerGroup.Any("template/create", cmsTemplateController.Create)
		routerGroup.Any("template/update", cmsTemplateController.Update)
		routerGroup.Any("template/delete", cmsTemplateController.Delete)

		cmsSelectController := new(admin.CmsSelectController)
		routerGroup.Any("select/getPermissionByUser", cmsSelectController.GetPermissionByUser)
		routerGroup.Any("select/site", cmsSelectController.GetSiteSelectList)
		routerGroup.Any("select/template", cmsSelectController.GetTemplateSelectList)
		routerGroup.Any("select/module", cmsSelectController.GetModuleSelectList)
		routerGroup.Any("select/classify", cmsSelectController.GetClassifySelectList)
		routerGroup.Any("select/author", cmsSelectController.GetAuthorSelectList)
		routerGroup.Any("select/product", cmsSelectController.GetProductSelectList)
		routerGroup.Any("select/tag", cmsSelectController.GetTagSelectList)
		routerGroup.Any("select/video_tag", cmsSelectController.GetVideoTagSelectList)

		cmsPictureController := new(admin.CmsPictureController)
		routerGroup.Any("picture/upload", cmsPictureController.Upload)
		routerGroup.Any("picture/upload_and_publish", cmsPictureController.UploadAndPublish)
		routerGroup.Any("picture/get_file_list", cmsPictureController.GetFileList)
		routerGroup.Any("picture/get_dir_list", cmsPictureController.GetDirList)
		routerGroup.Any("picture/create_dir", cmsPictureController.CreateDir)
		routerGroup.Any("picture/change_online_pic", cmsPictureController.ChangeOnlinePic)

		cmsMakeController := new(admin.CmsMakeController)
		routerGroup.Any("make/page_make", cmsMakeController.PageMake)
		routerGroup.Any("make/template_make", cmsMakeController.TemplateMake)
		routerGroup.Any("make/module_make", cmsMakeController.ModuleMake)
		routerGroup.Any("make/video_tag", cmsMakeController.VideoTagMake)

		cmsAuditController := new(admin.CmsAuditController)
		routerGroup.Any("audit/items", cmsAuditController.Items)
		routerGroup.Any("audit/delete", cmsAuditController.Delete)

		cmsAuditDetailController := new(admin.CmsAuditDetailController)
		routerGroup.Any("audit_detail/items", cmsAuditDetailController.Items)
		routerGroup.Any("audit_detail/delete", cmsAuditDetailController.Delete)

		cmsPublishController := new(admin.CmsPublishController)
		routerGroup.Any("publish/publish", cmsPublishController.Publish)

		cmsTagController := new(admin.CmsTagController)
		routerGroup.Any("tag/items", cmsTagController.Items)
		routerGroup.Any("tag/info", cmsTagController.Info)
		routerGroup.Any("tag/create", cmsTagController.Create)
		routerGroup.Any("tag/update", cmsTagController.Update)
		routerGroup.Any("tag/delete", cmsTagController.Delete)

		videoController := new(admin.VideoController)
		routerGroup.Any("video/items", videoController.Items)
		routerGroup.Any("video/info", videoController.Info)
		routerGroup.Any("video/create", videoController.Create)
		routerGroup.Any("video/update", videoController.Update)
		routerGroup.Any("video/delete", videoController.Delete)

		videoTagController := new(admin.VideoTagController)
		routerGroup.Any("video_tag/items", videoTagController.Items)
		routerGroup.Any("video_tag/info", videoTagController.Info)
		routerGroup.Any("video_tag/create", videoTagController.Create)
		routerGroup.Any("video_tag/update", videoTagController.Update)
		routerGroup.Any("video_tag/delete", videoTagController.Delete)

	}
}
