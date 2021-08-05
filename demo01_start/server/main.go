package main

import (
	"demo01/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, _ := net.Listen("tcp", ":8091")
	log.Println("grpc server starts...")
	rpcServer.Serve(lis)
}
