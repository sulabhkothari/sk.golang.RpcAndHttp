package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"sk.golang.RpcAndHttp/svcdef"
)

type basicServer struct{}

func (server *basicServer) Search(ctxt context.Context, request *svcdef.SearchRequest) (*svcdef.SearchResponse, error) {
	println("NewRequest________________"+request.Query)
	res := svcdef.SearchResponse{}
	res.Query = "HW"
	return &res,nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", "8080"))
	if err != nil {
		//logrus.Errorf("Error listening %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	println("starting server...")
	svcdef.RegisterSearchServiceServer(grpcServer, &basicServer{})
	println("to listen...")
	grpcServer.Serve(lis)
	println("listening")
}
