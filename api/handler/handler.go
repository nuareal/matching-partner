package api

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/nuareal/matching-partner/db"
	"github.com/nuareal/matching-partner/model"
	"golang.org/x/exp/slices"
)

// GetMatchingPartners		 godoc
// @Summary      Get Matching Partners
// @Description  Returns an order list of partners by most recommended.
// @Tags         partner
// @Produce      json
// @Param        customer  body     model.Customer true  "Costumer JSON"
// @Success      200   {array}  model.Partner
// @Router       /matchingPartners [post]
func GetMatchingPartners(c *gin.Context) {

	var matchingPartners []model.Partner
	distanceToCustomer := make(map[string]float64)
	costumer := model.Customer{}

	if err := c.BindJSON(&costumer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err, v := costumer.Valid(); !v {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	for _, partner := range db.Partners {
		if slices.Contains(partner.Materials, costumer.Material) {
			distance := model.Distance(partner.Address, costumer.Address)

			if distance < partner.OperatingRadius {
				distanceToCustomer[partner.NameId] = distance
				matchingPartners = append(matchingPartners, partner)
			}
		}
	}

	if len(matchingPartners) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no matching partners"})
		return
	}

	sort.Slice(matchingPartners, func(i, j int) bool {
		if matchingPartners[i].Rating != matchingPartners[j].Rating {
			return matchingPartners[i].Rating > matchingPartners[j].Rating
		}
		return distanceToCustomer[matchingPartners[i].NameId] < distanceToCustomer[matchingPartners[j].NameId]
	})

	var nameIds []string
	for _, partner := range matchingPartners {
		nameIds = append(nameIds, partner.NameId)
	}

	c.JSON(http.StatusOK, nameIds)
}

// GetPartner    godoc
// @Summary      Get Partner Info
// @Description  Returns info of a given partner.
// @Tags         partner
// @Produce      json
// @Param	     nameId	path	string true	"Partner nameId"
// @Success      200  {object}  model.Partner
// @Router       /partner/{nameId} [get]
func GetPartner(c *gin.Context) {

	if _, ok := db.Partners[c.Param("nameId")]; ok {
		c.JSON(http.StatusOK, db.Partners[c.Param("nameId")])
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "partner not found"})
}
