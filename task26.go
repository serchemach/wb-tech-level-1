package main

import (
	"flag"
	"fmt"
	"strings"
)

func isSymbolUnique(s string) bool {
	symSet := make(map[rune]struct{})
	for _, rn := range strings.ToLower(s) {
		if _, ok := symSet[rn]; ok {
			return false
		}
		symSet[rn] = struct{}{}
	}
	return true
}

func main() {
	var s string
	flag.StringVar(&s, "str", "abcd", "String to check for symbol uniqueness")
	flag.Parse()

	res := isSymbolUnique(s)
	if res {
		fmt.Println("All symbols are unique! (lowercase)")
	} else {
		fmt.Println("Not all symbols are unique! (lowercase)")
	}

}
