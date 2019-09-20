package main

import (
	pb "github.com/christinandrea/micro-grpc/internal/impl"
	protoGen "github.com/christinandrea/micro-grpc/internal/protoGen"

	gRPC "google.golang.org/grpc"

	"fmt"
	"log"
	"net"
)

func main() {
	netListener := getNetListener(9000)
	server := gRPC.NewServer()

	repositoryImpl := pb.NewRepoImpl()
	protoGen.RegisterRepositoryServiceServer(server, repositoryImpl)

	if err := server.Serve(netListener); err != nil {
		log.Fatalf("failed to start :%s", err)
	}
}

func getNetListener(port uint) net.Listener {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed :%v", err)
		panic(fmt.Sprintf("failed to listen :%v", err))
	}
	return conn
}
