package main

import "fmt"

// Используем руны потому что юникод
func reverse(input string) string {
	runeRepr := []rune(input)
	n := len(runeRepr)
	for i := range n / 2 {
		runeRepr[i], runeRepr[n-1-i] = runeRepr[n-1-i], runeRepr[i]
	}
	return string(runeRepr)
}

func main() {
	strArr := []string{"главрыба", "123123jdkajfkjdkfsj", "РыбаСевер213"}
	for _, str := range strArr {
		fmt.Printf("String before reversing:\n%s\nAfter reversing:\n%s\n\n", str, reverse(str))
	}
}
