package servers

import (
	"context"
	"github.com/codeedu/codeedu/appllication/usecases"
	"github.com/codeedu/codeedu/framework/pb"
	"github.com/codeedu/codeedu/domain"
)

type UserServer struct {
	User        domain.User
	UserUseCase usecases.UserUseCase
}

func (UserServer *UserServer) CreateUser(context context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	UserServer.User.Name = req.GetName()
	UserServer.User.Email = req.GetEmail()
	UserServer.User.Password = req.GetPassword()

	user, err := UserServer.UserUseCase.Create(&UserServer.User)

	if err != nil {

	}

	return &pb.UserResponse{Token: user.Token}, nil
}

func NewUserServer() *UserServer {
	return &UserServer{}
}
