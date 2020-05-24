package group

import (
	"hmdl-user-service/core"
	. "hmdl-user-service/handler"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"hmdl-user-service/repository/repoimpl"
)

func DM_ReportRoute(api core.DbData) {

	handler := DM_ReportHandler{
		Repo: repoimpl.NewDM_ReportRepo(api.DbSql01),
	}

	g := api.Echo.Group("/dm-report")
	sdk.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.GET("/phan-quyen/", handler.GetBaoCaoByPhanQuyenId)
	g.GET("/he-thong-bao-cao/phan-quyen", handler.GetReportTrangBaoCaoByPhanQuyenId)
	g.POST("", handler.Insert)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
