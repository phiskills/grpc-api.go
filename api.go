package grpc

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	grpcPort   = 50051
	bufferSize = 10
)

type api struct {
	listener     net.Listener
	name         string
	port         int
	server       *grpc.Server
	healthServer *health.Server
	status       status
	change       chan status
}

func New(name string) (a *api) {
	a = &api{name: name, change: make(chan status, bufferSize)}
	a.port = grpcPort
	a.server = grpc.NewServer()
	a.healthServer = health.NewServer()
	grpc_health_v1.RegisterHealthServer(a.server, a.healthServer)
	return
}

func (a *api) UseParams() {
	a.parseEnv()
	a.parseFlag()
}

func (a *api) UsePort(port int) {
	a.port = port
}

func (a *api) Start() {
	var err error
	addr := fmt.Sprintf(":%d", a.port)
	a.listener, err = net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	a.Activate()
	if err := a.server.Serve(a.listener); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}

func (a *api) Stop() {
	a.server.GracefulStop()
	a.Deactivate()
}

func (a *api) Activate() {
	a.updateStatus(Serving)
}

func (a *api) Deactivate() {
	a.updateStatus(NotServing)
}

func (a *api) Server() *grpc.Server {
	return a.server
}

func (a *api) Name() string {
	return a.name
}

func (a *api) Port() int {
	return a.port
}

func (a *api) Status() status {
	return a.status
}

func (a *api) Change() status {
	return <-a.change
}

func (a *api) updateStatus(status status) {
	a.healthServer.SetServingStatus(a.name, status.Value())
	a.status = status
	a.change <- a.status
}

func (a *api) parseFlag() {
	port := *flag.Int("port", 0, "the server port")
	flag.Parse()
	if port != 0 {
		a.port = port
	}
}

func (a *api) parseEnv() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil && port != 0 {
		a.port = port
	}
}
