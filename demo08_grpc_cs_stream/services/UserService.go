package services

import (
	"io"
	"log"
)

type UserService struct{}

func (*UserService) GetUserScoreByCSSteam(stream UserService_GetUserScoreByCSSteamServer) error {

	var score int32
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		for _, user := range req.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
		err = stream.Send(&UserScoreResponse{Users: users})
		if err != nil { // 失败
			log.Println(err) //不阻塞下一组send
		}
		users = (users)[0:0]
	}
}
