package prettyprintradians

import "fmt"

type Fraction struct {
	numerator   int
	denominator int
}

func newFraction(n, d int) Fraction {
	return Fraction{n, d}
}

func (f Fraction) String() string {
	if f.denominator != 1 {
		return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
	}
	return fmt.Sprintf("%d", f.numerator)
}

func (f Fraction) withPi() string {
	if f.denominator != 1 {
		return fmt.Sprintf("%d%s/%d", f.numerator, pi, f.denominator)
	}
	return fmt.Sprintf("%d%s", f.numerator, pi)
}
