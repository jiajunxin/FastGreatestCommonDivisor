package main

import (
	"math/big"
	"math/rand"
)

const testSize = 1000000
const seed = 2020
const bitLength = 2000

var one = big.NewInt(1)
var randNumRange = big.NewInt(0)

func init() {
	_ = randNumRange.Lsh(one, bitLength)
	_ = randNumRange.Sub(randNumRange, one)
}

// GenRandomInt returns a slice of random integers with the set size of testSize and
// the random number range as bitLength
func GenRandomInt() []big.Int {
	result := make([]big.Int, testSize)
	rng := rand.New(rand.NewSource(seed))
	var randNum big.Int

	for i := range result {
		randNum.Rand(rng, randNumRange)
		result[i].Set(&randNum)
	}
	return result
}

// GenRandomSet returns a slice of random integers with the set size of setSize and
// the random number range as bitLength
func GenRandomSet(setSize int64) []big.Int {
	result := make([]big.Int, setSize)
	rng := rand.New(rand.NewSource(seed))
	var randNum big.Int

	for i := range result {
		randNum.Rand(rng, randNumRange)
		result[i].Set(&randNum)
	}
	return result
}
