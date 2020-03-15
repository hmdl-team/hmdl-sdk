package handler

import (
	"github.com/labstack/echo/v4"
	"hmdl-user-service/helper/lib"
	"hmdl-user-service/repository"
)

type DanhMucHeThongHandler struct {
	DanhMucHeThongRepo repository.DanhMucHeThongRepository
}

func (u *DanhMucHeThongHandler) GetAllChucDanh(c echo.Context) error {

	//dbsql := drive.SqlServer.Db
	//db := drive.Postgres.Db

	//data, err := repoimpl.NewDanhMucHeThongRePo(dbsql,db).GetAllChucVu()
	data, err := u.DanhMucHeThongRepo.GetAllChucDanh()

	if err == nil {
		return c.JSON(200, lib.Response{
			Type:    "data",
			Message: "Sussess",
			Count:   len(data),
			Data:    data,
		})
	}
	return c.JSON(500, lib.Response{
		Type:    "error",
		Message: "Lỗi thực thi",
		Data:    nil,
	})

}

func (u *DanhMucHeThongHandler) GetAllChucVu(c echo.Context) error {

	data, err := u.DanhMucHeThongRepo.GetAllChucVu()

	if err == nil {
		return c.JSON(200, lib.Response{
			Type:    "data",
			Message: "Sussess",
			Count:   len(data),
			Data:    data,
		})
	}
	return c.JSON(500, lib.Response{
		Type:    "error",
		Message: "Lỗi thực thi",
		Data:    nil,
	})

}

//func DongBoDanhMucHeThong() {
//
//	dbsql := drive.SqlServer.Db
//	dbpos := drive.Postgres.Db
//	//defer log.Println(err)
//	repoimpl.NewDanhMucHeThongRePo(dbsql, dbpos).DongBoChucdanh()
//	repoimpl.NewDanhMucHeThongRePo(dbsql, dbpos).DongBoChucVu()
//
//}
