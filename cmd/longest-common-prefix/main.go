// https://leetcode.com/problems/longest-common-prefix
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("expected: fi\nresult: %s\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("expected: \nresult: %s\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	if strs[0] == "" {
		return ""
	}
	res := ""
	test := ""
	for _, char := range strs[0] {
		test += string(char)
		for _, v := range strs {
			if !strings.HasPrefix(v, test) {
				return res
			}
		}
		res = test
	}
	return res
}
