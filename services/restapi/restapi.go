package restapi

import (
	"github.com/fahmyabdul/golibs"
	"github.com/rintik-io/rintik-auth/app"
	"github.com/rintik-io/rintik-auth/configs"
	"github.com/rintik-io/rintik-auth/services/restapi/docs"
	"github.com/rintik-io/rintik-auth/services/restapi/routes"
)

type Restapi struct{}

func (p Restapi) Start() error {
	config := configs.Properties.Services.Restapi
	golibs.Log.Printf("| RestAPI | Starting | Port: %s, BasePath: %s\n", config.Port, config.BasePath)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = config.Swagger.Title
	docs.SwaggerInfo.Description = config.Swagger.Description
	docs.SwaggerInfo.Version = app.CurrentVersion
	docs.SwaggerInfo.BasePath = config.BasePath
	docs.SwaggerInfo.Schemes = config.Swagger.Schemes

	routes := routes.Routes{
		Port:     config.Port,
		BasePath: config.BasePath,
	}
	err := routes.Initialize()
	if err != nil {
		return err
	}

	return nil
}
