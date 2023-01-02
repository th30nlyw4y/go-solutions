// https://leetcode.com/problems/binary-search/
package leetcode

func Search(nums []int, target int) int {
	l, h := 0, len(nums)
	var m int
	for l < h {
		m = (h + l) / 2
		if nums[m] != target {
			if nums[m] > target {
				h = m
			} else {
				l = m + 1
			}
		} else {
			return m
		}
	}
	return -1
}
