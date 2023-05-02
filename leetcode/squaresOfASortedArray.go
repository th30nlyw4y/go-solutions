// https://leetcode.com/problems/squares-of-a-sorted-array/

package leetcode

func SortedSquares(nums []int) []int {
	var ans = make([]int, len(nums))
	var l, r = 0, len(nums) - 1
	for i := len(nums) - 1; l <= r && l < len(nums) && r >= 0; i-- {
		sqrL, sqrR := nums[l]*nums[l], nums[r]*nums[r]
		if sqrL > sqrR {
			ans[i] = sqrL
			l++
		} else {
			ans[i] = sqrR
			r--
		}
	}
	return ans
}
