package handler

import (
	"github.com/labstack/echo/v4"
	"hmdl-user-service/helper/lib"
	"hmdl-user-service/repository"
)

type PhanQuyenHandler struct {
	PhanQuyenRepo repository.DmPhanQuyenRepo
}

func (u *PhanQuyenHandler) GetAllPhanQuyen(c echo.Context) error {

	data, err := u.PhanQuyenRepo.GetAll(c)

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

