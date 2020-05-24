package server

import (
	"context"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"google.golang.org/grpc"
	"hmdl-user-service/core"
	"hmdl-user-service/services"

	"hmdl-user-service/pb"
	"hmdl-user-service/repository/repoimpl"

	"log"
	"net"
	"os"

	"sync"
)

type Greeter struct {
	wg sync.WaitGroup
	db *core.DbData
}

func New(db *core.DbData) *Greeter {
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
	 sdk.RegisterServiceWithConsul("hmdl-user-service-grpc", lis.Addr().(*net.TCPAddr).Port, consulAddress)

	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, &services.UserService{
		RepoNhanVien: repoimpl.NewNhanVienRepo(g.db),
		RepoThemSoHeThong: repoimpl.NewDmThamSoHeThongRepo(g.db),
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
