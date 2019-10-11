package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/brianvoe/devlog"
	"github.com/brianvoe/gofakeit/v4"
)

func main() {
	go devlog.Server()

	rand.Seed(time.Now().UnixNano())
	for {
		request, err := json.Marshal(struct {
			Level string      `json:"level"`
			Data  interface{} `json:"data"`
		}{
			Level: gofakeit.RandString([]string{"info", "debug", "warn", "error"}),
			Data:  gofakeit.Map(),
		})
		if err != nil {
			log.Fatalln(err)
		}

		client := http.Client{Timeout: time.Second * 2}
		resp, err := client.Post("http://localhost:8888/adddata", "application/json", bytes.NewBuffer(request))
		if err != nil {
			panic(err)
		}
		resp.Body.Close()

		// Random sleep
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	}
}
