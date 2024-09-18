package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
)

func main() {
	var (
		aStr string
		bStr string
	)
	flag.StringVar(&aStr, "num1", "100000000000000000000", "The first argument")
	flag.StringVar(&bStr, "num2", "50000800000", "The second argument")
	flag.Parse()

	a, b := big.NewInt(0), big.NewInt(0)
	a, ok := a.SetString(aStr, 10)
	if !ok {
		log.Fatal("The first number is not a number")
	}

	b, ok = b.SetString(bStr, 10)
	if !ok {
		log.Fatal("The second number is not a number")
	}

	sum := big.NewInt(0)
	sum.Add(a, b)
	fmt.Printf("The sum of the two numbers is:\n%s\n", sum.String())

	prod := big.NewInt(0)
	prod.Mul(a, b)
	fmt.Printf("The product of the two numbers is:\n%s\n", prod.String())

	sub := big.NewInt(0)
	sub.Sub(a, b)
	fmt.Printf("The difference num1-num2 is:\n%s\n", sub.String())
}
