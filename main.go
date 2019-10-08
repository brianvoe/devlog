package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"text/template"
	"time"
)

// Data is the primary struct that contains the main information
type Data struct {
	Level string `json:"level"`
	Data  string `json:"data"`
}

var DataMap sync.Map

// GetHTML will get the index html file
func GetHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

// GetData will get the current data set and return it back
func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	out := []interface{}{}
	DataMap.Range(func(key, value interface{}) bool {
		out = append([]interface{}{value}, out...)
		return true
	})

	json.NewEncoder(w).Encode(out)
}

// AddData will take a post request and add it to the data set
func AddData(w http.ResponseWriter, r *http.Request) {
	DataMap.Store(time.Now().Unix(), Data{
		Level: "info",
		Data:  "things",
	})

	fmt.Fprintf(w, "Added")
}

func main() {
	http.HandleFunc("/html", GetHTML)
	http.HandleFunc("/getdata", GetData)
	http.HandleFunc("/adddata", AddData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
