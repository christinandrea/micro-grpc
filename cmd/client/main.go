package main

import (
	protoGen "github.com/christinandrea/micro-grpc/internal/protoGen"
	pb "github.com/christinandrea/micro-grpc/internal/impl"

	gRPC "google.golang.org/gRPC"
	
	"context"
	"fmt"

	
)

func main() {
	serverAddress := "localhost:9000"

	conn, e := gRPC.Dial(serverAddress, gRPC.WithInsecure())

	if err != nil {
		panic(e)
	}

	client := protoGen.NewRepositoryServiceClient(conn)

	for i := range [10]int{} {
		repoModel := protoGen.Repository{
			Id:   uint64(i),
			name: string("grpc-trial"),
		}
		if respMessage, e := client.Add(context.Background(), &repoModel); e != nil {
			panic(fmt.Sprintf("cant insert : %v", e))

		} else {
			fmt.Println("inserting record...")
			fmt.Println(respMessage)
		}
	}
}
