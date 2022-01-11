package main

import (
	"context"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	dogenewstimeextractor "github.com/vivekab/crawler/pkg/extractor/dogeNewsTimeExtractor"
	"github.com/vivekab/crawler/pkg/parser/dogecoinnews"
	firebaseStore "github.com/vivekab/crawler/pkg/stores/firebase"
)

const (
	dogenewsSource = "dogecoinnewsnet"
	loaded         = "LOADED"
	processed      = "PROCESSED"
)

func main() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		fetchDogeNews()
	}

}

func fetchDogeNews() {
	c := colly.NewCollector()

	// Define the dogeNewsLogger to log the details extracted
	dogeNewsLogger := logrus.WithField("Source", dogenewsSource).Logger
	// dogeNewsLogger.Formatter = new(logrus.JSONFormatter)
	dogeNewsLogger.Level = logrus.InfoLevel
	dogeNewsLogger.Out = os.Stdout

	// Initialize the store
	ctx := context.Background()
	config := &firebase.Config{}
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEY_PATH"))
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.DatabaseWithURL(ctx, os.Getenv("FIREBASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	lastFetchedNewsTime := time.Time{}

	// create the local firebase store
	store := firebaseStore.New(client, dogeNewsLogger)

	// Get the last fetched latest news
	lastFetchedNews, err := store.GetLatest(dogenewsSource, loaded)
	if err != nil {
		dogeNewsLogger.Info("Error finding records in ", loaded)
		lastFetchedNews, err = store.GetLatest(dogenewsSource, processed)
		if err != nil {
			dogeNewsLogger.Info("Error finding records in ", loaded, " AND ", processed)
		}
	}

	if lastFetchedNews != nil {
		lastFetchedNewsTime = lastFetchedNews.Date
	} else {
		// if lastfetchedNews is not obtained fetch the news from the last two days
		lastFetchedNewsTime = time.Now().UTC().Add(-2 * 24 * time.Hour)
	}

	dogeNewsLogger.Info("Extracting news after ", lastFetchedNewsTime.Format(time.RFC3339Nano))

	// Create the timeExtractor and inject it into the parser and then run the parser
	dT := dogenewstimeextractor.New()
	dgParser := dogecoinnews.New(&lastFetchedNewsTime, dT, dogeNewsLogger)
	dgParser.OverrideCollector(c)

	c.OnHTML(dgParser.ParseHandler())
	c.Visit("https://www.dogecoinnews.net/")
	c.Wait()

	data := dgParser.Get()

	// Apppend data to firebase store
	n, err := store.Insert(dogenewsSource, loaded, data)
	if err != nil {
		dogeNewsLogger.Error("Error Inserting data to firebase Store", err)
	} else {
		dogeNewsLogger.Info("Added ", n, " new news to the database")
	}
}
