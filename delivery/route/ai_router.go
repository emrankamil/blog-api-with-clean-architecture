package route

import (
	"blog-api_with-clean-architecture/bootstrap"
	"blog-api_with-clean-architecture/delivery/controller"
	"blog-api_with-clean-architecture/infrastructure"
	"blog-api_with-clean-architecture/mongo"
	"blog-api_with-clean-architecture/repository"
	"blog-api_with-clean-architecture/usecase"
	"blog-api_with-clean-architecture/utils"
	"time"

	"github.com/gin-gonic/gin"
)


func NewAiRouter(env *bootstrap.Env,timeout time.Duration,db mongo.Database,group *gin.RouterGroup){
	air := repository.NewAIRepository(db)
	Llc := infrastructure.NewLlmClient(utils.MESSAGE_TELL_ROLE)
	aiu := usecase.NewChatUseCase(air,Llc)
	aic := controller.NewAIController(aiu)

	group.GET("/chat", aic.GetChats)
	group.GET("/chat/:id", aic.GetChat)
	group.POST("/chat", aic.CreateChat)
	group.PUT("/chat/:id", aic.UpdateChat)
	// group.DELETE("/chat/:id", aic.DeleteChat)
}