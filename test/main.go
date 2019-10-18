package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/brianvoe/devlog"
	"github.com/brianvoe/gofakeit/v4"
)

func main() {
	const port string = "8888"
	go devlog.Server(port)

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
		resp, err := client.Post("http://localhost:"+port+"/adddata", "application/json", bytes.NewBuffer(request))
		if err != nil {
			fmt.Println(err)
		}
		if resp.Body != nil {
			resp.Body.Close()
		}

		// Random sleep
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	}
}
