package group

import (
	"hmdl-user-service/db/core"
	. "hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	"hmdl-user-service/repository/repoimpl"
)

func DM_ReportRoute(api core.DbData) {

	handler := DM_ReportHandler{
		Repo: repoimpl.NewDM_ReportRepo(api.Db),
	}

	g := api.Echo.Group("/dm-report")
	middlewares.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.GET("/phan-quyen/", handler.GetBaoCaoByPhanQuyenId)
	g.GET("/he-thong-bao-cao/phan-quyen/", handler.GetReportTrangBaoCaoByPhanQuyenId)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
