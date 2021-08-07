package main

import (
	"context"
	"demo08/services"
	"fmt"
	"log"

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
	stream, err := userClient.GetUserScoreByCSSteam(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var i, uid int32
	uid = 1
	for j := 1; j <= 3; j++ {
		req := services.UserScoreRequest{}
		req.Users = make([]*services.UserInfo, 0)
		for i = 1; i <= 5; i++ {
			req.Users = append(req.Users, &services.UserInfo{UserId: uid})
			uid++
		}
		err = stream.Send(&req)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("client server stream batch resp:\n", resp)
	}
}
