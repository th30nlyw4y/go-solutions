// https://leetcode.com/problems/two-sum/
package leetcode

// v1
func TwoSumv1(nums []int, target int) []int {
	ans := []int{0, 0}
	nset := map[int]int{}
	for n := range nums {
		if _, ok := nset[nums[n]]; !ok {
			nset[nums[n]] = n
		}
	}
	for n := range nums {
		if index, ok := nset[target-nums[n]]; ok && n != index {
			ans[0] = n
			ans[1] = index
		}
	}
	return ans
}

// v2: leetcode solution
func TwoSumv2(nums []int, target int) []int {
	var ans []int
	nset := map[int]int{}
	for n := range nums {
		if _, ok := nset[nums[n]]; !ok {
			nset[nums[n]] = n
		}
		if index, ok := nset[target-nums[n]]; ok && n != index {
			ans = []int{n, index}
		}
	}
	return ans
}
