package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"googlemaps.github.io/maps"
)

func ginH(msg, in interface{}) gin.H {
	switch in.(type) {
	case error:
		return gin.H{"Message": msg, "Error": in.(error).Error()}
	default:
		return gin.H{"Message": msg, "Success": in}
	}
}

//GooglePlace _
func GooglePlace(c *gin.Context) {
	var (
		client *maps.Client
		err    error
		apiKey string
	)
	place := &models.Place{}

	if err := c.ShouldBindJSON(&place); err != nil {
		c.JSON(http.StatusBadRequest, ginH("failed to bind request", err))
		return
	}

	client, err = maps.NewClient(maps.WithAPIKey(*apiKey))

	r := &maps.FindPlaceFromTextRequest{
		Input:     &,
		InputType: parseInputType(*inputType),
	}
}
