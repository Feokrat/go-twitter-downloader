package http

import (
	v1 "go-twitter-downloader/pkg/delivery/http/v1"
	"go-twitter-downloader/pkg/service"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) Init(host string, port string) *gin.Engine {
	router := gin.Default()

	corsMiddleware := func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	}

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
		JSONAppErrorReporter(),
	)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
