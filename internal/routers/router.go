package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jellydelete/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// health
	r.GET("/metrics", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.GET("/start-level", controllers.GetInitPattern)
	return r
}

func startLevel() string {
	return "ok"
}
