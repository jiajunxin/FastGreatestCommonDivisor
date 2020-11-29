package euclid

import (
	"math/big"
	"math/rand"
	"testing"
)

var (
	three        = big.NewInt(3)
	five         = big.NewInt(5)
	six          = big.NewInt(6)
	randNumRange = big.NewInt(10000000)
)

const seed = 2020
const testSize = 1000

func CompareWithStandard(seed int64) bool {
	rng := rand.New(rand.NewSource(seed))
	var a, b, temp big.Int

	for i := 0; i < testSize; i++ {
		a.Rand(rng, randNumRange)
		b.Rand(rng, randNumRange)
		gcd1 := Gcd(&a, &b)
		temp.GCD(nil, nil, &a, &b)
		if gcd1.Cmp(&temp) != 0 {
			return false
		}
	}
	return true
}

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

	if !CompareWithStandard(seed) {
		t.Errorf("gcd result is different from standard lib")
	}
}
