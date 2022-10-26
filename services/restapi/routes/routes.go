package routes

import (
	"fmt"

	"github.com/fahmyabdul/golibs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Routes :
type Routes struct {
	Router    *gin.Engine
	Port      string
	BasePath  string
	MainGroup *gin.RouterGroup
}

// Initialize :
func (a *Routes) Initialize() error {
	// Set default writer to the golibs log writer
	gin.DefaultWriter = golibs.Log.Writer()

	// Disable debug mode
	gin.SetMode(gin.ReleaseMode)

	a.Router = gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	a.Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s | %d | %s | %s | %s %s | %s | %s %s\n",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.Request.Host,
			param.Method,
			param.Path,
			param.Request.UserAgent(),
			param.Request.Proto,
			param.ErrorMessage,
		)
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	a.Router.Use(gin.Recovery())

	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	a.Router.Use(cors.Default())

	// Set Endpoints
	setEndpoints, err := NewEndpoints(a.Router, a.BasePath)
	if err != nil {
		return err
	}
	setEndpoints.LoadEndpoints()

	a.Router.Run(":" + a.Port)

	return nil
}
