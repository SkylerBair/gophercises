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

	var ret []rune
	for _, ch := range input {
		ret = append(ret, cipher(ch, delta))
	}
	fmt.Println(string(ret))
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base 
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

// func main() {
// 	var length, delta int
// 	var input string
// 	fmt.Scanf("%d\n", &length)
// 	fmt.Scanf("%s\n", &input)
// 	fmt.Scanf("%d\n", &delta)

// 	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
// 	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	ret := ""
// 	for _, ch := range input {
// 		switch {
// 		case strings.IndexRune(alphabetLower, ch) >= 0:
// 			ret = ret + string(rotate(ch, delta, []rune(alphabetLower)))
// 		case strings.IndexRune(alphabetUpper, ch) >= 0:
// 			ret = ret + string(rotate(ch, delta, []rune(alphabetUpper)))
// 		default:
// 			ret = ret + string(ch)
// 		}
// 	}
// 	fmt.Println(ret)
// }

// func rotate(s rune, delta int, key []rune) rune {
// 	idx := strings.IndexRune(string(key), s)
// 	if idx < 0 {
// 		panic("idx < 0")
// 	}
// 	idx = (idx + delta) % len(key)
// 	return key[idx]
// }


