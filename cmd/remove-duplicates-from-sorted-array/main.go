// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums := []int{1, 1, 2}
	lens := removeDuplicates(nums)
	for i := 0; i < lens; i++ {
		print(nums[i], "\n")
	}
}

func removeDuplicates(nums []int) int {
	var (
		curent = nums[0] - 1
		k      = 0
	)

	for _, v := range nums {
		if v != curent {
			curent = v
			nums[k] = v
			k++
			continue
		}
	}
	return k
}
