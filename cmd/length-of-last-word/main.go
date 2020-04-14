// https://leetcode.com/problems/length-of-last-word/
package main

import "fmt"

func main() {
	l := lengthOfLastWord("Hello World")
	fmt.Printf("equeal: 5\nresult: %v\n", l)
	l = lengthOfLastWord("HelloWorld")
	fmt.Printf("equeal: 0\nresult: %v\n", l)
	l = lengthOfLastWord("a ")
	fmt.Printf("equeal: 0\nresult: %v\n", l)
}

func lengthOfLastWord(s string) int {
	var res int
	var isSpace bool
	for _, r := range s {
		if isSpace {
			res++
		}
		if fmt.Sprintf("%c", r) == " " {
			isSpace = true
			res = 0
		}
	}
	return res
}
