package models

import "time"

type StoreData struct {
	Tags   []string  `json:"tags"`
	Url    string    `json:"url"`
	Date   time.Time `json:"news_time"`
	Source string    `json:"source"`
	Status string    `json:"status"`
}

func NewStoreDataFromUrlData(u UrlData, source, status string) StoreData {
	return StoreData{
		Tags:   u.Tags,
		Url:    u.Url,
		Date:   u.Date,
		Source: source,
		Status: status,
	}
}
