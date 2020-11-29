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
	testNum := 100000000

	rng := rand.New(rand.NewSource(seed))
	var a, b big.Int
	var bufInt = big.NewInt(1)

	start := time.Now()
	for i := 0; i < testNum; i++ {
		printProgress(i, testNum)
		a.Rand(rng, randNumRange)
		b.Rand(rng, randNumRange)
		_ = bufInt.GCD(nil, nil, &a, &b)
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed = ", elapsed)
}
