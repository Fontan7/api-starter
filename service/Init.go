package service

import (
	"api-starter/controller"
	"api-starter/database"
	"api-starter/internal"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(app internal.App, db *database.DbConn) *gin.Engine {
	router := gin.Default()
	router.Use(setDb(db))
	router.Use(setApp(app))
	router.Use(corsConfig(router))
	router.Use(gin.Logger())
	makeRoutes(router, app)

	return router
}

// makeRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
func makeRoutes(router *gin.Engine, app internal.App) {
	router.GET("/health", func(c *gin.Context) { gHandler(c, controller.HealthCheck) })
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//we add this here and not in initRouter because health and swagger should be 100% public
	router.Use(checkAPIKey(app.ClientKey()))

	V := router.Group("api-starter/v1")
	//V2 := router.Group("/viber/v2")

	v1Public(V)
	v1Private(V)
}

func v1Public(rg *gin.RouterGroup) {
	rg.GET("/{id}", func(c *gin.Context) { gHandler(c, controller.GetPublicSomething) })
}

func v1Private(rg *gin.RouterGroup) {
	r := rg.Group("/private", validateAndSetToken())

	r.GET("/user/home", func(c *gin.Context) { gHandler(c, controller.GetPrivateSomething) })
}
