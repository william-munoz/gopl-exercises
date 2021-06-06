// Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. Compare the performance
// of the two versions. (Section 11.4 shows how to compare the performance of different implementations
// systematically.)

// Package popcount returns the population count.
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// TablePopCount returns the population count of x using a table reference.
func TablePopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// LoopPopCount uses a loop to return the population count of x.
func LoopPopCount(x uint64) int {
	count := 0
	for i := uint64(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
