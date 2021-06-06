// Package tempconv performs temperature calculations in Celsius, Fahrenheit, and absolute temperatures.
package tempconv

import "fmt"

// Celsius represents the temperature in Celsius.
type Celsius float64

// Fahrenheit represents the temperature in Fahrenheit.
type Fahrenheit float64

// Kelvin represents the absolute temperature.
type Kelvin float64

const (
	// AbsoluteZeroC is absolute zero.
	AbsoluteZeroC Celsius = -273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
