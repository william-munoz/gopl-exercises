// ch04 / ex01 counts the number of different bits in the two SHA256 hashes.
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "ch03/ex12: must have 2 arguments.")
		os.Exit(1)
	}
	fmt.Printf("%d\n", sha256PopCount(os.Args[1], os.Args[2]))
}

// sha256PopCount returns the number of different bits in the SHA256 hash of the strings a, b.
func sha256PopCount(a, b string) int {
	digesta := sha256.Sum256([]byte(a))
	digestb := sha256.Sum256([]byte(b))
	return popCount(digesta, digestb)
}

// popCount returns the population count of x.
func popCount(a, b [32]byte) int {
	pop := 0
	for i := range a {
		pop += int(pc[a[i]^b[i]])
	}
	return pop
}
