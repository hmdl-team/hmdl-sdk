package group

import (
	"hmdl-user-service/core"
	. "hmdl-user-service/handler"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"hmdl-user-service/repository/repoimpl"
)

func DmThamSoHeThongRoute(api *core.DbData) {

	handler := DmThamSoHeThongHandler{
		Repo: repoimpl.NewDmThamSoHeThongRepo(api),
	}

	g := api.Echo.Group("/tham-so-he-thong")
	sdk.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
