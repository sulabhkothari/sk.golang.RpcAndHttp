package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sk.golang.RpcAndHttp/svcdef"
)

type basicClient struct{}

func (client *basicClient) Search(ctx context.Context, in *svcdef.SearchRequest, opts grpc.CallOption) (*svcdef.SearchResponse, error) {
	println("Hi")
	conn, err := grpc.Dial(":9191", grpc.WithInsecure())
	if err == nil {
		println("GETTER")
		client1 := svcdef.NewSearchServiceClient(conn)
		res, err := client1.Search(ctx, in, opts)
		println("DONE")
		if err == nil {
			println("NoERR")
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
	println("Hi")
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	println(err)
	if err == nil {
		println("GETTER")
		client := svcdef.NewSearchServiceClient(conn)
		println("Setter")
		if client == nil {
			println("NIL CLient")
			return
		}
		res, err := client.Search(context.Background(), &svcdef.SearchRequest{})
		println("DONE")
		if err == nil {
			println("NoERR")
			println(res.Query)
		} else {
			println("_____________________Er")

			//println(err)
		}
	}
	defer conn.Close()
}
