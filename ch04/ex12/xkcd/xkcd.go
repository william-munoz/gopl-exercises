// Package xkcd provides Go's API for xkcd.
package xkcd

import (
	"fmt"
	"strconv"
)

func getComicURL(comicID int) string {
	return fmt.Sprintf("https://xkcd.com/%s/info.0.json", strconv.Itoa(comicID))
}

// ComicIndex represents the index of xkcd comics.
type ComicIndex struct {
	Comics []*Comic
}

// NewComicIndex returns a new xkcd comic index.
func NewComicIndex() ComicIndex {
	return ComicIndex{[]*Comic{}}
}

// Comic stands for xkcd comics.
type Comic struct {
	Alt        string
	Day        string
	Img        string
	Link       string
	Month      string
	News       string
	Num        int
	SafeTitle  string
	Title      string
	Transcript string
	Year       string
}
