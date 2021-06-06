// Exercise 1.1: Modify the echo program to also print os.Ars[0], the name of the command that invoked it.
// Displays the name of the command that invoked it and the command line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func joinWithSpace(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	fmt.Println(joinWithSpace(os.Args))
}
