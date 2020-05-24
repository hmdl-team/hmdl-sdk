package handler

import (
	"fmt"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

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
	phanQuyenId, err :=   sdk.CheckIntPar(c.QueryParam("phanquyenid"))

	if err != nil {
		fmt.Println(err)
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu phanquyenid không chính xác:"+ err.Error())
	}
	duAnId, err :=    sdk.CheckIntPar(c.QueryParam("duanid"))

	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu duanid không chính xác")
	}

	data, err := u.MenuRepo.GetMenuByPhanQuyenIdAndDuAnId(phanQuyenId, duAnId)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return  sdk.ResponseData(c, data)
}

func (u *MenuHandler) GetMenuByPhanQuyenId(c echo.Context) error {
	phanQuyenId, err :=    sdk.CheckIntPar(c.QueryParam("phanquyenid"))

	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	data, err := u.MenuRepo.GetMenuByPhanQuyenId(phanQuyenId)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return  sdk.ResponseData(c, data)
}

func (u *MenuHandler) GetAllMenu(c echo.Context) error {

	data, err := u.MenuRepo.GetAll(c)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return  sdk.ResponseData(c, data)
}

func (u *MenuHandler) GetAllMenuById(c echo.Context) error {
	parentId := c.Param("id")

	if len(parentId) == 0 {
		// Bắt lỗi trả về client
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valParentId, err := strconv.ParseInt(parentId, 0, 64)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	data, err := u.MenuRepo.GetById(c, int(valParentId))

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return  sdk.ResponseData(c, data)

}

func (u *MenuHandler) UpdateMenu(c echo.Context) (err error) {
	menu := new(data_user.DM_MenuWeb)
	if err = c.Bind(menu); err != nil {
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, err.Error())
	}

	err = u.MenuRepo.Update(c, *menu)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return  sdk.ResponseData(c, menu)
}

func (u *MenuHandler) DeleteMenuById(c echo.Context) error {
	parentId := c.Param("id")

	if len(parentId) == 0 {
		// Bắt lỗi trả về client
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valParentId, err := strconv.ParseInt(parentId, 0, 64)
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	err = u.MenuRepo.Delete(c, int(valParentId))

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return  sdk.ResponseData(c, "Delete thành công")

}

func (u *MenuHandler) InsertMenu(c echo.Context) (err error) {
	menu := new(data_user.DM_MenuWeb)
	if err = c.Bind(menu); err != nil {
		return
	}

	_, err = u.MenuRepo.Insert(c, *menu)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	return  sdk.ResponseData(c, menu)
}
