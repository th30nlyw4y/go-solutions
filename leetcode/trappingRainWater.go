// https://leetcode.com/problems/trapping-rain-water/
package leetcode

func Trap(height []int) int {
	ans, left, right := 0, 0, len(height)-1
	lMax, rMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= lMax {
				lMax = height[left]
			} else {
				ans += (lMax - height[left])
			}
			left++
		} else {
			if height[right] >= rMax {
				rMax = height[right]
			} else {
				ans += (rMax - height[right])
			}
			right--
		}
	}
	return ans
}
