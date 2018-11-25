package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"sk.golang.RpcAndHttp/streamsvc"
)

func main() {
	// Create a random number of random points
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//pointCount := int(r.Int31n(100)) + 2 // Traverse at least two points
	//var points []*pb.Point
	//for i := 0; i < pointCount; i++ {
	//	points = append(points, randomPoint(r))
	//}
	//log.Printf("Traversing %d points.", len(points))

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err == nil {
		client := streamsvc.NewStreamingSvcClient(conn)
		stream, err := client.Get(context.Background())
		stream.Send(&streamsvc.Chat{})
		if err == nil {
			for {
				chat, err := stream.Recv()
				if err == io.EOF {
					break
				}
				println(chat.Message)
			}
		}

	}
	defer conn.Close()

}
