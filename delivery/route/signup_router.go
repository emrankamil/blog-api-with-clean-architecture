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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	emailService := infrastructure.NewEmailService(env.SmtpServer, env.Mail, env.MailPassword)
	ur := repository.NewUserRepository(db, domain.UserCollection)
	su := usecase.NewSignupUsecase(ur, timeout, *emailService)
	uu := usecase.NewUserUsecase(ur, timeout)
	sc := controller.NewSignupController(uu, su, env)

	group.POST("/signup", sc.Signup)
	group.POST("/verify_email", sc.VerifyEmail)
}