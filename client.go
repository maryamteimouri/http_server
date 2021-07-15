package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type data struct {
	Main  rectangle   `json: main`
	Input []rectangle `json: input`
}

type rectangle struct {
	X      int `json: x`
	Y      int `json: y`
	Width  int `json: width`
	Height int `json: height`
}

type loc struct {
	Lat float32 `json: lat`
	Lon float32 `json: lon`
}

func client() {

	clientData := data{
		Main:  rectangle{2, 3, 4, 5},
		Input: []rectangle{{6, 7, 8, 9}, {1, 4, 7, 8}},
	}
	clientDataJson, err := json.Marshal(clientData)
	req, err := http.NewRequest("POST", "http://localhost:8088", bytes.NewBuffer(clientDataJson))
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
