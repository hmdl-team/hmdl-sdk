package group

import (
	"hmdl-user-service/db/core"
	. "hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	"hmdl-user-service/repository/repoimpl"
)

func DM_DuAnRoute(api core.DbData) {

	handler := DM_DuAnHandler{
		Repo: repoimpl.NewDM_DuAnRepo(api.DbSql01),
	}

	g := api.Echo.Group("/dm-du-an")
	middlewares.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
