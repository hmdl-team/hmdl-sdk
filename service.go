package sdk

import (
	"fmt"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

type Handler func(svc Service)
type CommandHandler func(svc Service, args []string)

type Service interface {
	Setup() error
	Stop()
	Command(command string, handler CommandHandler) Service
	Run()
	RestHandler(handler Handler) Service
	NatHandler(handler Handler) Service

	// Get echo server
	Echo() *echo.Echo

	// Get Nat client
	Nat() *nats.Conn

	// Get Nat Json client
	NatJson() *nats.EncodedConn
}

type service struct {
	server      *echo.Echo
	rootCmd     *cobra.Command
	restHandler Handler
	natHandler  Handler
}

// Get echo server
func (s *service) Echo() *echo.Echo {
	return s.server
}

func (s *service) Setup() error {
	if err := godotenv.Load(); err != nil {
		logrus.Warning("Không tìm thấy file .env, bạn có thể bỏ qua nếu đã thêm biến môi trường bằng tay")
	}

	if err := sdk.ConnectDb(); err != nil {
		return err
	}

	if err := sdk.ConnectNat(); err != nil {
		return err
	}

	if err := sdk.ConnectElastic(); err != nil {
		return err
	}

	return nil
}

func (s *service) Stop() {
	_ = sdk.Db.Close()
	sdk.Nat.Close()
}

func (s *service) Command(command string, handler CommandHandler) Service {
	s.rootCmd.AddCommand(&cobra.Command{
		Use: command,
		Run: func(cmd *cobra.Command, args []string) {
			handler(s, args)
		},
	})
	return s
}

func (s *service) Run() {
	s.rootCmd.AddCommand(&cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			if err := s.Setup(); err != nil {
				logrus.Fatalln(err)
				return
			}
			defer s.Stop()

			s.server.HideBanner = true

			if os.Getenv("ENVIRONMENT") == "dev" {
				s.server.Debug = true
			}

			s.server.Use(middleware.Logger())
			s.server.Use(middleware.Gzip())
			s.server.Use(middleware.RemoveTrailingSlash())
			s.server.Use(middleware.Recover())

			// Đăng ký HandlerContext
			s.server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					cc := &sdk.HandlerContext{Context: c}
					return next(cc)
				}
			})

			if s.restHandler != nil {
				s.restHandler(s)
			}

			if s.natHandler != nil {
				if s.Nat() == nil || s.NatJson() == nil {
					logrus.Fatalln("Chưa kết nối đến NAT server, nên không thể chạy NatHandler")
				}
				s.natHandler(s)
			}

			port := os.Getenv("PORT")
			if port == "" {
				port = "5000"
			}
			logrus.Fatalln(s.server.Start(fmt.Sprintf(":%s", port)))
		},
	})

	err := s.rootCmd.Execute()
	if err != nil {
		logrus.Panic(err)
	}
}

func (s *service) RestHandler(handler Handler) Service {
	s.restHandler = handler
	return s
}

func (s *service) NatHandler(handler Handler) Service {
	s.natHandler = handler
	return s
}

func (s *service) Nat() *nats.Conn {
	return sdk.Nat
}

func (s *service) NatJson() *nats.EncodedConn {
	return sdk.NatJson
}

func NewService() *service {
	return &service{
		server:  echo.New(),
		rootCmd: &cobra.Command{},
	}
}