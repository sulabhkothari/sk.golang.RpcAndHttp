package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sk.golang.RpcAndHttp/svcdef"
)

type basicClient struct{}

func (client *basicClient) Search(ctx context.Context, in *svcdef.SearchRequest, opts grpc.CallOption) (*svcdef.SearchResponse, error) {
	conn, err := grpc.Dial(":9191", grpc.WithInsecure())
	if err == nil {
		client1 := svcdef.NewSearchServiceClient(conn)
		res, err := client1.Search(ctx, in, opts)
		if err == nil {
			println(res.Query)
			return res, nil
		} else {
			println(err)
		}
	}
	defer conn.Close()
	return nil, err
}

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	println(err)
	if err == nil {
		client := svcdef.NewSearchServiceClient(conn)
		res, err := client.Search(context.Background(), &svcdef.SearchRequest{})
		if err == nil {
			println(res.Query)
		} else {
			println(err)
		}
	}
	defer conn.Close()
}
