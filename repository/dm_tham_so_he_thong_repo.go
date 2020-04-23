
	package repository

	import (
		."hmdl-user-service/models/data_user"
		"github.com/labstack/echo/v4"
	)

	type DmThamSoHeThongRepo interface {
		GetAll(ctx echo.Context) ([]DmThamSoHeThong, error)
		GetById(ctx echo.Context, id int) (*DmThamSoHeThong, error)
		Delete(ctx echo.Context, id int) error
		Insert(ctx echo.Context, item DmThamSoHeThong) (*DmThamSoHeThong, error)
		Update(ctx echo.Context, item DmThamSoHeThong) error
	}
       