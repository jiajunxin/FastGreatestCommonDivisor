package euclid

import (
	"math/big"
	"testing"
)

var (
	one   = big.NewInt(1)
	three = big.NewInt(3)
	five  = big.NewInt(5)
	six   = big.NewInt(6)
)

func TestGcd(t *testing.T) {
	var result big.Int
	var temp int64
	result = Gcd(six, three)
	if !result.IsInt64() {
		t.Errorf("gcd(6, 3) cannot be converted to Int64")
	}
	temp = result.Int64()
	if temp != 3 {
		t.Errorf("gcd(6, 3) not 3, temp = %d", temp)
	}

	result = Gcd(five, three)
	if !result.IsInt64() {
		t.Errorf("gcd(6, 3) cannot be converted to Int64")
	}
	temp = result.Int64()
	if temp != 1 {
		t.Errorf("gcd(6, 3) not 3, temp = %d", temp)
	}

	result = Gcd(five, six)
	if !result.IsInt64() {
		t.Errorf("gcd(6, 3) cannot be converted to Int64")
	}
	temp = result.Int64()
	if temp != 1 {
		t.Errorf("gcd(6, 3) not 3, temp = %d", temp)
	}
}
