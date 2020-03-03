// https://leetcode.com/problems/search-insert-position/submissions/
package main

import (
	"fmt"
)

func main() {
	nums := make(map[int][]int, 0)

	nums[5] = []int{1, 3, 5, 6} // 2
	nums[2] = []int{1, 3, 5, 6} // 1
	nums[7] = []int{1, 3, 5, 6} // 4
	nums[0] = []int{1, 3, 5, 6} // 0

	for k, v := range nums {
		r := searchInsert(v, k)
		fmt.Println("res:", r)
	}

}

func searchInsert(nums []int, target int) int {
	var r int
	for k, v := range nums {
		if v >= target {
			r = k
			break
		}
		if k == len(nums)-1 && target > v {
			r = k + 1
		}

	}
	return r
}
