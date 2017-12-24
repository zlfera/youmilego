package router

import (
	"os"
	"youmilego/handlers"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", handlers.HelloPage)
		v1.GET("/hello/:name", handlers.HelloName)
		v1.GET("/welcome", handlers.WelcomePage)
	}
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
