package main

import (
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

func appendFile(newData data) {
	filename := "file.json"
	// err := checkFile(filename)
	// if err != nil {
	//     log.Fatal(err)
	// }

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	clientData := []data{}

	// Here the magic happens!
	json.Unmarshal(file, &clientData)

	clientData = append(clientData, newData)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(clientData)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		fmt.Println("get")

		break
	case "POST":
		fmt.Println("post")
		newData := data{}
		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading the body", err)
		}
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(jsn, &newData)
		appendFile(newData)
		if err != nil {
			log.Fatal("Decoding error: ", err)
		}
		clientData := data{
			Main:  rectangle{2, 3, 4, 5},
			Input: []rectangle{{6, 7, 8, 9}, {1, 4, 7, 8}},
		}

		clientDataJson, err := json.Marshal(clientData)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")

		w.Write(clientDataJson)
	}
}

func server() {
	http.HandleFunc("/", fileHandler)
	http.ListenAndServe(":8088", nil)
}
func main() {
	server()
}
