package devlog

import "sync"

// Data is the primary struct that contains the main information
type Data struct {
	ID        string      `json:"id"`
	Level     string      `json:"level"`
	Data      interface{} `json:"data"`
	CreatedAt int64       `json:"created_at"`
}

// DataMap is the primary map where we store the data
var DataMap sync.Map
