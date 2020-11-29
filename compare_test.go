package main

import (
	"math/big"
	"testing"

	euclid "github.com/FastGreatestCommonDivisor/Euclidean"
)

// This file tests the benchmark for different algorithms generating
// greatest common divisors.

func BenchmarkEuclidean(b *testing.B) {
	testSet := GenRandomInt()
	b.ResetTimer()

	for i := 0; i < testSize-1; i++ {
		_ = euclid.Gcd(&testSet[i], &testSet[i+1])
	}

}

func BenchmarkGolangStandard(b *testing.B) {
	testSet := GenRandomInt()
	var bufInt = big.NewInt(1)
	b.ResetTimer()

	for i := 0; i < testSize-1; i++ {
		_ = bufInt.GCD(nil, nil, &testSet[i], &testSet[i+1])
	}

}
