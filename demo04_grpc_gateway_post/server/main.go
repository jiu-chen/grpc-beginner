package main

import (
	"demo04/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	//rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	services.RegisterOrderServiceServer(rpcServer, new(services.OrdersService))
	lis, err := net.Listen("tcp", ":8094")
	log.Println("grpc server starts...")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer.Serve(lis)
}
