package main

import (
	"encoding/json"
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

func isCovered(main rectangle, newRect rectangle) bool {

	var main_min_x int = main.X
	var main_max_x int = main.X + main.Width
	var main_min_y int = main.Y
	var main_max_y int = main.Y + main.Height

	var rect_min_x int = newRect.X
	var rect_max_x int = newRect.X + newRect.Width
	var rect_min_y int = newRect.Y
	var rect_max_y int = newRect.Y + newRect.Height

	if (main_min_x <= rect_max_x && main_min_x >= rect_min_x) || (main_max_x <= rect_max_x && main_max_x >= rect_min_x) {
		if (main_min_y <= rect_max_y && main_min_y >= rect_min_y) || (main_max_y <= rect_max_y && main_max_y >= rect_min_y) {
			return true
		}
	}

	if (main_min_x <= rect_max_x && main_max_x >= rect_max_x) || (main_max_x >= rect_min_x && main_min_x <= rect_min_x) {
		if (main_min_y <= rect_max_y && main_max_y >= rect_max_y) || (main_max_y >= rect_min_y && main_min_y <= rect_min_y) {
			return true
		}
	}
	return false
}

func appendFile(newData data) {
	filename := "file.json"

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	clientData := []rectangle{}

	json.Unmarshal(file, &clientData)

	for i := 0; i < len(newData.Input); i++ {
		if isCovered(newData.Main, newData.Input[i]) {
			newData.Input[i].CreationTime = time.Now().Local()
			clientData = append(clientData, newData.Input[i])
		}
	}

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
		filename := "file.json"

		file, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")

		w.Write(file)

		break

	case "POST":
		newData := data{}
		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading the body", err)
		}
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(jsn, &newData)
		if err != nil {
			log.Fatal("Decoding error: ", err)
		}
		appendFile(newData)
	}
}

func server() {
	http.HandleFunc("/", fileHandler)
	http.ListenAndServe(":8088", nil)
}
func main() {
	server()
}
