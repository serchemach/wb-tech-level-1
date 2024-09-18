package main

import (
	"fmt"
	"reflect"
)

// Хотелось использовать просто тайп свитч, но с каналами это не совсем хорошо работает
func inferType(val interface{}) string {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	default:
		return "wtfisthis"
	}
}

func main() {
	a := 1
	b := "asdf"
	c := make(chan int)
	d := false
	fmt.Printf("The types are as follows:\n")
	fmt.Printf("%v:%s\n", a, inferType(a))
	fmt.Printf("%v:%s\n", b, inferType(b))
	fmt.Printf("%v:%s\n", c, inferType(c))
	fmt.Printf("%v:%s\n", d, inferType(d))
}
