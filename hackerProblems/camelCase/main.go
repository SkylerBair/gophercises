package main

import "unicode"

func main() {
	s := "dylanLovesButts"

	wordCounter(s)

}

func wordCounter(s string) int {
	words := 0

	for _, r := range s {
		if words == 0 && s != "" {
			words = 1
		} else if unicode.IsUpper(r) {
			words++
		}

	}
	return words
}
