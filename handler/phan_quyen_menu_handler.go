package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/helper"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"

	"net/http"
	"strconv"
)

type PhanQuyenMenuHandler struct {
	PhanQuyenMenuRepo repository.PhanQuyenMenuRepository
}

func (u *PhanQuyenMenuHandler) GetAllPhanQuyenMenu(c echo.Context) error {

	data, err := u.PhanQuyenMenuRepo.GetAllPhanQuyenMenu()

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, data)
}

func (u *PhanQuyenMenuHandler) GetAllPhanQuyenMenuById(c echo.Context) error {
	parentId := c.Param("id")
	if len(parentId) == 0 {
		// Bắt lỗi trả về client
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valParentId, err := strconv.ParseInt(parentId, 0, 64)
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	data, err := u.PhanQuyenMenuRepo.GetById(int(valParentId))

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, data)
}

func (u *PhanQuyenMenuHandler) DeletePhanQuyenMenuById(c echo.Context) error {
	parentId := c.Param("id")

	if len(parentId) == 0 {
		// Bắt lỗi trả về client
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valParentId, err := strconv.ParseInt(parentId, 0, 64)
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	err = u.PhanQuyenMenuRepo.Delete(int(valParentId))

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, "Delete thành công")

}

func (u *PhanQuyenMenuHandler) InsertPhanQuyenMenu(c echo.Context) (err error) {
	menu := new(data_user.DM_PhanQuyenMenu)
	if err = c.Bind(menu); err != nil {
		return
	}

	err = u.PhanQuyenMenuRepo.Insert(menu)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, menu)
}
