package repository

import (
	"github.com/labstack/echo/v4"
	. "hmdl-user-service/models/data_user"
)

type DM_ReportRepo interface {
	GetAll(ctx echo.Context) ([]DM_Report, error)
	GetById(ctx echo.Context, id int) (*DM_Report, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item DM_Report) (*DM_Report, error)
	Update(ctx echo.Context, item DM_Report) error

	GetReportPhanQuyenId(phanQuyenId int) ([]DM_Report, error)
}
