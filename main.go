package main

import (
	"log"

	api "github.com/nuareal/matching-partner/api/router"
	_ "github.com/nuareal/matching-partner/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Finding Partner
// @version         1.0
// @description     An API to find partner company information.
// @termsOfService

// @contact.name   Nuno Areal
// @contact.email  nuno.areal.it@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {

	router := api.SetupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
