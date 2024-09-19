package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func main() {
	var (
		bitInd int
		bitVal int
		number int64
	)
	flag.IntVar(&bitInd, "index", 0, "Index of the bit to set")
	flag.IntVar(&bitVal, "value", 1, "New value of the bit")
	flag.Int64Var(&number, "number", 1213, "The initial number")
	flag.Parse()

	if bitInd > 63 || bitInd < 0 {
		log.Fatal("The bit index is out of range [0..63]")
	}

	if bitVal != 1 && bitVal != 0 {
		log.Fatal("The bit value is out of range [0,1]")
	}

	fmt.Printf("The initial number is: %d, or in bit form:\n", number)
	numberStr := fmt.Sprintf("%064b", uint64(number))
	fmt.Println(numberStr)

	arrowString := strings.Repeat(" ", len(numberStr))
	arrowString = replaceAtIndex(arrowString, '^', len(numberStr)-1-bitInd)
	fmt.Println(arrowString)

	previousBit := number & (1 << bitInd)
	fmt.Printf("Bit mask: %b\n", uint64(previousBit))
	fmt.Printf("Bit add mask: %b\n", uint64(bitVal)<<uint64(bitInd))
	number = previousBit ^ number ^ (int64(bitVal) << int64(bitInd))
	fmt.Printf("Number after setting the bit is: %d, or in bit form:\n", number)
	fmt.Printf("%063b\n", uint64(number))
	fmt.Println(arrowString)
}
