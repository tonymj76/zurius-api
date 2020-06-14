package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"googlemaps.github.io/maps"
)

var api string

func init() {
	api = os.Getenv(APIKEY)
	SetMode(api)
}

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
		fields = "photos,formatted_address,name,rating"
		apiKey = api
	)
	if apiKey == "" {
		log.Fatal("there must be a api key")
	}
	input := c.Query("input")
	radius, err := strconv.Atoi(c.Query("radius"))
	if err != nil {
		radius = 0
	}
	if input == "undefined" {
		c.JSON(http.StatusBadRequest, ginH("failed", errors.New("you need inputs")))
		return
	}

	client, err = maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, ginH("failed to create new client", err))
		return
	}

	r := &maps.FindPlaceFromTextRequest{
		Input:     input,
		InputType: maps.FindPlaceFromTextInputTypeTextQuery,
	}

	f, err := parseFields(fields)
	check(c, err)
	r.Fields = f
	if radius != 0 {
		r.LocationBiasRadius = radius
	}

	resp, err := client.FindPlaceFromText(context.Background(), r)
	check(c, err)

	c.JSON(http.StatusOK, ginH(resp, "success"))
	return
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
