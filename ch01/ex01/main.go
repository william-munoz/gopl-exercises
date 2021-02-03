// ex01 Displays the name of the command that invoked it and the command line arguments.
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
