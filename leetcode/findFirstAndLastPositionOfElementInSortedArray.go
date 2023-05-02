//https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package leetcode

func index(n *[]int, t int, f func(index int) bool) int {
	ans, l, r := -1, 0, len(*n)
	for l < r {
		m := (l + r) / 2
		if (*n)[m] == t {
			ans = m
		}
		if f(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return ans
}

func SearchRange(nums []int, target int) []int {
	lInd := index(&nums, target, func(index int) bool { return nums[index] >= target })
	if lInd == -1 {
		return []int{-1, -1}
	}
	rInd := index(&nums, target, func(index int) bool { return nums[index] > target })
	return []int{lInd, rInd}
}
