package main

import (
	"fmt"
	"grpc_server/services"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("./keys/ssl.crt", "./keys/ssl_no_password.key")
	if err != nil {
		log.Fatal(err)
	}

	// no cert
	// rpcServer := grpc.NewServer())

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	// 注册商品服务
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	// 注册订单服务
	services.RegisterOrderServiceServer(rpcServer, new(services.OrderService))

	// lis, _ := net.Listen("tcp", ":8081")
	fmt.Println("grpc server starts...")
	// rpcServer.Serve(lis)

	//提供http服务
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("client request: ", request)
		rpcServer.ServeHTTP(writer, request)
	})
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	// httpServer.ListenAndServe()
	httpServer.ListenAndServeTLS("./keys/ssl.crt", "./keys/ssl_no_password.key")
}
