package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"hmdl-user-service/models/response"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"hmdl-user-service/auth"
	"hmdl-user-service/helper"
	"hmdl-user-service/helper/lib"

	"hmdl-user-service/models/data_user"
	"hmdl-user-service/models/request"
	"hmdl-user-service/repository"

	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TaiKhoanHandler struct {
	TaiKhoanRepo repository.TaiKhoanRepository
}

// DanhMucDuoc godoc
// @Summary Login user
// @Description Get danh mục dược cho combobox
// @Tags tai-khoan
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body request.ReqSignIn true "user"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /login [post]
func (u *TaiKhoanHandler) LoginAcount(c echo.Context) (err error) {

	req := request.ReqSignIn{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(req); err != nil {
		fmt.Println(req)
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := u.TaiKhoanRepo.Login(c, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	mahoaPass := lib.MaHoa(req.Password)

	// check pass
	isTheSame := mahoaPass == user.MatKhauWeb
	if !isTheSame {
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Đăng nhập thất bại",
			Data:       nil,
		})
	}

	//	token, time, err := domain.GenToken(user)
	token, _, err := auth.GenTokenWithTime(user, 3)
	refeshToken, _, err := auth.GenTokenWithTime(user, 4)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	userRes := response.ResTaiKhoanToken{
		Token:        token,
		RefreshToken: refeshToken,
	}

	//user.Token = token
	//user.TokenTime = time

	return c.JSON(http.StatusOK, helper.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng nhập thành công",
		Data:       userRes,
	})
}

func (u *TaiKhoanHandler) GetAllTaiKhoan(c echo.Context) error {

	taikhoan, err := u.TaiKhoanRepo.GetAll(c)

	if err != nil {
		return c.JSON(500, lib.Response{
			Type:    "error",
			Message: "Lỗi thực thi",
			Data:    nil,
		})
	}

	if taikhoan == nil {
		return c.JSON(500, lib.Response{
			Type:    "error",
			Message: "Lỗi thực thi",
			Data:    nil,
		})
	}

	return c.JSON(200, lib.Response{
		Type:    "data",
		Message: "Sussess",
		Count:   len(taikhoan),
		Data:    taikhoan,
	})

}

// DanhMucDuoc godoc
// @Summary domain with token
// @Description Get danh mục dược cho combobox
// @Tags tai-khoan
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body request.ReqSignIn true "user"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /login [post]
func (u *TaiKhoanHandler) GetNhanVienByToken(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtClaims)
	userId := claims.UserId

	id, err := strconv.ParseInt(userId, 0, 64)

	if err != nil {
		return helper.ResponseWithCode(c, 401, "Unauthorized ")
	}

	taiKhoan, err := u.TaiKhoanRepo.GetById(c, int(id))

	if err != nil {
		return helper.ResponseWithCode(c, 404, "Không tìm thấy ")
	}

	resTaiKhoan := response.ResTaiKhoan{
		DM_NhanVienId:  taiKhoan.DM_NhanVienId,
		NhanVienSuDung: taiKhoan.NhanVienSuDung,
		PhanQuyenId:    taiKhoan.PhanQuyenId,
		DonVi:          taiKhoan.DonVi,
		DM_PhanQuyenID: taiKhoan.DM_PhanQuyenID,
		DM_PhanQuyen:   taiKhoan.DM_PhanQuyen,
	}
	return helper.ResponseData(c, resTaiKhoan)
}

// DanhMucDuoc godoc
// @Summary domain with token
// @Description Get danh mục dược cho combobox
// @Tags tai-khoan
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body request.ReqSignIn true "user"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /login [post]
func (u *TaiKhoanHandler) GetRefreshToken(c echo.Context) error {

	req := request.RefreshTokenRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(req); err != nil {
		fmt.Println(req)
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	claims := auth.DecodeToken(req.Token)

	if claims == nil {
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Lỗi chứng thực!",
			Data:       nil,
		})
	}

	userId := claims.UserId

	id, err := strconv.ParseInt(userId, 0, 64)

	if err != nil {
		return helper.ResponseWithCode(c, 401, "Unauthorized ")
	}

	taiKhoan, err := u.TaiKhoanRepo.GetById(c, int(id))

	if err != nil {
		return helper.ResponseWithCode(c, 404, "Không tìm thấy ")
	}

	token, _, err := auth.GenTokenWithTime(*taiKhoan, 3)
	refeshToken, _, err := auth.GenTokenWithTime(*taiKhoan, 4)

	resTaiKhoan := response.ResTaiKhoanToken{
		Token:        token,
		RefreshToken: refeshToken,
	}
	return helper.ResponseData(c, resTaiKhoan)
}

func (u *TaiKhoanHandler) GetRefreshToken2(c echo.Context) error {

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*auth.JwtClaims)

	if claims == nil {
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Lỗi chứng thực!",
			Data:       nil,
		})
	}

	userId := claims.UserId

	id, err := strconv.ParseInt(userId, 0, 64)

	if err != nil {
		return helper.ResponseWithCode(c, 401, "Unauthorized ")
	}

	taiKhoan, err := u.TaiKhoanRepo.GetById(c, int(id))

	if err != nil {
		return helper.ResponseWithCode(c, 404, "Không tìm thấy ")
	}

	token, _, err := auth.GenTokenWithTime(*taiKhoan, 3)
	refeshToken, _, err := auth.GenTokenWithTime(*taiKhoan, 4)

	resTaiKhoan := response.ResTaiKhoanToken{
		Token:        token,
		RefreshToken: refeshToken,
	}
	return helper.ResponseData(c, resTaiKhoan)
}

func (u *TaiKhoanHandler) UpdateTaiKhoan(c echo.Context) (err error) {
	data := data_user.DM_TaiKhoan{}

	if err = c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(data); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = u.TaiKhoanRepo.Update(c, data)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusInternalServerError, helper.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, helper.Response{
		StatusCode: http.StatusOK,
		Message:    "Cập nhật thành công",
		Data:       data,
	})
}

func (u *TaiKhoanHandler) DeleteTaiKhoanById(c echo.Context) error {
	phongKhamId := c.Param("id")

	if len(phongKhamId) == 0 {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	valConvert, err := strconv.ParseInt(phongKhamId, 0, 64)
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusBadRequest, "Dữ liệu không chính xác")
	}

	data, err := u.TaiKhoanRepo.GetById(c, int(valConvert))

	if err != nil {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	if data == nil {
		return helper.ResponseWithCode(c, http.StatusNotFound, "Không tìm thấy dữ liệu")
	}

	err = u.TaiKhoanRepo.Delete(c, int(valConvert))
	if err != nil {
		return helper.ResponseWithCode(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseWithCode(c, 200, "Xóa thành công")

}

func (u *TaiKhoanHandler) InsertTaiKhoan(c echo.Context) (err error) {

	data := new(data_user.DM_TaiKhoan)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := c.Validate(data); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, helper.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	item, err := u.TaiKhoanRepo.Insert(c, *data)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusInternalServerError, helper.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, helper.Response{
		StatusCode: http.StatusOK,
		Message:    "Thêm mới thành công",
		Data:       item,
	})
}

func (u *TaiKhoanHandler) GetTaiKhoanById(c echo.Context) error {
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
		return c.JSON(http.StatusBadRequest, lib.Response{
			Type:    "error",
			Message: "Dữ liệu không chính xác",
			Data:    nil,
		})
	}

	data, err := u.TaiKhoanRepo.GetById(c, int(valParentId))

	if err == nil {
		return c.JSON(200, lib.Response{
			Type:    "data",
			Message: "Sussess",
			Count:   1,
			Data:    data,
		})
	}
	return c.JSON(500, lib.Response{
		Type:    "error",
		Message: "Lỗi thực thi",
		Data:    nil,
	})

}
