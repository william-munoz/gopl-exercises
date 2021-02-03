// ex04 Displays the number of lines that appear more than once in a standard input or specified file and their text.
// If a file is specified, all filenames that contained each of the duplicate lines will be displayed.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	occurrences := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, occurrences)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occurrences)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, occurrences[line])
		}
	}
}

// countLines Reads the input line by line and records the number of lines that occur and the name of the input.
func countLines(f *os.File, counts map[string]int, occurrences map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		occurrences[input.Text()] = appendAsSet(occurrences[input.Text()], f.Name())
	}
}

// appendAsSet Adds the given string str to the given string array set.
// However, if set already contains str, do not add str.
func appendAsSet(set []string, str string) []string {
	if !includes(set, str) {
		return append(set, str)
	}
	return set
}

// includes Returns whether the array of the given strings contains the given string str.
func includes(array []string, str string) bool {
	for _, value := range array {
		if value == str {
			return true
		}
	}
	return false
}
