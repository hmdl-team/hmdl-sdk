package group

import (
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"hmdl-user-service/core"
	handler2 "hmdl-user-service/handler"

	"hmdl-user-service/repository/repoimpl"
)

func DanhSachNhanVienQuanLyRoute(api core.DbData) {

	handler := handler2.DanhSachNhanVienQuanLyHandler{
		Repo: repoimpl.NewDanhSachNhanVienQuanLyRepo(api.DbSql01),
	}

	g := api.Echo.Group("/danhsachnhanvienquanly")
	sdk.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
