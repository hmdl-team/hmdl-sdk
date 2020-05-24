package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lyquocnam/go-sdk"
	"github.com/lyquocnam/go-sdk/sdkcm"
	"github.com/spf13/cobra"
	"net/http"
)

func main() {
	service := sdk.NewService().Handler(router)
	service.Command(hashCommand(), setupCommand(service))
	service.Run()
}

func router(service sdk.Service) {
	e := service.Server()
	e.GET("/", func(c echo.Context) error {
		cc := sdkcm.GetHandlerContext(c)

		return cc.JSON(http.StatusOK, "ok")
	})
}

func hashCommand() *cobra.Command {
	return &cobra.Command{
		Use: "hash",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := sdkcm.HashPassword(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("-----------------------------------------------------------------")
				fmt.Println(p)
				fmt.Println("-----------------------------------------------------------------")
			}
		},
	}
}

func setupCommand(service sdk.Service) *cobra.Command {
	return &cobra.Command{
		Use: "setup",
		Run: func( cmd *cobra.Command, args []string) {
			fmt.Println("setup...")
		},
	}
}
