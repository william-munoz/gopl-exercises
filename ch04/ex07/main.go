// ch04 / ex07 reverses the UTF-8 encoded strings.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Lorem ipsum"
	fmt.Println(string(reverseUTF8([]byte(s)))) // "muspi meroL"
}

// reverseUTF8 reverses the UTF-8 encoded string.
func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
	return b
}

func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
