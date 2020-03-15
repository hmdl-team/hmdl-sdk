package group

import (
	"hmdl-user-service/db/core"
	"hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	repo "hmdl-user-service/repository/repoimpl"
)

func DanhMucHeThongRoute(api core.DbData) {

	g1 := api.Echo.Group("/chucdanh")
	g2 := api.Echo.Group("/chucvu")

	middlewares.SetJwtMiddlewares(g1)
	middlewares.SetJwtMiddlewares(g2)

	handlerHeThong := handler.DanhMucHeThongHandler{
		DanhMucHeThongRepo: repo.NewDanhMucHeThongRePo(api.Db, api.Db),
	}

	g1.GET("", handlerHeThong.GetAllChucDanh)
	g2.GET("", handlerHeThong.GetAllChucVu)

}
