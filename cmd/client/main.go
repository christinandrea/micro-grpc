package main

import (
	protoGen "github.com/christinandrea/micro-grpc/internal/protoGen"

	gRPC "google.golang.org/grpc"

	"context"
	"fmt"
)

func main() {
	serverAddress := "localhost:9000"

	conn, e := gRPC.Dial(serverAddress, gRPC.WithInsecure())

	if e != nil {
		panic(e)
	}

	client := protoGen.NewRepositoryServiceClient(conn)

	for i := range [10]int{} {
		repoModel := protoGen.Repository{
			Id:   uint64(i),
			Name: string("Project PKL"),
		}
		if respMessage, e := client.Add(context.Background(), &repoModel); e != nil {
			panic(fmt.Sprintf("cant insert : %v", e))

		} else {
			fmt.Println("Record inserted.........")
			fmt.Println(respMessage)
		}
	}
}
