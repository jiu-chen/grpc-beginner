package main

import (
	"demo03/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	//rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, err := net.Listen("tcp", ":8099")
	log.Println("grpc server starts...")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer.Serve(lis)
}
