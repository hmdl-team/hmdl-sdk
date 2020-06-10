package sdk

import (
	"errors"
	"fmt"
	"github.com/kevholditch/gokong"
	"os"
	"strings"
)

type KongServer struct {
	ServerAddress string
	ServiceName   string
	ServicePath   string
	ServicePort   string
}

func (u KongServer) Check() error {

	if len(u.ServerAddress) == 0 {
		return errors.New("KONG_ADDRESS chưa được khai báo, module Kong sẽ không được sử dụng.")
	}
	if len(u.ServiceName) == 0 {
		return errors.New("KONG_SERVICE_NAME chưa được khai báo, module Kong sẽ không được sử dụng.")
	}

	if len(u.ServicePort) == 0 {
		return errors.New("PORT chưa được khai báo, module Kong sẽ không được sử dụng.")
	}
	if len(u.ServicePath) == 0 {
		return errors.New("KONG_SERVICE_PATH chưa được khai báo, module Kong sẽ không được sử dụng.")
	}
	return nil
}

func (u KongServer) NewKongServerFromEnv() KongServer {
	u.ServerAddress = os.Getenv("KONG_ADDRESS")
	u.ServiceName = os.Getenv("KONG_SERVICE_NAME")
	u.ServicePort = os.Getenv("PORT")
	u.ServicePath = os.Getenv("KONG_SERVICE_PATH")

	return u
}

func (u KongServer) RegisterKong() error {

	if err := u.Check(); err != nil {
		return err
	}

	config := gokong.Config{HostAddress: u.ServerAddress}

	client := gokong.NewClient(&config)
	status, err := client.Status().Get()

	if err != nil {
		fmt.Println(status)
	}

	nameStream := strings.ToLower(u.ServiceName + "-upstream")
	// Tìm upstream
	upstream, err := client.Upstreams().GetByName(nameStream)

	//Nêu chưa có tạo mới upstream
	if upstream == nil {
		upstreamRequest := &gokong.UpstreamRequest{
			Name:  nameStream,
			Slots: 1000,
		}

		upstream, err = client.Upstreams().Create(upstreamRequest)

		if err != nil {
			fmt.Println(status)
			return err
		}

	}

	if upstream != nil {

		targetOld, err := client.Targets().GetTargetsFromUpstreamId(upstream.Id)

		if err != nil {
			return err
		}

		for _, tarOld := range targetOld {

			err = client.Targets().DeleteFromUpstreamById(upstream.Id, *tarOld.Id)
			if err != nil {
				return err
			}

		}

		targetRequest := &gokong.TargetRequest{
			Target:  GetLocalIP().String() + ":"+u.ServicePort,
			Weight: 100,
		}
		_, err = client.Targets().CreateFromUpstreamId(upstream.Id, targetRequest)

		if err != nil {

			return err
		}

		targets, err := client.Targets().GetTargetsFromUpstreamId(upstream.Id)

		for _, item := range targets {
			err := client.Targets().SetTargetFromUpstreamByIdAsHealthy(upstream.Id, *item.Id)
			if err != nil {

				return err
			}
		}

		if err != nil {
			fmt.Println(status)
			return err
		}

		// Đăng ký service
		serviceRequest := &gokong.ServiceRequest{
			Name:     gokong.String(strings.ToLower(u.ServiceName)),
			Protocol: gokong.String("http"),
			Host:     gokong.String(strings.ToLower(upstream.Name)),
		}

		// Kiểm tra đã có service chưa
		createdService, err := client.Services().GetServiceByName(*serviceRequest.Name)

		if err != nil {
			fmt.Println(status)
			return err
		}

		// Nếu chưa thì tạo service
		if createdService == nil {
			createdService, err = client.Services().Create(serviceRequest)
			if err != nil {
				fmt.Println(status)
				return err
			}
		}

		// Nếu đã tồn tại thì update
		if createdService != nil && createdService.Id != nil {
			createdService, err = client.Services().UpdateServiceById(*createdService.Id, serviceRequest)
			if err != nil {
				fmt.Println(status)
				return err
			}
		}

		if createdService != nil {
			routeRequest := &gokong.RouteRequest{
				Name:          gokong.String(strings.ToLower(*createdService.Name + "-ROUTER")),
				Protocols:     gokong.StringSlice([]string{"http", "https"}),
				Methods:       gokong.StringSlice([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}),
				Paths:         gokong.StringSlice([]string{strings.ToLower(u.ServicePath)}),
				RegexPriority: gokong.Int(0),
				StripPath:     gokong.Bool(true),
				PreserveHost:  gokong.Bool(true),
				Service:       gokong.ToId(*createdService.Id),
			}

			createdRoute, err := client.Routes().GetByName(*routeRequest.Name)

			if err != nil {
				fmt.Println(status)
				return err
			}

			// Nếu chưa thì tạo router
			if createdRoute == nil {
				_, err := client.Routes().Create(routeRequest)
				if err != nil {
					return err
				}
			} else {
				_, err = client.Routes().UpdateById(*createdRoute.Id, routeRequest)

				if err != nil {

					return err
				}
			}
		}
	}
	return nil
}