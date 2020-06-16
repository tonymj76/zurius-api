package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var url = "https://api.tomtom.com/search/2/search/Lagos Executive Cardiovascular Clinic.json?key=Sjt5p5YqryGxPvW8Wepvkq0nkSlo7iVc&countrySet=NG&lat=37.8085&lon=-122.423"

func main() {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	file, err := os.OpenFile("hospital.json", os.O_CREATE|os.O_RDWR, 0766)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(request)
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	js, err := json.MarshalIndent(data, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(file, bytes.NewBuffer(js))
	// strings.NewReader(string(js))
}
