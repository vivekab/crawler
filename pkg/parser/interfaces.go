package parser

import (
	"github.com/gocolly/colly"

	"github.com/vivekab/crawler/pkg/models"
)

// News provides functionality to parse links from different websites
type News interface {
	// OverrideCollector overrides the default colly collector according to the need of the parser
	OverrideCollector(c *colly.Collector)

	// ParseHandler is a function that takes colly as an argument and returns an HTMLELEMENT and a callback that needs to be done on that
	ParseHandler() (htmlElement string, callback colly.HTMLCallback)

	Get() []models.UrlData
}
