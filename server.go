package devlog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"text/template"
	"time"
)

// GetHTML will get the index html file
func GetHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t, _ := template.ParseFiles("../index.html")
	t.Execute(w, nil)
}

// GetData will get the current data set and return it back
func GetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Json unmarshal to data
	d := struct {
		Levels []string `json:"levels"`
	}{}
	if err := json.Unmarshal(body, &d); err != nil {
		http.Error(w, "Error unmarshalling data", http.StatusInternalServerError)
		return
	}

	out := []Data{}
	DataMap.Range(func(key interface{}, value interface{}) bool {
		if contains(d.Levels, (value.(Data)).Level) {
			out = append([]Data{value.(Data)}, out...)
		}
		return true
	})

	// Sort data by created at
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].CreatedAt > out[j].CreatedAt
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(out)
}

// AddData will take a post request and add it to the data set
func AddData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Json unmarshal to data
	d := Data{}
	if err := json.Unmarshal(body, &d); err != nil {
		http.Error(w, "Error unmarshalling data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set ID/CreateAt
	d.ID = uuid()
	d.CreatedAt = time.Now().Nanosecond()

	// Store data
	DataMap.Store(d.ID, d)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success")
}

// Server is the function that starts an http server
func Server() {
	http.HandleFunc("/getdata", GetData)
	http.HandleFunc("/adddata", AddData)
	http.HandleFunc("/", GetHTML)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
