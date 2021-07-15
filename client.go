package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type data struct {
	Main  rectangle   `json: main`
	Input []rectangle `json: input`
}

type rectangle struct {
	X            int       `json: x`
	Y            int       `json: y`
	Width        int       `json: width`
	Height       int       `json: height`
	CreationTime time.Time `json: time`
}

type loc struct {
	Lat float32 `json: lat`
	Lon float32 `json: lon`
}

func client() {

	clientData := data{
		Main:  rectangle{0, 0, 10, 20, time.Now().Local()},
		Input: []rectangle{{2, 18, 5, 4, time.Now().Local()}, {12, 18, 5, 4, time.Now().Local()}},
	}
	clientDataJson, err := json.Marshal(clientData)
	req, err := http.NewRequest("GET", "http://localhost:8088", bytes.NewBuffer(clientDataJson))
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
