package firebase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"firebase.google.com/go/v4/db"
	"github.com/sirupsen/logrus"

	"github.com/vivekab/crawler/pkg/models"
	"github.com/vivekab/crawler/pkg/stores"
)

type store struct {
	logger *logrus.Logger
	db     *db.Client
}

const (
	newsCollection = "news_data"
)

var (
	ErrNoRecords = errors.New("no records")
)

// New is the factory function to create the firebase store that implements the store interface
func New(con *db.Client, logger *logrus.Logger) stores.News {
	return &store{
		logger: logger,
		db:     con,
	}
}

// Insert function inserts the given set of news according to the given source and status
func (s *store) Insert(source, status string, newsData []models.UrlData) (n int, err error) {
	if len(newsData) <= 0 {
		s.logger.Error(ErrNoRecords)
		return 0, ErrNoRecords
	}

	newsStoreData := make([]models.StoreData, 0, len(newsData))
	for i := range newsData {
		newsStoreData = append(newsStoreData, models.NewStoreDataFromUrlData(newsData[i], source, status))
	}

	for i := range newsStoreData {
		_, err = s.db.NewRef(newsCollection).
			Push(context.Background(), newsStoreData[i])
		if err != nil {
			s.logger.Error("Error inserting news_data", err)
			continue
		}
		n++
	}

	return n, nil
}

// GetLatest function retrieves the latest news for the given source and status
func (s *store) GetLatest(source, status string) (m *models.UrlData, err error) {
	ms := make([]models.UrlData, 0, 2)

	if source == "" || status == "" {
		s.logger.Error(ErrNoRecords)
		return nil, ErrNoRecords
	}

	ref := s.db.NewRef("news_data")

	results, err := ref.OrderByChild("news_time").
		LimitToLast(1).
		GetOrdered(context.Background())
	if err != nil {
		s.logger.Error("Error fetching information ", err)
		return nil, err
	}

	for _, r := range results {
		score := models.UrlData{}
		if err := r.Unmarshal(&score); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}
		if score.Url == "" {
			continue
		} else {
			ms = append(ms, score)
		}
	}

	fmt.Println(ms)

	return &ms[0], nil
}

func (s *store) Update() {
	return
}
