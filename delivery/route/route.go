package route

import (
	"blog-api_with-clean-architecture/bootstrap"
	"blog-api_with-clean-architecture/delivery/middleware"
	"blog-api_with-clean-architecture/mongo"
	"blog-api_with-clean-architecture/redis"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine, redisClient redis.Client) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewFogetPWRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	NewAiRouter(env, timeout, db, publicRouter)

	NewBlogRouter(db, gin, redisClient)
	NewCommentRouter(env, db, gin)
	NewReactionRouter(env, db, gin)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewLogoutRouter(env, timeout, db, protectedRouter)
	NewPromoteRouter(env, timeout, db, protectedRouter)
}
