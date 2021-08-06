package services

import "io"

type UserService struct{}

func (*UserService) GetUserScoreByClientSteam(stream UserService_GetUserScoreByClientSteamServer) error {

	var score int32
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&UserScoreResponse{Users: users})
		}
		if err != nil {
			return err
		}
		for _, user := range req.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
	}
}
