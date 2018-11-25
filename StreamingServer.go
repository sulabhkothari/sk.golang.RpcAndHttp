package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sk.golang.RpcAndHttp/streamsvc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", "8080"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	streamsvc.RegisterStreamingSvcServer(grpcServer, &strSvc{})
	grpcServer.Serve(lis)
}

type strSvc struct {
}

func (server *strSvc) Get(strm streamsvc.StreamingSvc_GetServer) error {
	strm.Recv()
	if err := strm.Send(&streamsvc.Chat{Id: int32(90), Message: string("initiate stream.....")}); err != nil {
		return err
	}
	if err := strm.Send(&streamsvc.Chat{Id: int32(91), Message: string(" stream")}); err != nil {
		return err
	}
	if err := strm.Send(&streamsvc.Chat{Id: int32(92), Message: string("end stream")}); err != nil {
		return err
	}

	return nil
}
