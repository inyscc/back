package routes

import (
	"github.com/SunspotsInys/thedoor/services"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	e := gin.Default()

	api := e.Group("/api")
	api.Use(services.ParseJWT())
	{
		api.GET("/postTot", services.GetPostTotal)
		api.GET("/posts", services.GetPostList)
		api.GET("/post", services.GetPostDetail)
		api.GET("/tags", services.GetTagList)
		api.GET("/comments", services.GetComments)
		api.POST("/comment", services.NewComments)

		api.POST("/signin", services.Signin)

		admin := api.Group("/admin")
		admin.Use(services.JudgeAdmin())
		{
			admin.GET("/sysinfo", services.GetSysInfo)
			admin.GET("/post", services.GetPostSimpleList)
			admin.POST("/post", services.NewPost)
		}
	}

	return e
}
