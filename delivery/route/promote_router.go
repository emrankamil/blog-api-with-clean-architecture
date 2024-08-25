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

func NewPromoteRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserCollection)
	pu := usecase.NewPromoteUsecase(ur, timeout)
	promoteUserController := controller.NewPromoteController(ur, pu, env)

	group.PUT("/promote-user/:id", promoteUserController.PromoteUser)
	group.PUT("/demote-user/:id", promoteUserController.DemoteUser)
}