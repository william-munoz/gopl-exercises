// Package massconv calculates the weight of pounds and kilograms.
package massconv

import "fmt"

// Pound represents the weight of a pound.
type Pound float64

// Kilogram represents the weight of a kilogram.
type Kilogram float64

const (
	// PoundPerKilogram is the value of pounds per kilogram.
	PoundPerKilogram float64 = 0.45359237
)

func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
