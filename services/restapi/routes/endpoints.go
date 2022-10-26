package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rintik-io/rintik-auth/configs"
	"github.com/rintik-io/rintik-auth/services/restapi/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Endpoints struct {
	Router    *gin.Engine
	BasePath  string
	MainGroup *gin.RouterGroup
}

func NewEndpoints(routes *gin.Engine, basePath string) (*Endpoints, error) {
	if routes == nil {
		return nil, fmt.Errorf("routes must not null")
	}
	return &Endpoints{Router: routes, BasePath: basePath}, nil
}

// LoadEndpoints : Function for all possible route
func (p *Endpoints) LoadEndpoints() {
	// Main Router Group
	p.MainGroup = p.Router.Group("/").Group(p.BasePath)
	{
		p.MainGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		p.MainGroup.GET("/heartbeat", (&controllers.Heartbeat{}).GetVersion)
		// If prometheus set to true in configs, then open prometheus endpoints
		if configs.Properties.Services.Prometheus.Status == true {
			p.MainGroup.GET("/metrics", gin.WrapH(promhttp.Handler()))
		}
		p.MainGroup.POST("/register", (&controllers.Auth{}).Register)
		p.MainGroup.POST("/claims", (&controllers.Auth{}).Claims)
		p.MainGroup.POST("/validate", (&controllers.Auth{}).Validate)
	}

	// Set Komoditas Endpoints
	// p.KomoditasEndpoints()
}
