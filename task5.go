package main

import "time"

func main() {
	var timeToLive int
	flag.IntVar(&timeToLive, "n", 10, "The lifetime of the program in seconds")
	flag.Parse()

	go func(){
		time.Sleep(timeToLive * time.Second)
	}

}
