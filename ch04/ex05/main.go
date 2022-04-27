// ch04 / ex05 removes adjacent duplicates from within the slice.
package main

import (
	"fmt"
)

func main() {
	s := []string{"A", "A", "B", "B", "A", "C"}
	fmt.Println(removeDup(s)) // "[A B A C]"
}

// removeDup removes adjacent duplicates from within the slice.
func removeDup(s []string) []string {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}
