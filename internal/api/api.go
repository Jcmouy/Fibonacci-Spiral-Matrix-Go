package api

import (
	"fibonacci-spiral-matrix-go/internal/pkg/tracer"
	"fibonacci-spiral-matrix-go/internal/wired"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

//AppServer application runtime
type AppServer struct {
	Close func() error
	Start func() error
}

func registerHandlers(apiRouteGroup *gin.RouterGroup) error {

	public := apiRouteGroup.Group("/")
	{
		authHandler, err := wired.InitializeAuthHandler()
		if err != nil {
			return err
		}
		authHandler.Register(public)
	}

	protected := apiRouteGroup.Group("/user")
	{
		fiboSpiralMatrixHandler, err := wired.InitializeFiboSpiralMatrixHandler()
		if err != nil {
			return err
		}
		fiboSpiralMatrixHandler.Register(protected)
	}

	healthHandler := wired.InitializeHealthHandler()
	healthHandler.Register(apiRouteGroup)

	return nil
}

func NewAppServer() (*AppServer, error) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	apiRouteGroup := router.Group("/api")

	if err := registerHandlers(apiRouteGroup); err != nil {
		return nil, err
	}

	//Opentracing configuration
	traceCloser := tracer.Init()

	//Swagger Configuration
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))

	return &AppServer{
		Close: func() error {
			return traceCloser.Close()
		},
		Start: func() error {
			return router.Run()
		},
	}, nil
}
