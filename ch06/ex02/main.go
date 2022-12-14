// ch06/ex01 は、IntSet の AddAll 関数を実行します。
package main

import (
	"fmt"

	"github.com/williammunozr/gopl-exercises/ch06/ex02/intset"
)

func main() {
	is := &intset.IntSet{}

	is.Add(1)

	fmt.Println(is) // "{1}"

	is.AddAll(2, 3, 4)

	fmt.Println(is) // "{1 2 3 4}"
}
