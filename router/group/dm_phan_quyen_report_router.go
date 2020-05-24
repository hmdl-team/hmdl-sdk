package group

import (
	"hmdl-user-service/core"
	. "hmdl-user-service/handler"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"hmdl-user-service/repository/repoimpl"
)

func DM_PhanQuyen_ReportRoute(api core.DbData) {

	handler := DM_PhanQuyen_ReportHandler{
		Repo: repoimpl.NewDM_PhanQuyen_ReportRepo(api.DbSql01),
	}

	g := api.Echo.Group("/dm-phan-quyen-report")
	sdk.SetJwtMiddlewares(g)

	g.GET("", handler.GetAll)
	g.POST("", handler.Insert)
	g.POST("/update-phan-quyen", handler.UpdatePhanQuyenReport)
	g.PUT("/:id", handler.Update)
	g.DELETE("/id/:id", handler.Delete)
	g.GET("/id/:id", handler.GetById)
}
