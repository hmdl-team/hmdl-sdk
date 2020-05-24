package sdk

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lyquocnam/go-sdk/sdkcm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

type Service interface {
	Setup() error
	Stop()
	Command(cmds ...*cobra.Command) Service
	Run()
	Handler(handler ServiceHandler) Service
	Subscriber(handler ServiceHandler) Service
	Server() *echo.Echo
}
type service struct {
	server  *echo.Echo
	rootCmd *cobra.Command
	routerHandler ServiceHandler
	subscribeHandler ServiceHandler
}
func (s *service) Server() *echo.Echo {
	return s.server
}
func (s *service) Setup() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := sdkcm.ConnectDb(); err != nil {
		return err
	}

	if err := sdkcm.ConnectNat(); err != nil {
		panic(err)
	}

	return nil
}
func (s *service) Stop() {
	_ = sdkcm.Db.Close()
	sdkcm.Nat.Close()
}
func (s *service) Command(commands ...*cobra.Command) Service {
	s.rootCmd.AddCommand(commands...)
	return s
}
func (s *service) Run() {
	s.rootCmd.AddCommand(&cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			if err := s.Setup(); err != nil {
				logrus.Fatalln(err)
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
					cc := &sdkcm.HandlerContext{Context: c}
					return next(cc)
				}
			})

			if s.subscribeHandler != nil {
				s.subscribeHandler(s)
			}
			if s.routerHandler != nil {
				s.routerHandler(s)
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
func (s *service) Handler(handler ServiceHandler) Service {
	s.routerHandler = handler
	return s
}
func (s *service) Subscriber(handler ServiceHandler) Service  {
	s.subscribeHandler = handler
	return s
}
func NewService() *service {
	return &service{
		server:  echo.New(),
		rootCmd: &cobra.Command{},
	}
}
type ServiceHandler func(svc Service)