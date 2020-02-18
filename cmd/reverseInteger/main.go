// https://leetcode.com/problems/reverse-integer/
package main

import "fmt"

func main() {
	res := reverse(123)
	fmt.Printf("equeal: 321\nresult: %v\n", res)
	res = reverse(-123)
	fmt.Printf("equeal: -321\nresult: %v\n", res)
	res = reverse(121)
	fmt.Printf("equeal: 121\nresult: %v\n", res)
	res = reverse(8085774586302733229)
	fmt.Printf("equeal: 0\nresult: %v\n", res)
}

func reverse(n int) int {
	denial := false

	if n < 0 {
		denial = true
		n *= -1
	}
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		if newInt == int(^uint(0)>>1)-7 {
			return 0
		}
		newInt += remainder
		n /= 10
	}
	if denial {
		newInt *= -1
	}

	return newInt
}
