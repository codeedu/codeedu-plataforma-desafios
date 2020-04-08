package server

import (
	"github.com/codeedu/codeedu/application/repositories"
	"github.com/codeedu/codeedu/framework/utils"
	"github.com/codeedu/codeedu/domain"
	"github.com/jessevdk/go-flags"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var configOptions utils.ConfigOptions
var parser = flags.NewParser(&configOptions, flags.Default)
var db *gorm.DB

func main() {
	db = utils.ConnectDB()

	user := domain.NewUser()
	user.Name = "Wes"
	user.Email = "wes@wes" + uuid.NewV4().String() + ".com"
	user.Password = "123"

	repo := repositories.UserRepositoryDb{Db: db}
	repo.Insert(user)

	//utils.ConnectDB()
	//db = utils.ConnectDB()
	//utils.CheckConfigFlags(parser)

	//usecases.Login{}.Auth("wesleywillians@gmail.com","123")
}
