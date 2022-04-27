// ch04 / ex13 gets the movie poster image from the Open Movie Database.
package main

import (
	"fmt"
	"os"

	"github.com/williammunozr/gopl-exercises/ch04/ex13/omdb"
)

func main() {
	err := omdb.GetPoster(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/ex12: %v\n", err)
		os.Exit(1)
	}
}
