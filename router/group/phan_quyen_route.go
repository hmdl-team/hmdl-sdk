package group

import (
	"hmdl-user-service/db/core"
	"hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	repo "hmdl-user-service/repository/repoimpl"
)

func PhanQuyenRoute(api core.DbData) {

	g := api.Echo.Group("/phan-quyen")

	middlewares.SetJwtMiddlewares(g)

	handlerPhanQuyen := handler.PhanQuyenHandler{
		PhanQuyenRepo: repo.NewDmPhanQuyenRepo(api.DbSql01),
	}

	//Phân quyền
	g.GET("", handlerPhanQuyen.GetAllPhanQuyen)

}
