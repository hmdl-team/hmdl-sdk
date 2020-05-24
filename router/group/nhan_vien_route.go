package group

import (
	"hmdl-user-service/db/core"
	"hmdl-user-service/handler"
	"hmdl-user-service/middlewares"
	repo "hmdl-user-service/repository/repoimpl"
)

func NhanVienRoute(api *core.DbData) {

	nvg := api.Echo.Group("/nhan-vien")

	middlewares.SetJwtMiddlewares(nvg)

	handlerNhanVien := handler.NhanVienHandler{
		NhanVienRepo: repo.NewNhanVienRepo(api),
	}

	// Thông tin nhân viên
	nvg.GET("", handlerNhanVien.GetAllNhanVien)
	nvg.POST("", handlerNhanVien.InsertNhanVien)
	nvg.PUT("", handlerNhanVien.UpdateNhanVien)
	nvg.GET("/id/:id", handlerNhanVien.GetNhanVienById)
	api.Echo.GET("/bacsi/id/:id", handlerNhanVien.GetNhanVienById)
	nvg.DELETE("/id/:id", handlerNhanVien.DeleteNhanVienById)

	nvg.GET("/combobox", handlerNhanVien.GetAllNhanVienCombobox)
	nvg.GET("/username", handlerNhanVien.GetNhanVienByUserName)
	nvg.GET("/danhsachbacsi", handlerNhanVien.GetDanhSachBacSi)
	nvg.GET("/chucdanhid/:id", handlerNhanVien.GetNhanVienByChucDanhId)
	nvg.GET("/:id/phongbanid", handlerNhanVien.GetNhanVienByPhongBanId)
}
