package group

import (
	"hmdl-user-service/db/core"
	"hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	repo "hmdl-user-service/repository/repoimpl"
)

func PhanQuyenMenuRoute(api core.DbData) {

	g := api.Echo.Group("/phanquyenmenu")

	middlewares.SetJwtMiddlewares(g)

	handlerPhanQuyenMenu := handler.PhanQuyenMenuHandler{
		PhanQuyenMenuRepo: repo.NewPhanQuyenMenuRepo(api.Db),
	}

	//Phân quyền Menu
	g.GET("", handlerPhanQuyenMenu.GetAllPhanQuyenMenu)
	g.POST("", handlerPhanQuyenMenu.InsertPhanQuyenMenu)
	g.GET("/id/:id", handlerPhanQuyenMenu.GetAllPhanQuyenMenuById)
	g.DELETE("/id/:id", handlerPhanQuyenMenu.DeletePhanQuyenMenuById)

}
