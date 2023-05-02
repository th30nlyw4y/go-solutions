// https://leetcode.com/problems/search-insert-position/

package leetcode

func SearchInsertPosition(nums []int, target int) int {
	l, h := 0, len(nums)
	var m int
	for l < h {
		m = (l + h) / 2
		if nums[m] == target {
			return m
		} else {
			if nums[m] > target {
				h = m
			} else {
				l = m + 1
			}
		}
	}
	return l
}
