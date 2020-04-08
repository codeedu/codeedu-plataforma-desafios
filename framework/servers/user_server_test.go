package servers_test

import (
	"context"
	"github.com/codeedu/codeedu/appllication/repositories"
	"github.com/codeedu/codeedu/appllication/usecases"
	"github.com/codeedu/codeedu/framework/pb"
	"github.com/codeedu/codeedu/framework/servers"
	"github.com/codeedu/codeedu/framework/utils"
	"github.com/codeedu/codeedu/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

var db *gorm.DB

func TestServerCreateUser(t *testing.T) {
	db = utils.ConnectDB()
	t.Parallel()

	user := domain.NewUser()
	user.ID = ""
	user.Name = "Wesley"
	user.Email = "wes@wesli"+uuid.NewV4().String()+"d.com"

	userNoID := domain.NewUser()
	userNoID.ID = ""
	userNoID.Name = "Wesley"
	userNoID.Email = "wes@wesli"+uuid.NewV4().String()+"d.com"

	userInvalidID := domain.NewUser()
	userInvalidID.ID = "invalid"
	userInvalidID.Name = "Wesley Invalid"
	userInvalidID.Email = "wes@wesli"+uuid.NewV4().String()+"d.com"

	repo := repositories.UserRepositoryDb{Db: db}

	testCases := []struct {
		name    string
		user    domain.User
		service usecases.UserUseCase
		code    codes.Code
	}{
		{
			name:    "success_with_id",
			user:    *user,
			code:    codes.OK,
			service: usecases.UserUseCase{UserRepository: repo},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &pb.UserRequest{
				Name:     tc.user.Name,
				Email:    tc.user.Email,
				Password: tc.user.Password,
			}

			server := servers.NewUserServer()
			server.UserUseCase = tc.service
			server.User = tc.user

			res, err := server.CreateUser(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Token)
				if len(tc.user.ID) > 0 {
					require.Equal(t, tc.user.Token, res.GetToken())
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}

}
