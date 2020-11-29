package fkm

import "math/big"

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)

// Gcd implements the a fast hybrid GCD computation algorithm by Faraoun Kamel Mohamed
// in the paper "A novel fast hybrid GCD computation algorithm", Int. J. Computing and
// Science and Mathematics, Vol. 5, No. 1, 2014
func Gcd(u, v *big.Int) big.Int {
	if u.Sign() == 0 {
		if v.Sign() == 0 {
			//gcd(0, 0) = 0
			return *zero
		}
		// gcd(0, v) = |v|
		var ret big.Int
		return *ret.Abs(v)
	}
	g := big.NewInt(1)

	for isModTwoZero(u) && isModTwoZero(v) {
		u.Rsh(u, 1)
		u.Rsh(v, 1)
		g.Add(g, one)
	}
	for isModTwoZero(u) && !isZero(u) {
		u.Rsh(u, 1)
	}
	for isModTwoZero(v) && !isZero(v) {
		v.Rsh(v, 1)
	}
	if u.Cmp(v) < 0 { // if u < v, exchange (u, v)
		exchange(u, v)
	}

	for v.Cmp(one) > 0 {
		u.Sub(u, v)
		for isModTwoZero(u) && !isZero(u) {
			u.Rsh(u, 1)
		}
	}
}

// exchange the input value
func exchange(u, v *big.Int) {
	var temp big.Int
	temp.Set(u)
	u.Set(v)
	v.Set(&temp)
}

// isZero returns true if input = 0
func isZero(input *big.Int) bool {
	if input.Cmp(zero) == 0 {
		return true
	}
	return false
}

// isModTwoZero returns true if input mod 2 = 0
func isModTwoZero(input *big.Int) bool {
	z := big.NewInt(1)
	z.Mod(input, two)
	if isZero(z) {
		return true
	}
	return false
}
