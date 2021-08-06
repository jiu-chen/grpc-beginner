package main

import (
	"context"
	"fmt"
	"log"

	"demo07/services"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8097", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	userClient := services.NewUserServiceClient(conn)
	ctx := context.Background()
	stream, err := userClient.GetUserScoreByClientSteam(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var i int32
	for j := 1; j <= 3; j++ {
		req := services.UserScoreRequest{}
		req.Users = make([]*services.UserInfo, 0)
		for i = 1; i <= 5; i++ {
			req.Users = append(req.Users, &services.UserInfo{UserId: i})
		}
		err = stream.Send(&req)
		if err != nil {
			log.Fatal(err)
		}
	}
	resp, _ := stream.CloseAndRecv()
	fmt.Println("client stream resp:\n", resp)
}
