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

func NewReactionRouter(env *bootstrap.Env, db mongo.Database, gin *gin.Engine) {
	rr := repository.NewReactionRepository(db, domain.CollectionReactions)
	br := repository.NewBlogRepository(db, domain.CollectionBlogs)
	// ur := repository.NewUserRepository(db, domain.UserCollection)
	Ru := usecase.NewReactionUseCase(rr, br)
	Rc := controller.ReactionController{
		ReactionUseCase: Ru,
	}
	protectedRoute := gin.Group("")
	// publicRoute := gin.Group("")
	protectedRoute.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	protectedRoute.POST("/blog/:blog_id/like", Rc.LikeBlog)
	protectedRoute.POST("/blog/:blog_id/unlike", Rc.UnLikeBlog)
	protectedRoute.DELETE("/blog/:blog_id/delete", Rc.DeleteLike)

}
