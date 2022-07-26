package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LengthOfLongestSubstring(s string) int {
	ans, l, r, chars, sList := 0, 0, 0, map[rune]int{}, []rune(s)
	for r < len(sList) {
		if lastIndex, ok := chars[sList[r]]; ok {
			ans, l = max(r-l, ans), max(lastIndex+1, l)
		}
		chars[sList[r]] = r
		r++
	}
	return max(ans, r-l)
}
