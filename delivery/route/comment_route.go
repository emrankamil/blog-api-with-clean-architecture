package route

import (
	"blog-api_with-clean-architecture/bootstrap"
	"blog-api_with-clean-architecture/delivery/controller"
	"blog-api_with-clean-architecture/delivery/middleware"
	"blog-api_with-clean-architecture/domain"
	"blog-api_with-clean-architecture/mongo"
	"blog-api_with-clean-architecture/repository"
	"blog-api_with-clean-architecture/usecase"

	"github.com/gin-gonic/gin"
)

func NewCommentRouter(env *bootstrap.Env, db mongo.Database, gin *gin.Engine) {
	tr := repository.NewCommentRepository(db, domain.CollectionComments)
	br := repository.NewBlogRepository(db, domain.CollectionBlogs)
	ur := repository.NewUserRepository(db, domain.UserCollection)
	cu := usecase.NewCommentUseCase(tr, br, ur)
	cc := controller.CommentController{
		CommentUseCase: cu,
	}
	protectedRoute := gin.Group("")
	publicRoute := gin.Group("")
	protectedRoute.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	publicRoute.GET("/comments/:comment_id", cc.GetComment)
	protectedRoute.POST("/comments/:blog_id", cc.CreateComment)
	publicRoute.PUT("/comments/:comment_id", cc.UpdateComment)
	publicRoute.DELETE("/comments/:comment_id", cc.DeleteComment)

}
