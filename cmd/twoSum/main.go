// https://leetcode.com/problems/two-sum/
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums, 9)
	fmt.Printf("equeal: [0 1]\nresult: %v\n", res)
}

func twoSum(nums []int, target int) []int {
	var res []int
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if v1+v2 == target {
				res = []int{k1, k2}
				return res
			}
		}
	}
	return res
}
