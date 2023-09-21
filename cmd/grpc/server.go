package main

import (
	"fmt"
	"net"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/logger"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	conf   *config.AppConfig
	server *grpc.Server
	lis    net.Listener
}

func NewGRPCServer(conf *config.AppConfig, server *grpc.Server, lis net.Listener) *GRPCServer {
	return &GRPCServer{
		conf:   conf,
		server: server,
		lis:    lis,
	}
}

func (s *GRPCServer) Run() {
	s.server.Serve(s.lis)
}

func NewNetListener(conf *config.AppConfig) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GetString("app.grpc.port")))
	if err != nil {
		logger.Console.Error(err)
		panic(err)
	}
	logger.Console.Info("listen at :", conf.GetString("app.grpc.port"))
	return lis
}
