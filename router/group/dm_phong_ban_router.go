package group

import (
	"hmdl-user-service/core"
	. "hmdl-user-service/handler"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"hmdl-user-service/repository/repoimpl"
)

func DM_PhongBanRoute(api core.DbData) {

	handler := DmPhongbanhandler{
		Repo: repoimpl.NewDM_PhongBanRepo(api.DbSql01),
	}

	g := api.Echo.Group("/phong-ban")
	sdk.SetJwtMiddlewares(g)

	g.GET("/combobox", handler.GetPhongBanComBobox)
	g.GET("/cayphongban", handler.GetCayPhongBanTheoTaiKhoan)
	g.GET("", handler.GetAll)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
