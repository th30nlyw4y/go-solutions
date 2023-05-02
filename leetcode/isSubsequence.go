// https://leetcode.com/problems/is-subsequence/
package leetcode

func IsSubsequenceV1(s, t string) bool {
	if s != "" {
		sLen, cnt := len(s), 0
		for i := 0; i < len(t) && cnt < sLen; i++ {
			if s[cnt] == t[i] {
				cnt++
			}
		}
		return cnt == sLen
	} else {
		return true
	}
}
