package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type weatherData struct {
	LocationName string   `json: locationName`
	Weather      string   `json: weather`
	Temperature  int      `json: temperature`
	Celsius      bool     `json: celsius`
	TempForecast []int    `json: temp_forecast`
	Wind         windData `json: wind`
}

type windData struct {
	Direction string `json: direction`
	Speed     int    `json: speed`
}
type loc struct {
	Lat float32 `json: lat`
	Lon float32 `json: lon`
}

func client() {

	locJson, err := json.Marshal(loc{Lat: 35.14326, Lon: -116.104})
	req, err := http.NewRequest("POST", "http://localhost:8088", bytes.NewBuffer(locJson))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}

func main() {
	client()
}
