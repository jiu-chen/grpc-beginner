package services

import (
	"time"
)

type UserService struct{}

func (*UserService) GetUserScoreByServerSteam(in *UserScoreRequest, stream UserService_GetUserScoreByServerSteamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)

	for index, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)

		if (index+1)%2 == 0 && index > 0 {
			err := stream.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			users = (users)[0:0] //置空
		}
		time.Sleep(time.Second * 1)
	}

	if len(users) != 0 {
		err := stream.Send(&UserScoreResponse{Users: users})
		if err != nil {
			return err
		}
	}
	return nil
}
