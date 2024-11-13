package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	args := os.Args[1]
	fmt.Println(abc(args))
}

func abc(s string) bool {
	var res string
	for _, char := range s {
		alph := strings.ToLower(string(char))
		if unicode.IsLetter(char) && !strings.Contains(res, alph) {
			res += alph
		}
	}
	return len(res) == 26
}
