// ex08 Displays what is in the URL.
// URL If does not have the prefix http://, add it before retrieving the content.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		completeURL := appendPrefix(url)
		resp, err := http.Get(completeURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", completeURL, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// appendPrefix Returns the string with prefix http:// if the given string does not have it.
// If it already has the prefix http://, it returns the given string as is.
func appendPrefix(str string) string {
	if !strings.HasPrefix(str, "http://") {
		return "http://" + str
	}
	return str
}
