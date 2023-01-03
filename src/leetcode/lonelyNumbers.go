// https://leetcode.com/problems/find-all-lonely-numbers-in-the-array/

package leetcode

func FindLonely(nums []int) []int {
	tmp, ans := map[int]bool{}, []int{}
	for _, number := range nums {
		_, okNum := tmp[number]
		_, okNumL := tmp[number-1]
		_, okNumR := tmp[number+1]
		if !(okNum || okNumL || okNumR) {
			tmp[number] = true
			continue
		}
		if okNumL {
			tmp[number-1] = false
		}
		if okNumR {
			tmp[number+1] = false
		}
		tmp[number] = false
	}
	for k, v := range tmp {
		if v {
			ans = append(ans, k)
		}
	}
	return ans
}
