package main

import (
	"context"
	"demo04/services"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.Dial(":8094", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	orderClient := services.NewOrderServiceClient(conn)
	ctx := context.Background()
	resp1, _ := orderClient.NewOrder(ctx,
		&services.OrderRequest{
			OrderMain: &services.OrderMain{
				OrderId:    1001,
				OrderNo:    "1",
				OrderMoney: 99.99,
				OrderTime:  timestamppb.New(time.Now()),
				UserId:     123,
			},
		},
	)
	log.Println("get order resp: ", resp1)
}
