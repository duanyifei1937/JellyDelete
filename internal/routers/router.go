package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"jellydelete/internal/controllers"
	"jellydelete/pkg/prom_metrics"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// health
	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// add ops count test metrics
	go prom_metrics.RecordMetrics()

	r.GET("/start-level", controllers.GetInitPattern)
	r.GET("/move", controllers.MovePattern)

	return r
}
