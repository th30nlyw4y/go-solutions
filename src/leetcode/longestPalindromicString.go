// https://leetcode.com/problems/longest-palindromic-substring/

package leetcode

func maxStrLen(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}

func longestPalindrome(s string) string {
	var ans string
	var l, h int
	switch len(s) {
	case 1:
		return s
	case 2:
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	default:
		for i := 1; i < len(s)-1; i++ {
			l, h = i, i
			for {
				if l > 0 && s[l-1] == s[i] {
					l--
					continue
				}
				if h < len(s)-1 && s[h+1] == s[i] {
					h++
					continue
				}
				break
			}
			for l > 0 && h < len(s)-1 && s[l-1] == s[h+1] {
				l--
				h++
			}
			ans = maxStrLen(ans, s[l:h+1])
		}
	}
	return ans
}
