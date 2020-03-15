package group

import (
	"hmdl-user-service/db/core"
	handler2 "hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	"hmdl-user-service/repository/repoimpl"
)

func DanhSachNhanVienQuanLyRoute(api core.DbData) {

	handler := handler2.DanhSachNhanVienQuanLyHandler{
		Repo: repoimpl.NewDanhSachNhanVienQuanLyRepo(api.Db),
	}

	g := api.Echo.Group("/danhsachnhanvienquanly")
	middlewares.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
