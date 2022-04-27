// ch07/ex13 は、文字列を式としてパースした構文ツリーを、式のような文字列で表示します。
package main

import (
	"fmt"
	"log"

	"github.com/williammunozr/gopl-exercises/ch07/ex13/eval"
)

func main() {
	expr, err := eval.Parse("pow(x, 3) + pow(y, 3)")
	if err != nil {
		log.Fatalf("ch07/ex13: %v", err)
	}
	fmt.Printf("%s\n", expr)
}
