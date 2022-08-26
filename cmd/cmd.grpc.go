package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
	"urbox-viettel-go/configs"
)

func GRPCServe(ctx *cli.Context) error {
	listen, err := net.Listen("tcp", ":"+configs.GlobalConf.GRPCPort)
	if err != nil {
		log.Fatalln(err)
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 30 * time.Second,
		}),
	}

	serve := grpc.NewServer(opts...)
	reflection.Register(serve)
	fmt.Println("GRPC Running: " + configs.GlobalConf.GRPCPort)
	return serve.Serve(listen)
}
