package Server

import (
	"context"

	"google.golang.org/grpc"
	"hmdl-user-service/Services"
	"hmdl-user-service/helper"
	"hmdl-user-service/pb"
	"hmdl-user-service/repository/repoimpl"
	"hmdl-user-service/router"
	"log"
	"net"
	"os"

	"sync"
)

type Greeter struct {
	wg sync.WaitGroup
	db *router.API
}

func New(db *router.API) *Greeter {
	return &Greeter{
		db: db,
	}
}

func (g *Greeter) Start() {
	g.wg.Add(2)
	go func() {
		log.Fatal(g.startGRPC())
		g.wg.Done()
	}()

	go func() {
		log.Fatal(g.startREST())
		g.wg.Done()
	}()

}

func (g *Greeter) WaitStop() {
	g.wg.Wait()
}

func (g *Greeter) startGRPC() error {

	lis, err := net.Listen("tcp", ":0")

	if err != nil {
		return err
	}
	consulAddress := os.Getenv("CONSUL_ADDRESS")
	helper.RegisterServiceWithConsul("hmdl-user-service-grpc", lis.Addr().(*net.TCPAddr).Port, consulAddress)

	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, &Services.NhanVienServicePro{
		RepoNhanVien: repoimpl.NewNhanVienRepo(g.db.Db),
	})
	err = srv.Serve(lis)

	if err != nil {
		return err
	}
	return nil
}

func (g *Greeter) startREST() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return g.db.Echo.Start(":7001")
}
