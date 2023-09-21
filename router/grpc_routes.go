package router

import (
	helloWorldServer "github.com/adityarifqyfauzan/go-boilerplate/app/domain/pb/proto/helloworld"
	"github.com/adityarifqyfauzan/go-boilerplate/app/service"
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"google.golang.org/grpc"
)

func InitGRPCRouter(
	server *grpc.Server,
	conf *config.AppConfig,

	// register all GRPC service here
	helloWorldService service.HelloWorldService,
) {
	// register all services here
	helloWorldServer.RegisterHelloWorldServiceServer(server, helloWorldService)
}
