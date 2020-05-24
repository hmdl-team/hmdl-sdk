package sdk

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net"
	"os"
)

func RegisterServiceWithConsul(serviceId string, port int, serverConsulIp string) {
	config := consulapi.DefaultConfig()

	config.Address = serverConsulIp
	config.Scheme = "http"

	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = serviceId   //replace with service id
	registration.Name = serviceId //replace with service name
	address := GetLocalIP()
	registration.Address = address.String()
	registration.Port = port
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, port)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "5s"

	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println(err)
	}
}

func LookupServiceWithConsul(serviceID string) (string, error) {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)
	if err != nil {
		return "", err
	}
	services, err := client.Agent().Services()
	if err != nil {
		return "", err
	}
	srvc := services[serviceID]
	address := srvc.Address
	port := srvc.Port
	return fmt.Sprintf("http://%s:%v", address, port), nil
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}

func GetUrlService(port string) string {
	return fmt.Sprintf("http://%s:%s", GetLocalIP(), port)
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
