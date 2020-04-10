package servers

import (
	"context"
	"github.com/codeedu/codeedu-plataforma-desafios/application/usecases"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/codeedu/codeedu-plataforma-desafios/framework/pb"
	"log"
)

type UserServer struct {
	User        domain.User
	UserUseCase usecases.UserUseCase
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (UserServer *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	UserServer.User.Name = req.GetName()
	UserServer.User.Email = req.GetEmail()
	UserServer.User.Password = req.GetPassword()

	user, err := UserServer.UserUseCase.Create(&UserServer.User)

	if err != nil {
		log.Fatalf("Error during the RPC Create User: %v", err)
	}

	return &pb.UserResponse{
		Token: user.Token,
	}, nil
}
