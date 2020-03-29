package repository

import (
	"github.com/labstack/echo/v4"
	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/models/request"
)

type DM_PhanQuyen_ReportRepo interface {
	GetAll(ctx echo.Context) ([]DM_PhanQuyen_Report, error)
	GetById(ctx echo.Context, id int) (*DM_PhanQuyen_Report, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item DM_PhanQuyen_Report) (*DM_PhanQuyen_Report, error)
	Update(ctx echo.Context, item DM_PhanQuyen_Report) error
	UpdatePhanQuyen(ctx echo.Context, req request.PhanQuyenBaoCaoReq) error
}
