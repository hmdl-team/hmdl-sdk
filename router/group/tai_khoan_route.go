package group

import (
	"hmdl-user-service/core"
	"hmdl-user-service/handler"
	"github.com/congnguyendl/hmdl-sdk/sdk"

	repo "hmdl-user-service/repository/repoimpl"
)

func TaiKhoanRoute(api core.DbData) {

	g := api.Echo.Group("/taikhoan")

	sdk.SetJwtMiddlewares(g)

	handlerTaiKhoan := handler.TaiKhoanHandler{
		TaiKhoanRepo: repo.NewDmTaiKhoanRepo(api.DbSql01),
	}

	// Tài khoản JWT
	g.GET("", handlerTaiKhoan.GetAllTaiKhoan)
	g.POST("", handlerTaiKhoan.InsertTaiKhoan)
	g.PUT("", handlerTaiKhoan.UpdateTaiKhoan)
	g.GET("/id/:id", handlerTaiKhoan.GetTaiKhoanById)
	g.DELETE("/id/:id", handlerTaiKhoan.DeleteTaiKhoanById)

	//public
	g.GET("/me", handlerTaiKhoan.GetNhanVienByToken)
	api.Echo.POST("/refresh-token", handlerTaiKhoan.GetRefreshToken)
	api.Echo.POST("/refresh-token-new", handlerTaiKhoan.GetRefreshToken2)
	api.Echo.POST("/login", handlerTaiKhoan.LoginAcount)

}
