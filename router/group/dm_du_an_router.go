package group

import (
	"hmdl-user-service/core"
	. "hmdl-user-service/handler"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"hmdl-user-service/repository/repoimpl"
)

func DM_DuAnRoute(api core.DbData) {

	handler := DM_DuAnHandler{
		Repo: repoimpl.NewDM_DuAnRepo(api.DbSql01),
	}

	g := api.Echo.Group("/dm-du-an")
	sdk.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
