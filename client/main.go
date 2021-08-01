package main

import (
	"context"
	"fmt"
	"grpc_server/services"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("./keys/ssl.crt", "jtthink.com")
	if err != nil {
		log.Fatal(err)
	}

	// conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal("err")
	}
	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)
	resp, err := prodClient.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId: 12, ProdArea: services.ProdArea_C})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("grpc get stock: %d\n", resp.ProdStock)

	response, err := prodClient.GetProdStocks(context.Background(), &services.QuerySize{Size: 1})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("grpc list stock: %v\n", response.Prodres)

	ctx := context.Background()
	modelResp, err := prodClient.GetProdInfo(ctx, &services.ProdRequest{ProdId: 1})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("get prod model: %v\n", modelResp)

	orderClient := services.NewOrderServiceClient(conn)
	orderResp, err := orderClient.NewOrder(ctx,
		&services.OrderMain{
			OrderId:    1,
			OrderNo:    "1000",
			UserId:     100,
			OrderMoney: 99.99,
			OrderTime:  timestamppb.New(time.Now().UTC()),
		})
	if err != nil {
		log.Fatal(err)
	}
	// get order resp: status:"ok"  message:"success
	fmt.Printf("get order resp: %v\n", orderResp)
}

/*
grpc get stock: 33
grpc list stock: [prod_stock:1 prod_stock:2 prod_stock:3 prod_stock:4]
get prod model: prod_id:1  prod_name:"phone"  prod_price:3999
get order resp: status:"ok"  message:"success
*/
