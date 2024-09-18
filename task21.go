package main

import "fmt"

// Давайте адаптируем наш тип к интерфейсу error
type Adaptee struct {
}

func (a *Adaptee) doubledError() (string, string) {
	return "We have an", "error"
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Error() string {
	s1, s2 := a.doubledError()
	return s1 + s2
}

func NewAdapter(a *Adaptee) error {
	return &Adapter{a}
}

func main() {
	adapter := NewAdapter(&Adaptee{})
	fmt.Printf("The error is: %s", adapter.Error())
}
