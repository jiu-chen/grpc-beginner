package main

import (
	"context"
	"demo04/services"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8094", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	prodClient := services.NewProdServiceClient(conn)
	ctx := context.Background()
	resp1, _ := prodClient.GetProdStock(ctx, &services.ProdRequest{ProdId: 1})

	log.Println("get prod stock resp: ", resp1)
}
