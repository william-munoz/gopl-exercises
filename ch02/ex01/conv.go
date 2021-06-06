// Package tempconv performs temperature calculations in Celsius, Fahrenheit, and absolute temperatures.
package tempconv

// CToF converts the temperature in Celsius to the temperature in Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts the temperature in Fahrenheit to the temperature in Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts the temperature in Celsius to absolute temperature.
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToC converts absolute temperature to temperature in degrees Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

// FToK converts Fahrenheit temperature to absolute temperature.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToF converts absolute temperature to Fahrenheit temperature.
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
