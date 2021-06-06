// Package floats perform calculations on floating point numbers.
package floats

import "math"

// IsFinite returns whether f is a finite value.
func IsFinite(f float64) bool {
	if math.IsInf(f, 0) {
		return false
	}
	if math.IsNaN(f) {
		return false
	}
	return true
}
