package dogecoinnews

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"

	"github.com/vivekab/crawler/pkg/extractor"
	"github.com/vivekab/crawler/pkg/models"
	"github.com/vivekab/crawler/pkg/parser"
)

// dogeNewsPageParser gets the latest news on the website identified by https://www.dogecoinnews.net/
// stores inside ParsedNews
type dogeNewsPageParser struct {
	from          *time.Time
	ParsedNews    []models.UrlData
	timeExtractor extractor.TimeExtractor
	logger        *logrus.Logger
}

// New is the Factory function to create a dogeNewsPageParser
func New(from *time.Time, tE extractor.TimeExtractor, infoLog *logrus.Logger) parser.News {
	return &dogeNewsPageParser{
		from:          from,
		ParsedNews:    []models.UrlData{},
		timeExtractor: tE,
		logger:        infoLog,
	}
}

// Parsehandler returns an html element that needs to be parsed with a callback
func (d *dogeNewsPageParser) ParseHandler() (htmlElement string, callback colly.HTMLCallback) {
	return "li", d.fetchNews
}

// fetchNews is a function where the logic of fetching the individual news resides
func (d *dogeNewsPageParser) fetchNews(h *colly.HTMLElement) {
	data := models.UrlData{
		Url:        h.ChildAttr("a", "href"),
		DateString: h.DOM.Find("div").Find("small").Text(),
	}

	if data.DateString == "" {
		d.logger.Info("No date string found ", data, h)
		return
	}

	t, err := d.timeExtractor.Extract(data.DateString)
	if err != nil {
		d.logger.Errorln("Error Extracting dogeNews time ", data.DateString, " from date string", err)
		return
	}

	if t.Before(*d.from) || t.Equal(*d.from) {
		return
	}

	data.Date = t

	d.ParsedNews = append(d.ParsedNews, data)

	d.logger.Info("Extracted ", data)
}

// OverrideCollector override the colly collector and applies the specific constraint on the colly Collector
func (d *dogeNewsPageParser) OverrideCollector(c *colly.Collector) {
	c.AllowURLRevisit = false
	c.MaxDepth = 2

	c.DisallowedDomains = []string{"mailto:contact@dogecoinnews.net"}
}

func (d *dogeNewsPageParser) Get() []models.UrlData {
	return d.ParsedNews
}
