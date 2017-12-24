package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}
func HelloName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s", name)
}
func WelcomePage(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "zlf")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "hello %s %s", firstname, lastname)
}
