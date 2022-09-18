package api

import (
	"github.com/gin-gonic/gin"
	api "github.com/nuareal/matching-partner/api/handler"
)

// SetupRouter setup api endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/matchingPartners/", api.GetMatchingPartners)
		v1.GET("/partner/:nameId", api.GetPartner)

	}
	return router
}
