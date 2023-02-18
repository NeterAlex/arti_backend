package routers

import (
	"arti_backend/middleware"
	"arti_backend/pkg/setting"
	"arti_backend/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())

	gin.SetMode(setting.RunMode)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiV1 := router.Group("/api")
	apiV1.GET("/articles", api.GetArticles)
	apiV1.GET("/articles/:id", api.GetArticle)
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/articles/create", api.AddArticle)
		apiV1.PUT("/articles/:id", api.EditArticle)
		apiV1.DELETE("/articles/:id", api.DeleteArticle)
	}

	auth := router.Group("/api")
	auth.POST("/auth/login", api.GetAuth)
	auth.POST("/auth/register", api.AddUser)
	auth.Use(middleware.JWT())
	{
		auth.PUT("/auth/:id", api.EditUser)
		auth.DELETE("/auth/:username", api.DeleteUser)
		auth.GET("/auth/:username", api.GetUserInfoByName)
	}

	file := router.Group("/api")
	auth.Use(middleware.JWT())
	{
		file.POST("/upload_avatar", api.UploadAvatar)
	}
	return router
}
