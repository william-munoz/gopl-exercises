// Exercise 1.10: Find a web site that produces a large amount of data. Investigate caching by
// running fetchall twice in succession to see whether the reported time changes much. Do
// you get the same content each time? Modify fetchall to print its output to a file so it can be
// examined.
// Will retrieve the URLs in parallel and display the time and size twice.
// It also saves the contents obtained by the extraction to a file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fetchToDir(os.Args[1:], "out/1")
	fetchToDir(os.Args[1:], "out/2")
}

// fetchToDir Extracts URLs in parallel to display the time.
// It also saves the contents obtained by the retrieval in the specified directory.
func fetchToDir(urls []string, dirName string) {
	start := time.Now()
	ch := make(chan string)
	for idx, url := range urls {
		file, err := os.Create(fmt.Sprintf("%s/%d.txt", dirName, idx))
		if err != nil {
			fmt.Println(err) // send to channel ch
			return
		}
		defer file.Close()
		go fetch(url, file, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch Extracts the URL and displays the time and size.
// It also writes the content obtained by the retrieval to the given writer.
func fetch(url string, writer io.Writer, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(writer, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
