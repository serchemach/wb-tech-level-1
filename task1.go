package main

import "fmt"

type Human struct {
	name string
}

func (man *Human) say() {
	fmt.Printf("My name is: %s\n", man.name)
}

// В го нет наследования, но есть композиция структур (что на самом
// деле во многом тоже самое).
type Action struct {
	Human
}

func main() {
	h := Human{
		name: "human",
	}
	h.say()

	a := Action{
		Human{name: "action"},
	}
	a.say()
}
