package main

import (
	"assignment/commons/bootstrap"
	"assignment/interfaces/http/api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

//	@title			Assignment Notes API
//	@version		1.0
//	@description	This is a Notes API

// @host							localhost:5000
// @securitydefinitions.bearerauth	BearerAuth
// @externalDocs.description		OpenAPI
// @externalDocs.url				https://swagger.io/resources/open-api/
func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.DB

	timeout := time.Duration(env.ContextTimeout) * time.Second
	router := gin.Default()
	routes.Setup(env, timeout, db, router)
	swagger := router.Group("/swagger")
	{
		swagger.StaticFile("/doc.json", "./docs/swagger.json")
	}
	swaggerUrl := ginSwagger.URL("http://localhost:5000/swagger/doc.json")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	router.Run(env.ServerAddr)
}
