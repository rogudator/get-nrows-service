package transport

import (
	"github.com/gin-gonic/gin"
	_ "github.com/rogudator/get-nrows-service/docs"
	"github.com/rogudator/get-nrows-service/internal/service"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Transport struct {
	serv *service.Service
}

func NewTransport(serv *service.Service) *Transport {
	return &Transport{
		serv: serv,
	}
}

func (t *Transport) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/rows", t.getRows)

	return router
}
