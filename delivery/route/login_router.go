package route

import (
	"blog-api_with-clean-architecture/bootstrap"
	"blog-api_with-clean-architecture/delivery/controller"
	"blog-api_with-clean-architecture/domain"
	"blog-api_with-clean-architecture/mongo"
	"blog-api_with-clean-architecture/repository"
	"blog-api_with-clean-architecture/usecase"
	"time"

	"github.com/gin-gonic/gin"
)


func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	ur := repository.NewUserRepository(db, domain.UserCollection)
	su := usecase.NewLoginUsecase(ur, timeout)
	uu := usecase.NewUserUsecase(ur, timeout)
	sc := controller.LoginController{
		UserUsecase: uu,
		LoginUsecase: su,
		Env: env,
	}

	group.POST("/login", sc.Login)
}