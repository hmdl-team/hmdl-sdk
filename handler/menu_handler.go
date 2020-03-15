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

type MenuHandler struct {
	MenuRepo          repository.MenuRepository
	PhanQuyenMenuRepo repository.PhanQuyenMenuRepository
}

func (u *MenuHandler) GetMenuByPhanQuyenIdAnDuAnId(c echo.Context) error {
	phanQuyenId, err := helper.CheckIntPar(c.QueryParam("phanquyenid"))

	if err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}
	duAnId, err := helper.CheckIntPar(c.QueryParam("duanid"))

	if err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	data, err := u.MenuRepo.GetMenuByPhanQuyenId(phanQuyenId, duAnId)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, data)
}

func (u *MenuHandler) GetAllMenu(c echo.Context) error {

	data, err := u.MenuRepo.GetAll(c)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, data)
}

func (u *MenuHandler) GetAllMenuById(c echo.Context) error {
	parentId := c.Param("id")

	if len(parentId) == 0 {
		// Bắt lỗi trả về client
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valParentId, err := strconv.ParseInt(parentId, 0, 64)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	data, err := u.MenuRepo.GetById(c, int(valParentId))

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, data)

}

func (u *MenuHandler) UpdateMenu(c echo.Context) (err error) {
	menu := new(data_user.DM_MenuWeb)
	if err = c.Bind(menu); err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, err.Error())
	}

	err = u.MenuRepo.Update(c, *menu)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, menu)
}

func (u *MenuHandler) DeleteMenuById(c echo.Context) error {
	parentId := c.Param("id")

	if len(parentId) == 0 {
		// Bắt lỗi trả về client
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valParentId, err := strconv.ParseInt(parentId, 0, 64)
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	err = u.MenuRepo.Delete(c, int(valParentId))

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, "Delete thành công")

}

func (u *MenuHandler) InsertMenu(c echo.Context) (err error) {
	menu := new(data_user.DM_MenuWeb)
	if err = c.Bind(menu); err != nil {
		return
	}

	_, err = u.MenuRepo.Insert(c, *menu)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, menu)
}
