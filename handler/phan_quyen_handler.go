package handler

import (
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"github.com/labstack/echo/v4"

	"hmdl-user-service/repository"
)

type PhanQuyenHandler struct {
	PhanQuyenRepo repository.DmPhanQuyenRepo
}

func (u *PhanQuyenHandler) GetAllPhanQuyen(c echo.Context) error {

	data, err := u.PhanQuyenRepo.GetAll(c)

	if err == nil {


		return c.JSON(200, sdk.NewAppError(err))
	}
	return c.JSON(500, sdk.Response{
		Message: "Sussess",
		Data:    data,
	})

}

