package router

import (
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hmdl-user-service/core"
	"hmdl-user-service/handler/impl"
	"hmdl-user-service/queue"
	"hmdl-user-service/repository/repoimpl"
	"hmdl-user-service/router/group"
	"net/http"
)

type API struct {
	Echo *echo.Echo
	Db   *gorm.DB
}

func (api *API) NewRouter() {
	//cau hinh các Option
	structValidator := sdk.NewStructValidator()
	structValidator.RegisterValidate()
	api.Echo.Validator = structValidator

	//cau hinh các Option
	api.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
			http.MethodConnect,
			http.MethodHead,
			http.MethodTrace,
		},
	}))

	api.Echo.GET("/healthcheck", impl.HealthCheck)

	// Đăng ký HandlerContext
	api.Echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &sdk.HandlerContext{Context: c}
			return next(cc)
		}
	})

	db := core.DbData{
		Echo:    api.Echo,
		DbSql01: api.Db,
	}


	subUser := queue.UserSubscribeService{
		RepoNhanVien:      repoimpl.NewNhanVienRepo(&db),
		RepoThemSoHeThong: repoimpl.NewDmThamSoHeThongRepo(&db),
	}

	subUser.Subscribes()

	group.MenuRoute(db)
	group.NhanVienRoute(&db)
	group.PhanQuyenRoute(db)
	group.PhanQuyenMenuRoute(db)
	group.TaiKhoanRoute(db)
	group.DanhSachNhanVienQuanLyRoute(db)
	group.DM_DuAnRoute(db)
	group.DM_PhanQuyen_ReportRoute(db)
	group.DM_ReportRoute(db)
	group.DM_PhongBanRoute(db)
	group.DmThamSoHeThongRoute(&db)
	group.DmChucDanhRoute(&db)
	group.DmChucVuRoute(&db)

	//g := server.New(&db)
	//g.Start()
	//g.WaitStop()
}
