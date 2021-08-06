package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"demo06/services"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8096", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	userClient := services.NewUserServiceClient(conn)
	ctx := context.Background()

	var i int32
	req := services.UserScoreRequest{}
	req.Users = make([]*services.UserInfo, 0)
	for i = 1; i < 8; i++ {
		req.Users = append(req.Users, &services.UserInfo{UserId: i})
	}
	stream, err := userClient.GetUserScoreByServerSteam(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(resp)

		// 这里可以开协程去处理一个批次
		wg.Add(1)
		go func(resp *services.UserScoreResponse) {
			defer wg.Done()
			fmt.Println(resp)
		}(resp)
	}
	wg.Wait() // 等待协程全部完成
	fmt.Println("main exit")

}
