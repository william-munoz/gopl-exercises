// Package lengthconv calculates the length of feet and meters.
package lengthconv

// FootToMeter converts foot lengths to metric lengths.
func FootToMeter(f Foot) Meter { return Meter(float64(f) / FootPerMeter) }

// MeterToFoot converts the length of a meter to the length of a foot.
func MeterToFoot(m Meter) Foot { return Foot(float64(m) * FootPerMeter) }
