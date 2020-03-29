package repository

import (
	"github.com/labstack/echo/v4"
	. "hmdl-user-service/models/data_user"
)

type DM_DuAnRepo interface {
	GetAll(ctx echo.Context) ([]DM_DuAn, error)
	GetById(ctx echo.Context, id int) (*DM_DuAn, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item DM_DuAn) (*DM_DuAn, error)
	Update(ctx echo.Context, item DM_DuAn) error
}
