// Package complexrat uses big.Rat to represent complex numbers.
package complexrat

import "math/big"

// ComplexRat uses big.Rat to represent complex numbers.
type ComplexRat struct {
	Re *big.Rat
	Im *big.Rat
}

// AbsCompare compares the absolute value of a complex number with n.
// Returns a positive value when the absolute value of the complex is greater.
func (c *ComplexRat) AbsCompare(n *big.Rat) int {
	re2 := big.NewRat(0, 1).Mul(c.Re, c.Re)
	im2 := big.NewRat(0, 1).Mul(c.Im, c.Im)
	n2 := big.NewRat(0, 1).Mul(n, n)
	return big.NewRat(0, 1).Add(re2, im2).Cmp(n2)
}

// Square returns the square of the complex number c.
func (c *ComplexRat) Square() *ComplexRat {
	re2 := big.NewRat(0, 1).Mul(c.Re, c.Re)
	im2 := big.NewRat(0, 1).Mul(c.Im, c.Im)
	reim := big.NewRat(0, 1).Mul(c.Re, c.Im)

	c.Re.Sub(re2, im2)
	c.Im.Add(reim, reim)
	return c
}

// Add returns the sum of complex numbers a and b.
func (c *ComplexRat) Add(z *ComplexRat) *ComplexRat {
	c.Re.Add(c.Re, z.Re)
	c.Im.Add(c.Im, z.Im)
	return c
}
