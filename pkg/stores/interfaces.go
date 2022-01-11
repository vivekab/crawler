package stores

import "github.com/vivekab/crawler/pkg/models"

type News interface {
	// Insert news into the store and report the number of records inserted or an error
	Insert(source, status string, newsData []models.UrlData) (n int, err error)
	// Get the latest news from the fire store
	GetLatest(source, status string) (m *models.UrlData, err error)
	// Update method updates the given news
}
