// Exercise 1.3: Experimient to measure the difference in running time between our potentially
// inefficient verions and the one that uses strings.Join. (Section 1.6 illustrates part of the
// time package, and Section 11.4 shows how to write benchmark tests for systematic performance
// evaluation.)
// Package ex03 Displays line arguments.
package ex03

import (
	"os"
	"strings"
)

// Echo1 Uses a for loop to display command line arguments.
func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// fmt.Println(s)
}

// Echo2 Displays command line arguments using a for loop and range.
func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

// Echo3 Displays command line arguments using strings.Join.
func Echo3() {
	strings.Join(os.Args[1:], " ")
	// fmt.Println(strings.Join(os.Args[1:], " "))
}
