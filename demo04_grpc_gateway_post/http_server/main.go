package main

import (
	"context"
	"demo04/services"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	gwmux := runtime.NewServeMux()
	defer cancel()
	opt := []grpc.DialOption{grpc.WithInsecure()}
	//8099是对应的grpc的监听端口，必须与grpc开放的端口一致,所以启动前务必启动grpc服务
	err := services.RegisterProdServiceHandlerFromEndpoint(ctx, gwmux, "localhost:8094", opt)
	if err != nil {
		log.Fatal(err)
	}
	err = services.RegisterOrderServiceHandlerFromEndpoint(ctx, gwmux, "localhost:8094", opt)
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	log.Println("grpc http server starts...")
	httpServer.ListenAndServe()
}
