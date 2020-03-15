package group

import (
	"hmdl-user-service/db/core"
	"hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	repo "hmdl-user-service/repository/repoimpl"
)

func PhanQuyenRoute(api core.DbData) {

	g := api.Echo.Group("/phanquyen")

	middlewares.SetJwtMiddlewares(g)

	handlerPhanQuyen := handler.PhanQuyenHandler{
		PhanQuyenRepo: repo.NewDmPhanQuyenRepo(api.Db),
	}

	//Phân quyền
	g.GET("", handlerPhanQuyen.GetAllPhanQuyen)


}
