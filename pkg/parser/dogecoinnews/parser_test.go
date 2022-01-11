package dogecoinnews

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"

	dogenewstimeextractor "github.com/vivekab/crawler/pkg/extractor/dogeNewsTimeExtractor"
)

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		b, _ := ioutil.ReadFile("success-1.html")
		w.Write(b)
		w.WriteHeader(200)
	})

	return httptest.NewServer(mux)
}

func Test_dogeNewsPageParser_fetchNews(t *testing.T) {

	ts := newTestServer()
	defer ts.Close()

	cR := colly.NewCollector()
	// Define the dogeNewsLogger to log the details extracted
	dogeNewsLogger := logrus.New().WithField("Source", "dogeNews").Logger
	// dogeNewsLogger.Formatter = new(logrus.JSONFormatter)
	dogeNewsLogger.Level = logrus.InfoLevel
	dogeNewsLogger.Out = os.Stdout

	cR.MaxDepth = 2
	cR.Async = false

	// Create the timeExtractor and inject it into the parser and then run the parser
	fromTime := time.Date(2021, time.May, 26, 11, 42, 0, 0, time.UTC)
	dgParser := New(&fromTime, dogenewstimeextractor.New(), dogeNewsLogger)
	dgParser.OverrideCollector(cR)
	cR.OnHTML(dgParser.ParseHandler())

	if err := cR.Visit(ts.URL + "/"); err != nil {
		t.Error(err)
	}
	cR.Wait()
	parserInfo := dgParser.Get()
	if len(parserInfo) != 1 {
		t.Error("Error parsing data", parserInfo)
	}
}
