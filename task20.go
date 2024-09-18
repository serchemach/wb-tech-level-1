package main

import "fmt"

// На всякий случай возьмём руны, иначе опять юникод заболеет
func reverseWords(input string) string {
	runeRepr := []rune(input)
	n := len(runeRepr)
	outputRepr := make([]rune, n)
	lastStrBegin := 0
	for i := range n {
		if runeRepr[i] == ' ' && i != 0 {
			copy(outputRepr[n-i:], runeRepr[lastStrBegin:i])
			outputRepr[n-i-1] = ' '
			lastStrBegin = i + 1
		}
	}

	if n > 0 {
		copy(outputRepr[0:], runeRepr[lastStrBegin:])
	}

	return string(outputRepr)
}

func main() {
	sentences := []string{"snow dog sun", "Разработать программу которая переворачивает слова в строке"}
	for _, sentence := range sentences {
		fmt.Printf("Sentence before reversing the words:\n%s\nAfter reversing:\n%s\n\n", sentence, reverseWords(sentence))
	}

}
