package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	api "github.com/bharat-rajani/grpc-products-demo/api"
	pb "github.com/bharat-rajani/grpc-products-demo/gen/proto/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	// grpcPort = os.Getenv("GRPC_PORT")
	grpcPort = "8080"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		panic(err)
	}

	errs := make(chan error, 1)

	go func() {
		// create grpc server
		grpcServer := grpc.NewServer()

		// create product server struct
		var vendorServices = map[string][]string{
			"google": {"compute", "storage"},
			"aws":    {"compute", "storage"},
			"oracle": {"compute", "storage"},
		}

		productServer := api.NewProductServer(vendorServices)

		pb.RegisterProductServiceServer(grpcServer, productServer)
		reflection.Register(grpcServer)

		log.Printf("Starting grpc server on GRPC_PORT=[%s]", grpcPort)
		err := grpcServer.Serve(lis)
		errs <- err
	}()

	// Catch shutdown
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT)
		s := <-sig
		errs <- fmt.Errorf("caught signal %v", s)
	}()

	log.Fatal(<-errs)
}
