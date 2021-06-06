// Package massconv calculates the weight of pounds and kilograms.
package massconv

// PoundToKilogram converts the weight of a pound to the weight of a kilogram.
func PoundToKilogram(p Pound) Kilogram { return Kilogram(float64(p) / PoundPerKilogram) }

// KilogramToPound converts the weight of a kilogram to the weight of a pound.
func KilogramToPound(k Kilogram) Pound { return Pound(float64(k) * PoundPerKilogram) }
