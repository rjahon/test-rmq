package api

import (
	V1 "github.com/rjahon/labs-rmq/storage/api/V1"
	"github.com/rjahon/labs-rmq/storage/config"
	"github.com/rjahon/labs-rmq/storage/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := cors.DefaultConfig()

	cfg.AllowHeaders = append(cfg.AllowHeaders, "*")
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true

	router.Use(cors.New(cfg))

	handlerV1 := V1.New(&V1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	// Phone --->
	router.GET("/phone/:id", handlerV1.GetPhone)
	// <---

	return router
}
