package group

import (
	"hmdl-user-service/db/core"
	"hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	repo "hmdl-user-service/repository/repoimpl"
)

func MenuRoute(api core.DbData) {

	g := api.Echo.Group("/menu")

	middlewares.SetJwtMiddlewares(g)

	handlerMenu := handler.MenuHandler{
		PhanQuyenMenuRepo: repo.NewPhanQuyenMenuRepo(api.DbSql01),
		MenuRepo:          repo.NewMenuWebRepo(api.DbSql01),
	}

	//Menu
	g.GET("", handlerMenu.GetAllMenu)
	g.POST("", handlerMenu.InsertMenu)
	g.PUT("", handlerMenu.UpdateMenu)
	g.GET("/id/:id", handlerMenu.GetAllMenuById)
	g.DELETE("/id/:id", handlerMenu.DeleteMenuById)
	api.Echo.GET("menu/id/:id", handlerMenu.GetAllMenuById)
	api.Echo.GET("/menu/", handlerMenu.GetMenuByPhanQuyenIdAnDuAnId)
	g.GET("/phan-quyen/", handlerMenu.GetMenuByPhanQuyenId)
}
