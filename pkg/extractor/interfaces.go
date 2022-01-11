package extractor

import "time"

// TimeExtractor retrieves the time information from the given string
type TimeExtractor interface{
	// Extract function takes the input as a string and returns a time.Time and an error
	// if parsing goes fine it returns a valid time or returns the respective associated error
	Extract(input string) (time.Time,error)
}