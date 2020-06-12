package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tonymj76/zurius-api/models"
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

func check(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, ginH("failed to fetch request", err))
		return
	}
}

//GooglePlace _
func GooglePlace(c *gin.Context) {
	var (
		client *maps.Client
		err    error
		apiKey = ""
	)
	place := &models.Place{}

	if err := c.ShouldBindJSON(place); err != nil {
		c.JSON(http.StatusBadRequest, ginH("failed to bind request", err))
		return
	}

	client, err = maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, ginH("failed to create new client", err))
		return
	}

	r := &maps.FindPlaceFromTextRequest{
		Input:     place.Input,
		InputType: maps.FindPlaceFromTextInputTypeTextQuery,
	}

	if place.LocationBias != "" {
		lb, err := maps.ParseFindPlaceFromTextLocationBiasType(place.LocationBias)
		if err != nil {
			c.JSON(http.StatusBadRequest, ginH("failed to get location bias", err))
			return
		}
		r.LocationBias = lb
		switch lb {
		case maps.FindPlaceFromTextLocationBiasPoint:
			l, err := maps.ParseLatLng(place.Point)
			check(c, err)
			r.LocationBiasPoint = &l
		case maps.FindPlaceFromTextLocationBiasCircular:
			l, err := maps.ParseLatLng(place.Center)
			check(c, err)
			r.LocationBiasCenter = &l
			r.LocationBiasRadius = place.Radius
		case maps.FindPlaceFromTextLocationBiasRectangular:
			sw, err := maps.ParseLatLng(place.SouthWest)
			check(c, err)
			r.LocationBiasSouthWest = &sw
			ne, err := maps.ParseLatLng(place.NorthEast)
			check(c, err)
			r.LocationBiasNorthEast = &ne
		}
	}

	if place.Fields != "" {
		f, err := parseFields(place.Fields)
		check(c, err)
		r.Fields = f
	}

	resp, err := client.FindPlaceFromText(context.Background(), r)
	check(c, err)

	c.JSON(http.StatusOK, ginH(resp, "success"))
}

func parseFields(fields string) ([]maps.PlaceSearchFieldMask, error) {
	var res []maps.PlaceSearchFieldMask
	for _, s := range strings.Split(fields, ",") {
		f, err := maps.ParsePlaceSearchFieldMask(s)
		if err != nil {
			return nil, err
		}
		res = append(res, f)
	}
	return res, nil
}
