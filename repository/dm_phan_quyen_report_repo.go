package repository

import (
	"github.com/labstack/echo/v4"
	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/models/request"
)

type DmPhanquyenReportrepo interface {
	GetAll(ctx echo.Context) ([]DmPhanquyenReport, error)
	GetById(ctx echo.Context, id int) (*DmPhanquyenReport, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item DmPhanquyenReport) (*DmPhanquyenReport, error)
	Update(ctx echo.Context, item DmPhanquyenReport) error
	UpdatePhanQuyen(ctx echo.Context, req request.PhanQuyenBaoCaoReq) error
}
