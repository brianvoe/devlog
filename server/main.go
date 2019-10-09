package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/brianvoe/devlog"
)

func main() {
	go devlog.Server()

	rand.Seed(time.Now().UnixNano())
	for {
		request, err := json.Marshal(map[string]interface{}{
			"level": "info",
			"data": map[string]interface{}{
				"number": rand.Intn(100000000),
			},
		})
		if err != nil {
			log.Fatalln(err)
		}

		http.Post("http://localhost:8888/adddata", "application/json", bytes.NewBuffer(request))

		// Random sleep
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	}
}
