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

type DmPhongbanhandler struct {
	Repo repository.DM_PhongBanRepo
}

func (u *DmPhongbanhandler) Insert(c echo.Context) error {
	data := new(DM_PhongBan)

	if err := c.Bind(data); err != nil {
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

func (u *DmPhongbanhandler) Update(c echo.Context) (err error) {

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

	data := new(DM_PhongBan)

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
		Message:    "Cập nhật thành công",
		Data:       data,
	})
}

func (u *DmPhongbanhandler) Delete(c echo.Context) (err error) {
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

func (u *DmPhongbanhandler) GetById(c echo.Context) (err error) {

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

func (u *DmPhongbanhandler) GetAll(c echo.Context) (err error) {

	data, err := u.Repo.GetAll(c)
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return  sdk.ResponseData(c, data)

}

func (u *DmPhongbanhandler) GetPhongBanComBobox(c echo.Context) (err error) {

	data, err := u.Repo.GetPhongBanComBobox(c)
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return  sdk.ResponseData(c, data)

}

func (u *DmPhongbanhandler) GetCayPhongBan(c echo.Context) (err error) {
	cc :=  sdk.GetHandlerContext(c)
	data, err := u.Repo.GetCayPhongBan(c)
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return cc.Ok(data)

}

func (u *DmPhongbanhandler) GetCayPhongBanTheoTaiKhoan(c echo.Context) (err error) {
	cc :=  sdk.GetHandlerContext(c)

	user := cc.GetUid()

	data, err := u.Repo.GetCayPhongBanByUserId(c.Request().Context(), user)
	if err != nil {
		return  sdk.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return cc.Ok(data)

}
