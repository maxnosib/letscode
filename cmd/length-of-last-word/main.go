// https://leetcode.com/problems/length-of-last-word/
package main

import (
	"fmt"
	"strings"
)

func main() {
	l := lengthOfLastWord("Hello World")
	fmt.Printf("equeal: 5\nresult: %v\n", l)
	l = lengthOfLastWord("HelloWorld")
	fmt.Printf("equeal: 10\nresult: %v\n", l)
	l = lengthOfLastWord("a ")
	fmt.Printf("equeal: 1\nresult: %v\n", l)
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	var res int
	for _, r := range s {
		res++
		if fmt.Sprintf("%c", r) == " " {
			res = 0
		}
	}
	return res
}
