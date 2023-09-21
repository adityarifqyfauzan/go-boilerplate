package service

import (
	pb "github.com/adityarifqyfauzan/go-boilerplate/app/domain/pb/proto/helloworld"
)

type HelloWorldService struct {
	pb.UnimplementedHelloWorldServiceServer
}

func NewHelloWorldService() HelloWorldService {
	return HelloWorldService{}
}

func (s *HelloWorldService) HelloWorld(req *pb.HelloWorldRequest) pb.WebResponse {
	return pb.WebResponse{
		Code:    0,
		Message: "Hello," + req.Name,
	}
}
