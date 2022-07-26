// https://leetcode.com/problems/find-pivot-index/
package leetcode

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func PivotIndex(nums []int) int {
	ans, lSum, rSum := -1, 0, sum(nums)
	for i, v := range nums {
		rSum -= v
		if lSum == rSum {
			ans = i
			break
		}
		lSum += v
	}
	return ans
}
