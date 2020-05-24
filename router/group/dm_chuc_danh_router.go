package group

import (
	"hmdl-user-service/core"
	handler2 "hmdl-user-service/handler"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"hmdl-user-service/repository/repoimpl"
)

func DmChucDanhRoute(api *core.DbData) {

	handler := handler2.DmChucDanhHandler{
		Repo: repoimpl.NewDmChucDanhRepo(api),
	}

	g := api.Echo.Group("/chuc-danh")
	sdk.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
