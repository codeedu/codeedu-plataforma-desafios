package servers

import (
	"context"
	"github.com/codeedu/codeedu-plataforma-desafios/application/usecases"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/codeedu/codeedu-plataforma-desafios/framework/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	UserUseCase usecases.UserUseCase
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (UserServer *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	user, err := domain.NewUser(req.GetName(), req.GetEmail(), req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error during the validation: %v", err)
	}

	newUser, err := UserServer.UserUseCase.Create(user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error persisting information: %v", err)
	}

	return &pb.UserResponse{
		Token: newUser.Token,
	}, nil
}

func (UserServer *UserServer) Auth(ctx context.Context, req *pb.UserAuthRequest) (*pb.UserResponse, error) {

	user, err := UserServer.UserUseCase.Auth(req.GetEmail(), req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Email and password don't match")
	}

	return &pb.UserResponse{
		Token: user.Token,
	}, nil

}
