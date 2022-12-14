// ch13 / ex03 is a bzip compression tool that utilizes the C library.
package main

import (
	"io"
	"log"
	"os"

	"github.com/williammunozr/gopl-exercises/ch13/ex03/bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}
