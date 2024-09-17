package main

import (
	"flag"
	"fmt"
)

func main() {
	var numberOfWorkers int
	flag.IntVar(&numberOfWorkers, "n", 10, "The number of workers")
	flag.Parse()

	fmt.Println(numberOfWorkers)
}
