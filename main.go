package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func printProgress(icounter, size int) {
	tenPercent := size / 10
	if icounter == tenPercent*1 {
		fmt.Println("10% finished")
	}
	if icounter == tenPercent*3 {
		fmt.Println("30% finished")
	}
	if icounter == tenPercent*5 {
		fmt.Println("50% finished")
	}
	if icounter == tenPercent*7 {
		fmt.Println("70% finished")
	}
	if icounter == tenPercent*9 {
		fmt.Println("90% finished")
	}
}

func GetRandomOdd(rnd *rand.Rand, n *big.Int) *big.Int {
	var ret, temp big.Int
	ret.Rand(rnd, n)
	temp.Mod(&ret, two)
	if temp.Cmp(one) != 0 {
		ret.Sub(&ret, one)
	}
	return &ret
}

// ExploreNum explores if the input number has small factors starting from 3
func ExploreNum(input *big.Int) {
	fmt.Println("Begin to explore factors of ", input.String())
	var temp big.Int
	var copyInput big.Int
	copyInput.Set(input)
	starter := big.NewInt(3)

	for i := 0; i < 10000; i++ {
		temp.GCD(nil, nil, &copyInput, starter)
		if temp.Cmp(one) != 0 {
			fmt.Println("Found one small factor, gcd = ", temp.String(), ", factor = ", starter.String())
		}
		copyInput.Div(&copyInput, &temp)
		starter.Add(starter, two)
	}
	fmt.Println("End of exploring factors.")
}

func main() {
	fmt.Println("Hello Greatest common divisor")
	// main function is used to test some large test numbers. Golang benchmark framework can run
	// limited time.
	testNum := 10000000

	rng := rand.New(rand.NewSource(222))
	var a, b big.Int
	var gcd = big.NewInt(1)
	var bufInt = big.NewInt(1)

	start := time.Now()
	a = *GetRandomOdd(rng, randNumRange)
	ExploreNum(&a)
	for i := 0; i < testNum; i++ {
		printProgress(i, testNum)

		b = *GetRandomOdd(rng, randNumRange)
		bufInt.GCD(nil, nil, &a, &b)
		if bufInt.Cmp(one) != 0 {
			fmt.Println("Found one factor = ", bufInt.String())
			gcd.Mul(gcd, bufInt)
			a.Div(&a, bufInt)
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Bit length of a = ", a.BitLen())
	ExploreNum(&a)
	fmt.Println("Time elapsed = ", elapsed)
	fmt.Println("Bitlength of gcd = ", gcd.BitLen())
}
