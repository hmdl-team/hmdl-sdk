package group

import (
	"hmdl-user-service/db/core"
	. "hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	"hmdl-user-service/repository/repoimpl"
)

func DM_PhongBanRoute(api core.DbData) {

	handler := DM_PhongBanHandler{
		Repo: repoimpl.NewDM_PhongBanRepo(api.Db),
	}

	g := api.Echo.Group("/phong-ban")
	middlewares.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
