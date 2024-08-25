package route

import (
	"blog-api_with-clean-architecture/bootstrap"
	"blog-api_with-clean-architecture/delivery/controller"
	"blog-api_with-clean-architecture/domain"
	"blog-api_with-clean-architecture/infrastructure"
	"blog-api_with-clean-architecture/mongo"
	"blog-api_with-clean-architecture/repository"
	"blog-api_with-clean-architecture/usecase"
	"time"

	"github.com/gin-gonic/gin"
)


func NewFogetPWRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	ur := repository.NewUserRepository(db, domain.UserCollection)
	emailService := infrastructure.NewEmailService(env.SmtpServer, env.Mail, env.MailPassword)

	fpu := usecase.NewForgetPWUsecase(ur, timeout, *emailService)
	uu := usecase.NewUserUsecase(ur, timeout)

	fpc := controller.ForgetPWController{
		Userusecase: uu,
		ForgetPWUsecase: fpu,
		Env: env,
	}

	group.POST("/forget-password", fpc.ForgetPW)
	group.POST("/recover-password", fpc.ResetPW)
}