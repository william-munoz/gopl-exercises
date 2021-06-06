// ch03 / ex10 inserts a comma in the string without making a recursive call.
package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

// comma inserts a comma into the string without making a recursive call.
func comma(s string) string {
	var buf bytes.Buffer
	i := (3 - utf8.RuneCountInString(s)%3) % 3
	for _, r := range s {
		if i == 3 {
			buf.WriteByte(',')
			i = 0
		}
		buf.WriteRune(r)
		i++
	}
	return buf.String()
}
