package models

import (
	"encoding/json"
	"time"
)

type UrlData struct {
	Tags       []string  `json:"tags"`
	Url        string    `json:"url"`
	DateString string    `json:"-"`
	Date       time.Time `json:"news_time"`
}

func (u UrlData) String() string {
	b, _ := json.Marshal(u)
	return string(b)
}
