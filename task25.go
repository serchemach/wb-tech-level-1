package main

import (
	"fmt"
	"os/exec"
	"time"
)

// Это не совсем честно, поэтому давайте ещё сделаем тупую версию
func sysSleep(seconds float64) error {
	cmd := exec.Command("sleep", fmt.Sprintf("%f", seconds))
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func dumbSleep(seconds float64) error {
	dur, err := time.ParseDuration(fmt.Sprintf("%fs", seconds))
	if err != nil {
		return err
	}
	endTime := time.Now().Add(dur)

	// fmt.Println(time.Now(), endTime)
	for time.Now().Before(endTime) {
	}
	return nil
}

func main() {
	fmt.Printf("Let's syssleep for 0.5 seconds!\n")
	err := sysSleep(0.5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sysslept for 0.5 seconds!\n")
	}

	fmt.Printf("Let's dumbsleep for 0.7 seconds!\n")
	err = dumbSleep(0.7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Dumbslept for 0.7 seconds!\n")
	}

}
