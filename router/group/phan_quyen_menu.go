package group

import (
	"hmdl-user-service/db/core"
	"hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	repo "hmdl-user-service/repository/repoimpl"
)

func PhanQuyenMenuRoute(api core.DbData) {

	g := api.Echo.Group("/phan-quyen-menu")

	middlewares.SetJwtMiddlewares(g)

	handlerPhanQuyenMenu := handler.PhanQuyenMenuHandler{
		PhanQuyenMenuRepo: repo.NewPhanQuyenMenuRepo(api.DbSql01),
	}

	//Phân quyền Menu
	g.POST("/update-phan-quyen", handlerPhanQuyenMenu.UpdatePhanQuyenMenu)
	g.GET("", handlerPhanQuyenMenu.GetAllPhanQuyenMenu)
	g.POST("", handlerPhanQuyenMenu.InsertPhanQuyenMenu)
	g.GET("/id/:id", handlerPhanQuyenMenu.GetAllPhanQuyenMenuById)
	g.DELETE("/id/:id", handlerPhanQuyenMenu.DeletePhanQuyenMenuById)

}
