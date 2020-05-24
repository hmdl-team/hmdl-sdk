package handler

import (

	"github.com/congnguyendl/hmdl-sdk/sdk"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
	"net/http"
	"strconv"
)

type DmChucVuHandler struct {
	Repo repository.DmChucVuRepo
}

func (u *DmChucVuHandler) Insert(c echo.Context) error {
	data := new(DmChucVu)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, sdk.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(data); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest,  sdk.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	item, err := u.Repo.Insert(c, *data)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusInternalServerError,  sdk.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK,  sdk.Response{
		StatusCode: http.StatusOK,
		Message:    "Thêm mới thành công",
		Data:       item,
	})
}

func (u *DmChucVuHandler) Update(c echo.Context) (err error) {

	parentId := c.Param("id")

	if len(parentId) == 0 {
		// Bắt lỗi trả về client
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valParentId, err := strconv.ParseInt(parentId, 0, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest,  sdk.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	data := new(DmChucVu)

	data, _ = u.Repo.GetById(c, int(valParentId))

	if err = c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest,  sdk.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(data); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest,  sdk.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = u.Repo.Update(c, *data)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusInternalServerError,  sdk.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK,  sdk.Response{
		StatusCode: http.StatusOK,
		Message:    "Thêm mới thành công",
		Data:       data,
	})
}

func (u *DmChucVuHandler) Delete(c echo.Context) (err error) {
	phongKhamId := c.Param("id")

	if len(phongKhamId) == 0 {
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valConvert, err := strconv.ParseInt(phongKhamId, 0, 64)
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	data, err := u.Repo.GetById(c, int(valConvert))
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	if data == nil {
		return  sdk.ResponseWithCode(c, http.StatusNotFound, "Không tìm thấy dữ liệu")
	}

	err = u.Repo.Delete(c, int(valConvert))
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return  sdk.ResponseWithCode(c, 200, "Xóa thành công")
}

func (u *DmChucVuHandler) GetById(c echo.Context) (err error) {

	id := c.Param("id")

	if len(id) == 0 {
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valConvert, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	data, err := u.Repo.GetById(c, int(valConvert))
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}
	if data == nil {
		return  sdk.ResponseWithCode(c, http.StatusNotFound, "Không tìm thấy dữ liệu")
	}
	return  sdk.ResponseData(c, data)
}

func (u *DmChucVuHandler) GetAll(c echo.Context) (err error) {

	data, err := u.Repo.GetAll(c)
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return  sdk.ResponseData(c, data)

}
