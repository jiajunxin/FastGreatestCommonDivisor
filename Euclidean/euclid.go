package euclid

import (
	"math/big"
)

var zero = big.NewInt(0)
var two = big.NewInt(2)

// Gcd calculates the greatest common divisor of a and b using Euclidean algorithm
func Gcd(a, b *big.Int) big.Int {
	if a.Sign() == 0 {
		if b.Sign() == 0 {
			//gcd(0, 0) = 0
			return *zero
		}
		// gcd(0, b) = |b|
		var ret big.Int
		return *ret.Abs(b)
	}
	if b.Sign() == 0 {
		// gcd(a, 0) = |a|
		var ret big.Int
		return *ret.Abs(a)
	}

	for isModTwoZero(a) && !isZero(a) {
		a.Rsh(a, 1)
	}
	for isModTwoZero(b) && !isZero(b) {
		b.Rsh(b, 1)
	}

	var aTemp, bTemp big.Int
	aTemp.Abs(a)
	bTemp.Abs(b)
	if aTemp.Cmp(&bTemp) == -1 {
		return gcd(&bTemp, &aTemp)
	}
	return gcd(&aTemp, &bTemp)
}

// gcd(a, b) where a > b
func gcd(a, b *big.Int) big.Int {
	//fmt.Println("test here 0, a = ", a.String())
	//fmt.Println("b = ", b.String())
	if b.Sign() == 0 {
		return *a
	}
	var reminder big.Int
	reminder.Mod(a, b)
	if reminder.Sign() == 0 {
		//fmt.Println("qq b = ", b.String())
		return *b
	}
	return gcd(b, &reminder)
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
