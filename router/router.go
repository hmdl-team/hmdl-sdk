package router

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"hmdl-user-service/Services"
	"hmdl-user-service/db/core"
	"hmdl-user-service/handler/impl"
	"hmdl-user-service/helper"
	"hmdl-user-service/pb"
	"hmdl-user-service/repository/repoimpl"
	"hmdl-user-service/router/group"
	"net"
	"net/http"
	"os"
)

type API struct {
	Echo *echo.Echo
	Db   *gorm.DB
}

func (api API) NewRouter() {
	//Đăng ký service với consul
	//consulAddress := os.Getenv("CONSUL_ADDRESS")
	//helper.RegisterServiceWithConsul("HMDL-USER-SERVICE", 7001, consulAddress)

	kongAddress := os.Getenv("KONG_ADDRESS")
	kong := helper.KongServer{
		ServerKong:  kongAddress,
		NameService: "HMDL-USER-SERVICE",
		PathService: "/user-service",
		UrlService:  fmt.Sprintf("http://%s:%s", helper.GetLocalIP(), "7001"),
		IpService:   helper.GetLocalIP().String() + ":7001",
	}

	consulAddress := os.Getenv("CONSUL_ADDRESS")
	helper.RegisterServiceWithConsul("hmdl-user-service", 7001, consulAddress)

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println(err)
	}

	helper.RegisterServiceWithConsul("hmdl-user-service-grpc", listener.Addr().(*net.TCPAddr).Port, consulAddress)

	srv := grpc.NewServer()
	pb.RegisterNhanVienServiceServer(srv, &Services.NhanVienServicePro{
		RepoNhanVien: repoimpl.NewNhanVienRepo(api.Db),
	})
	reflection.Register(srv)

	go func() {
		if e := srv.Serve(listener); e != nil {
			panic(e)
		}
	}()

	err = kong.RegisterKong()

	if err != nil {
		fmt.Println(err)
	}

	// show log api request

	//api.Echo.Use(middleware.Logger())
	//api.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: "method=${method}, uri=${uri}, status=${status}, remote_ip=${remote_ip}\n",
	//}))
	//api.Echo.Use(middleware.Recover())

	//cau hinh các Option
	structValidator := helper.NewStructValidator()
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

	db := core.DbData{
		Echo: api.Echo,
		Db:   api.Db,
	}

	group.MenuRoute(db)
	group.NhanVienRoute(db)
	group.PhanQuyenRoute(db)
	group.PhanQuyenMenuRoute(db)
	group.TaiKhoanRoute(db)
	group.DanhSachNhanVienQuanLyRoute(db)
	group.DM_DuAnRoute(db)
	group.DM_PhanQuyen_ReportRoute(db)
	group.DM_ReportRoute(db)
	group.DM_PhongBanRoute(db)
}
