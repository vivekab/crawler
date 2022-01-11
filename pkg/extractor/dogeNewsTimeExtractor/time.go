package dogenewstimeextractor

import (
	"strings"
	"time"

	"github.com/vivekab/crawler/pkg/extractor"
)

// dogeNewsTime implements the functionality of converting the given string information into time
type dogeNewsTime struct{}

const dogeNewsFormat = "January 2, 2006 at 15:04 pm"

// New is the factory function to produce a dogeNewsTime implementor
func New() extractor.TimeExtractor {
	return &dogeNewsTime{}
}

// Extract function takes a given string and extracts the time information out of it if exist else returns an error
// Example: time format by David Cox on `May 28, 2021 at 11:42 am` is converted to 2021-05-28T11:42:00
func (d *dogeNewsTime) Extract(input string) (t time.Time, err error) {
	strs := strings.Split(input, "on")
	st := strings.TrimLeft(strs[len(strs)-1], " ")
	st = strings.TrimRight(st, " ")
	return time.Parse(dogeNewsFormat, st)

}
