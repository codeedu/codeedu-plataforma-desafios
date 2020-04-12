package servers

import (
	"context"
	"github.com/codeedu/codeedu-plataforma-desafios/application/usecases"
	"github.com/codeedu/codeedu-plataforma-desafios/domain"
	"github.com/codeedu/codeedu-plataforma-desafios/framework/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChallengeServer struct {
	ChallengeUseCase usecases.ChallengeUseCase
}

func NewChallengeServer() *ChallengeServer {
	return &ChallengeServer{}
}

func (ChallengeServer *ChallengeServer) CreateChallenge(ctx context.Context, req *pb.NewChallengeRequest) (*pb.NewChallengeResponse, error) {

	challenge := domain.NewChallenge()
	challenge.Title = req.GetChallenge().GetTitle()
	challenge.Slug = req.GetChallenge().GetSlug()
	challenge.Description = req.GetChallenge().GetDescription()
	challenge.Requirements = req.GetChallenge().GetRequirements()
	challenge.Tags = req.GetChallenge().GetTags()
	challenge.Level = req.GetChallenge().GetLevel()

	var files []*domain.ChallengeFile
	for _, f := range req.Challenge.GetFile() {
		file, err := domain.NewChallengeFile(f.Name, f.Url)

		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Error during the validation: %v", err)
		}
		file.ChallengeID = challenge.ID
		files = append(files, file)
	}

	challenge.ChallengeFiles = files

	if err := challenge.Valid(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error during the validation: %v", err)
	}

	newChallenge, err := ChallengeServer.ChallengeUseCase.Create(challenge)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error persisting information: %v", err)
	}

	return &pb.NewChallengeResponse{
		ChallengeId: newChallenge.ID,
	}, nil
}
