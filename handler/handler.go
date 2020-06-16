package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

//RequestToTomTom _
func RequestToTomTom(c *gin.Context) {
	input := c.Query("input")
	radius := c.Query("radius")

	if input == "undefined" {
		c.JSON(http.StatusBadRequest, ginH("failed", errors.New("you need inputs")))
		return
	}
	if radius == "undefined" {
		radius = ""
	}
	APIKey := os.Getenv("APIKEY")
	url := fmt.Sprintf("https://api.tomtom.com/search/2/search/%s.json?key=%s&countrySet=NG&lat=37.8085&lon=-122.423&radius=%s", input, APIKey, radius)
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")
	check(c, err)
	resp, err := client.Do(request)
	check(c, err)
	c.JSON(http.StatusOK, ginH(resp, "success"))
	return
}
