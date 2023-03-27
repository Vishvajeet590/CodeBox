package main

import (
	"CodeBox/api/handlers"
	"CodeBox/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.Default()) //allows all origin
	v1 := router.Group("/api/v1/")
	{
		jobApi := v1.Group("/runner/")
		{
			jobApi.POST("/runCode", handlers.CreateJob)
			jobApi.GET("/result/:id", handlers.GetJobResult)
		}

		questionApi := v1.Group("/question/")
		{
			questionApi.POST("/add", middleware.TokenAuthMiddleware("admin"), handlers.AddProblem)
			questionApi.GET("/:id", handlers.GetProblemById)
			questionApi.GET("/list", handlers.GetProblems)
		}
		adminApi := v1.Group("/admin/")
		{
			adminApi.POST("/login", handlers.AdminSignIn)
			adminApi.POST("/signup", handlers.AdminSignUp)
		}
		v1.GET("/ping", handlers.Ping)
	}

	return router
}
