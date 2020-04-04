package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/helper"
	"hmdl-user-service/helper/lib"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
	"net/http"
	"strconv"
)

type NhanVienHandler struct {
	NhanVienRepo repository.NhanVienRepository
}

func (u *NhanVienHandler) GetNhanVienByUserName(c echo.Context) error {
	name := c.QueryParam("username")

	//user2 := c.Get("user").(*jwt.Token)

	nhanvien := u.NhanVienRepo.GetNhanVienByUserName(name)
	return helper.ResponseData(c, nhanvien)
}

func (u *NhanVienHandler) GetNhanVienByChucDanhId(c echo.Context) error {
	id := c.Param("id")

	valId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	nhanvien := u.NhanVienRepo.GetDanhSachNhanVienByChucDanhId(int(valId))

	return helper.ResponseData(c, nhanvien)
}

func (u *NhanVienHandler) GetDanhSachBacSi(c echo.Context) error {

	nhanvien := u.NhanVienRepo.GetDanhSachBacSi()

	return helper.ResponseData(c, nhanvien)
}

func (u *NhanVienHandler) GetNhanVienById(c echo.Context) error {
	parentId := c.Param("id")

	if len(parentId) == 0 {
		// Bắt lỗi trả về client

		return c.JSON(http.StatusBadRequest, lib.Response{
			Type:    "error",
			Message: "Dữ liệu không chính xác",
			Data:    nil,
		})
	}

	valParentId, err := strconv.ParseInt(parentId, 0, 64)
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	data := u.NhanVienRepo.GetNhanVienById(int(valParentId))

	return helper.ResponseData(c, data)

}

func (u *NhanVienHandler) DeleteNhanVienById(c echo.Context) error {
	parentId := c.Param("id")

	if len(parentId) == 0 {
		// Bắt lỗi trả về client
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valParentId, err := strconv.ParseInt(parentId, 0, 64)
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	err = u.NhanVienRepo.Delete(int(valParentId))

	if err != nil {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, "Delete thành công")

}

func (u *NhanVienHandler) InsertNhanVien(c echo.Context) (err error) {
	nhanvien := new(data_user.NhanVien)

	if err = c.Bind(nhanvien); err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, err.Error())
	}

	err = u.NhanVienRepo.Insert(nhanvien)
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, nhanvien)
}

func (u *NhanVienHandler) UpdateNhanVien(c echo.Context) (err error) {

	nhanvien := new(data_user.NhanVien)
	if err = c.Bind(nhanvien); err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, err.Error())
	}

	err = u.NhanVienRepo.Update(nhanvien)
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, nhanvien)
}

func (u *NhanVienHandler) GetAllNhanVien(c echo.Context) error {

	data, err := u.NhanVienRepo.GetAll()

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, data)
}
