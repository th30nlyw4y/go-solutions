package leetcode

func MinimumBuckets(street string) int {
	ans, i, s := 0, 0, []rune(street)
	for i < len(s) {
		if s[i] == 'H' {
			if i < len(s)-1 && s[i+1] == '.' {
				ans += 1
				i += 2
			} else if i > 0 && s[i-1] == '.' {
				ans += 1
			} else {
				return -1
			}
		}
		i++
	}
	return ans
}
