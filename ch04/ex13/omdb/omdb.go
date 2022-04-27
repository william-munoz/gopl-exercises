// Package omdb provides Go's API for Open Movie Database.
package omdb

import (
	"fmt"
	"net/url"
	"strings"
)

func searchURL(terms []string) string {
	return fmt.Sprintf("http://www.omdbapi.com/?t=%s", url.QueryEscape(strings.Join(terms, " ")))
}

// Movie represents the movie information in the Open Movie Database.
type Movie struct {
	Poster   string
	Response string
}
