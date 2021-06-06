// ch03 / ex11 inserts a comma in a floating point string with a sign.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", commaSigned(os.Args[i]))
	}
}

// commaSigned inserts a comma in a floating-point number string with a sign.
func commaSigned(s string) string {
	var start, end int

	if strings.HasPrefix(s, "-") {
		start = 1
	} else {
		start = 0
	}

	if strings.Contains(s, ".") {
		end = strings.Index(s, ".")
	} else {
		end = len(s)
	}

	return s[:start] + comma(s[start:end]) + s[end:]
}

// comma inserts a comma in the string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
