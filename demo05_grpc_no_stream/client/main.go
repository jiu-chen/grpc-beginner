package main

import (
	"context"
	"log"

	"demo05/services"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8095", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	userClient := services.NewUserServiceClient(conn)
	ctx := context.Background()

	var i int32
	req := services.UserScoreRequest{}
	req.Users = make([]*services.UserInfo, 0)
	for i = 0; i < 6; i++ {
		req.Users = append(req.Users, &services.UserInfo{UserId: i})
	}
	resp, err := userClient.GetUserScore(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("get usre store resp: ", resp.Users)

}
