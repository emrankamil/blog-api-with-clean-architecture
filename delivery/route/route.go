package route

import (
	"blog-api_with-clean-architecture/bootstrap"
	"blog-api_with-clean-architecture/delivery/middleware"
	"blog-api_with-clean-architecture/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewFogetPWRouter(env, timeout, db, publicRouter)
	
	NewAiRouter(env,timeout,db,publicRouter)

	NewBlogRouter(db, gin)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewLogoutRouter(env, timeout, db, protectedRouter)
	NewPromoteRouter(env, timeout, db, protectedRouter)
}
