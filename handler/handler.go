package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

//Location _
type Location struct {
	Results []interface{}
	Summary map[string]interface{}
}

func ginH(msg, in interface{}) gin.H {
	switch in.(type) {
	case error:
		return gin.H{"Message": msg, "Error": in.(error).Error()}
	default:
		return gin.H{"Message": msg, "Success": in}
	}
}

//RequestToTomTom _
func RequestToTomTom(c *gin.Context) {
	var (
		url      string
		location = Location{}
	)
	input := c.Query("input")
	radius := c.Query("radius")
	APIKey := os.Getenv("APIKEY")

	if input == "undefined" {
		c.JSON(http.StatusBadRequest, ginH("failed", errors.New("you need inputs")))
		return
	}
	if radius == "undefined" || radius == "" {
		url = fmt.Sprintf("https://api.tomtom.com/search/2/search/%s.json?key=%s&countrySet=NG&lat=37.8085&lon=-122.423", input, APIKey)
	} else {
		url = fmt.Sprintf("https://api.tomtom.com/search/2/search/%s.json?key=%s&countrySet=NG&lat=37.8085&lon=-122.423&radius=%s", input, APIKey, radius)
	}
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, ginH("failed to fetch request", err))
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginH("failed to fetch repsonse", err))
		return
	}
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		c.JSON(http.StatusBadRequest, ginH("failed to fetch request", err))
		return
	}
	c.JSON(http.StatusOK, location)
	return
}
