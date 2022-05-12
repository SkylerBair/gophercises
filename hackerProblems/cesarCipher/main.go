package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	fmt.Println("length: %d\n", length)
	fmt.Println("input: %s\n", input)
	fmt.Println("delta: %d\n", delta)

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetLower)))
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetUpper)))
		default:
			ret = ret + string(ch)
		}
	}
	fmt.Println(ret)
}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return key[idx]
}

// func caesarCipher(s string, k []byte) string {
// 	// Write your code here
// 	abc := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
// 	trimmed := strings.TrimSpace(s)
// 	buf := []byte("")
// 	for letterIndex, letter := range trimmed {
// 		for abcIndex, abcLetter := range abc {
// 			if letter == abcLetter {
// 				encrypted := abc[abcIndex+k]
// 				buf = append(buf, []byte(encrypted))
// 			}
// 		}

// 	}
// 	return trimmed
// }
