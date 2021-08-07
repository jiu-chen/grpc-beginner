package main

import (
	"demo08/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterUserServiceServer(rpcServer, new(services.UserService))
	lis, _ := net.Listen("tcp", ":8097")
	log.Println("grpc server starts...")
	rpcServer.Serve(lis)
}
