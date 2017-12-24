package router

import (
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
	r.Run(":5000")
}
