package group

import (
	"hmdl-user-service/core"
	"hmdl-user-service/handler"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	repo "hmdl-user-service/repository/repoimpl"
)

func PhanQuyenRoute(api core.DbData) {

	g := api.Echo.Group("/phan-quyen")

	sdk.SetJwtMiddlewares(g)

	handlerPhanQuyen := handler.PhanQuyenHandler{
		PhanQuyenRepo: repo.NewDmPhanQuyenRepo(api.DbSql01),
	}

	//Phân quyền
	g.GET("", handlerPhanQuyen.GetAllPhanQuyen)

}
