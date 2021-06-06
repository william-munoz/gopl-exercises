// Package lengthconv calculates the length of feet and meters.
package lengthconv

import "fmt"

// Foot represents the length of the foot.
type Foot float64

// Meter represents the length of the meter.
type Meter float64

const (
	// FootPerMeter is the value of feet per meter.
	FootPerMeter float64 = 0.3048
)

func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
