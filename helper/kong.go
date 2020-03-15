package helper

import (
	"fmt"
	"github.com/kevholditch/gokong"
	"strings"
)

type KongServer struct {
	ServerKong  string `json:"url"`
	UrlService  string `json:"url"`
	IpService   string `json:"url"`
	NameService string `json:"name"`
	PathService string `json:"path"`
}

func (u *KongServer) RegisterKong() error {

	config := gokong.Config{HostAddress: u.ServerKong}

	client := gokong.NewClient(&config)
	status, err := client.Status().Get()

	if err != nil {
		fmt.Println(status)
	}

	nameStream := strings.ToLower(u.NameService + "-upstream")
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
		targetRequest := &gokong.TargetRequest{
			Target: u.IpService,
			Weight: 100,
		}
		_, err := client.Targets().CreateFromUpstreamId(upstream.Id, targetRequest)

		if err != nil {
			fmt.Println(status)
			return err
		}

		targets, err := client.Targets().GetTargetsFromUpstreamId(upstream.Id)

		for _, item := range targets {
			err := client.Targets().SetTargetFromUpstreamByIdAsHealthy(upstream.Id, *item.Id)
			if err != nil {
				fmt.Println(status)
				return err
			}
		}

		if err != nil {
			fmt.Println(status)
			return err
		}

		// Đăng ký service

		serviceRequest := &gokong.ServiceRequest{
			Name:     gokong.String(strings.ToLower(u.NameService)),
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
				Paths:         gokong.StringSlice([]string{strings.ToLower(u.PathService)}),
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
					fmt.Println(status)
					return err
				}
			} else {
				_, err = client.Routes().UpdateById(*createdRoute.Id, routeRequest)

				if err != nil {
					fmt.Println(status)
					return err
				}
			}

		}

	}

	return nil
}
