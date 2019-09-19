package main

import (
	"internal/protoGen"
	"internal/impl"

	"google.golang.org/gRPC"

	"fmt"
	"log"
	"net"
)

func main() {
	netListener := getNetListener(9000)
	grpcServer := gRPC.NewServer()

	repositoryImpl := impl.NewRepoImpl()
	protoGen.RegisterRepositoryServiceServer(grpcServer, repositoryImpl)

	if err := grpcServer.Serve(netListener); err != nil {
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
