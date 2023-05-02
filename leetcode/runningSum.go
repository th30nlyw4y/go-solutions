// https://leetcode.com/problems/running-sum-of-1d-array/
package leetcode

func RunningSum(nums []int) []int {
	c := 0
	for i := range nums {
		c += nums[i]
		nums[i] = c
	}
	return nums
}
