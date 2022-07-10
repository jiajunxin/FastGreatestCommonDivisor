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

func main() {
	fmt.Println("Hello Greatest common divisor")
	// main function is used to test some large test numbers. Golang benchmark framework can run
	// limited time.
	testNum := 10000

	rng := rand.New(rand.NewSource(seed))
	var one = big.NewInt(1)
	var two = big.NewInt(2)
	var a, b, temp big.Int
	var gcd = big.NewInt(1)
	var bufInt = big.NewInt(1)

	start := time.Now()
	a.Rand(rng, randNumRange)
	temp.Mod(&a, two)
	if temp.Cmp(one) != 0 {
		a.Sub(&a, one)
	}

	for i := 0; i < testNum; i++ {
		printProgress(i, testNum)

		b.Rand(rng, randNumRange)
		// Make sure b is odd
		temp.Mod(&b, two)
		if temp.Cmp(one) != 0 {
			b.Sub(&b, one)
		}
		bufInt.GCD(nil, nil, &a, &b)
		if bufInt.Cmp(one) != 0 {
			fmt.Println("Found one factor = ", bufInt.String())
			gcd.Mul(gcd, bufInt)
			a.Div(&a, bufInt)
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Bit length of a = ", a.String())
	fmt.Println("Time elapsed = ", elapsed)
	fmt.Println("Bitlength of gcd = ", gcd.BitLen())
}
