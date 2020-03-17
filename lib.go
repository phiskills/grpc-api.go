package grpc

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type api struct {
	Port     *int
	Server   *grpc.Server
	listener net.Listener
}

func New() (a *api) {
	var err error
	a = &api{}
	a.Port = flag.Int("port", 50051, "the server port")
	flag.Parse()
	a.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", *a.Port))
	if err != nil {
		log.Fatalf("failed to listen:%v", err)	
	}
	a.Server = grpc.NewServer()
	return
}

func (a *api) Start() {
	if err := a.Server.Serve(a.listener); err != nil {
		log.Fatalf("failed to serve grpc:%v", err)
	}
}
