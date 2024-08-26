package prettyprintradians

import (
	"math"
)

// for a first pass, lets just have a couple of public function
// PrettyPrint(angle float64) string
// and the worry about any configuration later.

const pi = "π"

func PrettyPrint(angle float64) string {

	angle = angle / math.Pi
	f := fromDecimal(angle)
	return f.withPi()
}

func fromDecimal(v float64) Fraction {
	accuracy := 0.001

	sgn := sign(v)
	if sgn == -1 {
		v = -1.0 * v
	}

	var maxError float64
	if sgn == 0 {
		maxError = accuracy
	} else {
		maxError = v * accuracy
	}

	f := math.Floor(v)
	n := int(f)
	v -= f

	if v < maxError {
		return Fraction{sgn * n, 1}
	}

	if 1-maxError < v {
		return Fraction{sgn * (n + 1), 1}
	}

	// lower fraction is 0/1
	ln := 0
	ld := 1

	// upper fraction is 1/1
	un := 1
	ud := 1

	for {
		mn := ln + un
		md := ld + ud

		if float64(md)*(v+maxError) < float64(mn) {
			un = mn
			ud = md
		} else if float64(mn) < (v-maxError)*float64(md) {
			ln = mn
			ld = md
		} else {
			return Fraction{
				(n*md + mn) * sgn,
				md,
			}
		}
	}

}

func sign(v float64) int {
	if v > 0 {
		return 1
	} else if v < 0 {
		return -1
	} else {
		return 0
	}
}
