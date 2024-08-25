package route

import (
	"blog-api_with-clean-architecture/delivery/controller"
	"blog-api_with-clean-architecture/domain"
	"blog-api_with-clean-architecture/mongo"
	"blog-api_with-clean-architecture/redis"
	"blog-api_with-clean-architecture/repository"
	"blog-api_with-clean-architecture/usecase"

	"github.com/gin-gonic/gin"
)

func NewBlogRouter(db mongo.Database, gin *gin.Engine, redisClient redis.Client) {
	tr := repository.NewBlogRepository(db, domain.CollectionBlogs)
	tu := usecase.NewBlogUseCase(tr)
	tc := controller.NewBlogController(tu, redisClient)

	// protectedRoute := gin.Group("")
	publicRoute := gin.Group("")
	// protectedRoute.Use(infrastructure.AdminOnlMiddleware(), infrastructure.JWTAuthMiddleware())
	publicRoute.GET("/blogs", tc.GetBlogs)
	publicRoute.GET("/blogs/:id", tc.GetBlog)
	publicRoute.POST("/blogs", tc.CreateBlog)
	publicRoute.PUT("/blogs/:id", tc.UpdateBlog)
	publicRoute.DELETE("/blogs/:id", tc.DeleteBlog)
	publicRoute.POST("/blogs/:id/like", tc.LikeBlog)
	publicRoute.POST("/blogs/:id/unlike", tc.UnlikeBlog)
	publicRoute.POST("/blogs/:id/comment", tc.CommentBlog)

}
